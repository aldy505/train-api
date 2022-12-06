package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

func main() {
	hostname, ok := os.LookupEnv("HTTP_HOSTNAME")
	if !ok {
		hostname = "127.0.0.1"
	}

	port, ok := os.LookupEnv("HTTP_PORT")
	if !ok {
		port = "5000"
	}

	stations, err := ParseStationConfiguration()
	if err != nil {
		log.Fatalf("parsing station configuration: %s", err.Error())
	}

	routes, err := ParseRoutesConfiguration()
	if err != nil {
		log.Fatalf("parsing routes configuration: %s", err.Error())
	}

	var lines []*Line
	for i, route := range routes {
		trainForRoute := GenerateTrains(route.Trains, 100*(i+1), route.ID, 0, 0)

		line, err := MapStationsToLine(route, stations, trainForRoute)
		if err != nil {
			log.Fatalf("mapping stations to line: %s", err.Error())
		}

		lines = append(lines, &line)
	}

	app := chi.NewRouter()
	app.Use(middleware.CleanPath)
	app.Use(middleware.Recoverer)
	app.Use(middleware.NoCache)
	app.Use(cors.AllowAll().Handler)
	app.Mount("/", NewHttpRouter(lines))

	httpServer := &http.Server{
		Addr:              net.JoinHostPort(hostname, port),
		Handler:           app,
		ReadTimeout:       time.Minute,
		ReadHeaderTimeout: time.Minute,
		WriteTimeout:      time.Minute,
		IdleTimeout:       time.Minute,
	}

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt)

	for _, line := range lines {
		go func(line *Line) {
			log.Printf("Starting new worker for %d - %s", line.ID, line.Name)
			worker := NewWorker(line)
			_ = worker.Start()
		}(line)
	}

	go func() {
		log.Printf("HTTP server is listening on: %s", httpServer.Addr)
		err := httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listening http server: %s", err.Error())
		}
	}()

	<-exitSignal

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Minute)
	defer shutdownCancel()

	err = httpServer.Shutdown(shutdownCtx)
	if err != nil {
		log.Printf("shutting down http server: %s", err.Error())
	}
}
