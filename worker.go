package main

import (
	"log"
	"math/rand"
	"time"

	"train-api/clock"
	"train-api/haversine"
)

type Worker struct {
	line *Line
}

func NewWorker(line *Line) *Worker {
	return &Worker{
		line: line,
	}
}

func (w *Worker) Start() error {
	for i, train := range w.line.Trains {
		go func(index int, train *Train) {
			log.Printf("[%s] Deploying a train", w.line.Name)
			w.startTrain(train)
		}(i, train)

		time.Sleep(time.Minute * 5)
	}

	log.Printf("[%s] All train has been deployed", w.line.Name)

	return nil
}

func (w *Worker) startTrain(train *Train) {
	// Create a goroutine that will change the OutOfService to true during line closing.
	go func(t *Train) {
		for {
			if clock.Now().After(w.line.OperationalStop) {
				train.Lock()
				train.OutOfService = true
				train.Unlock()

				log.Printf("[%s] Train ID %d has been marked as out of service", w.line.Name, train.ID)
				time.Sleep(time.Hour * 23)
			}

			time.Sleep(time.Second)
		}
	}(train)

	var currentStationIndex = 0
	var nextStationIndex = -1
	var nextDirection TrainDirection
	var currentTimeSecondIndex float64 = 0
	var nextStationArrivalTime time.Time
	var tripTimeToNextStation float64 = 0

	train.CurrentLatitude = w.line.Stations[0].Latitude
	train.CurrentLongitude = w.line.Stations[0].Longitude
	train.OutOfService = false
	train.StoppingAtStation = true

	for {
		if clock.Now().Before(w.line.OperationalStart) {
			sleepDuration := time.Since(time.Date(
				time.Now().Year(),
				time.Now().Month(),
				time.Now().Day(),
				w.line.OperationalStart.Hour,
				w.line.OperationalStart.Minute,
				w.line.OperationalStart.Second,
				0,
				time.Local))

			log.Printf("[%s] Train ID %d will be sleeping for %s", w.line.Name, train.ID, sleepDuration.String())
			time.Sleep(sleepDuration)

			train.OutOfService = false
			train.StoppingAtStation = true

			log.Printf("[%s] Train ID %d is back on service", w.line.Name, train.ID)
		}

		if clock.Now().After(w.line.OperationalStop) {
			time.Sleep(time.Minute * 30)
			continue
		}

		if train.OutOfService && currentStationIndex == 0 {
			train.Lock()
			train.TotalPassengers = 0
			train.StoppingAtStation = false
			nextStationIndex = -1
			nextStationArrivalTime = time.Time{}
			train.Unlock()
			continue
		}

		if train.StoppingAtStation {
			if nextStationIndex >= 0 {
				currentStationIndex = nextStationIndex
			}

			log.Printf("[%s] Train ID %d has arrived on station %s", w.line.Name, train.ID, w.line.Stations[currentStationIndex].Name)

			train.CurrentLongitude = w.line.Stations[currentStationIndex].Longitude
			train.CurrentLatitude = w.line.Stations[currentStationIndex].Latitude

			// distance to next station
			nextStationIndex, nextDirection = getNextStationIndex(w.line.Stations, currentStationIndex, train.Direction)
			_, distance := haversine.Distance(
				haversine.NewCoordinate(w.line.Stations[currentStationIndex].Latitude, w.line.Stations[currentStationIndex].Longitude),
				haversine.NewCoordinate(w.line.Stations[nextStationIndex].Latitude, w.line.Stations[nextStationIndex].Longitude),
			)

			tripTimeToNextStation = measureTripTime(distance)
			nextStationArrivalTime = time.Now().Add(time.Second * time.Duration(tripTimeToNextStation))

			var nextScheduleIndex int
			for i, schedule := range w.line.Stations[nextStationIndex].NextSchedule {
				if schedule.LineId == w.line.ID && schedule.Direction == nextDirection {
					nextScheduleIndex = i
				}
			}

			time.Sleep(w.line.Stations[currentStationIndex].TrainWaitingTime)

			w.line.Stations[nextStationIndex].NextSchedule[nextScheduleIndex].TrainId = train.ID
			w.line.Stations[nextStationIndex].NextSchedule[nextScheduleIndex].Arrival = nextStationArrivalTime
			w.line.Stations[nextStationIndex].NextSchedule[nextScheduleIndex].Departure = nextStationArrivalTime.Add(w.line.Stations[nextStationIndex].TrainWaitingTime)

			train.Lock()
			train.Direction = nextDirection
			train.TotalPassengers = rand.Int31n(300) + 1
			train.StoppingAtStation = false
			train.DestinationStationId = w.line.Stations[nextStationIndex].ID
			train.Unlock()

			log.Printf("[%s] Train ID %d has departed from station %s bound for %s", w.line.Name, train.ID, w.line.Stations[currentStationIndex].Name, w.line.Stations[nextStationIndex].Name)
		}

		time.Sleep(time.Millisecond * 100)
		currentTimeSecondIndex += 0.1
		latitude, longitude := measureCurrentCoordinate(
			currentTimeSecondIndex,
			tripTimeToNextStation,
			w.line.Stations[currentStationIndex].Latitude, w.line.Stations[currentStationIndex].Longitude,
			w.line.Stations[nextStationIndex].Latitude, w.line.Stations[nextStationIndex].Longitude)

		train.CurrentLongitude = longitude
		train.CurrentLatitude = latitude

		if time.Now().After(nextStationArrivalTime) {
			train.Lock()
			train.StoppingAtStation = true
			train.Unlock()
		}
	}
}

func getNextStationIndex(stations []*Station, currentIndex int, direction TrainDirection) (int, TrainDirection) {
	if len(stations)-1 == currentIndex && direction == TrainDirectionUp {
		return currentIndex, TrainDirectionDown
	}

	if direction == TrainDirectionUp {
		return currentIndex + 1, TrainDirectionUp
	}

	if currentIndex == 0 {
		return 0, TrainDirectionUp
	}

	return currentIndex - 1, TrainDirectionDown
}

func measureTripTime(distanceKm float64) float64 {
	return distanceKm / 0.0138888889 // 50 km/hours in km/second
}

func measureCurrentCoordinate(elapsedTime float64, totalTime float64, initialLatitude float64, initialLongitude float64, destinationLatitude float64, destinationLongitude float64) (latitude float64, longitude float64) {
	dx := destinationLongitude - initialLongitude
	dy := destinationLatitude - initialLatitude

	longitude = initialLongitude + elapsedTime/totalTime*dx
	latitude = initialLatitude + elapsedTime/totalTime*dy
	return
}
