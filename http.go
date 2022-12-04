package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HttpDependency struct {
	lines []*Line
}

func NewHttpRouter(lines []*Line) *chi.Mux {
	dependency := &HttpDependency{lines: lines}

	r := chi.NewRouter()

	r.Get("/stations", dependency.GetStations)
	r.Get("/trains", dependency.GetTrains)
	r.Get("/lines", dependency.GetLines)

	return r
}

func (h *HttpDependency) GetStations(w http.ResponseWriter, r *http.Request) {
	var stations []*Station
	for _, line := range h.lines {
		stations = append(stations, line.Stations...)
	}

	marshalled, err := json.Marshal(stations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshalled)
}

func (h *HttpDependency) GetTrains(w http.ResponseWriter, r *http.Request) {
	var trains []*Train
	for _, line := range h.lines {
		trains = append(trains, line.Trains...)
	}

	marshalled, err := json.Marshal(trains)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshalled)
}

func (h *HttpDependency) GetLines(w http.ResponseWriter, r *http.Request) {
	marshalled, err := json.Marshal(h.lines)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshalled)
}
