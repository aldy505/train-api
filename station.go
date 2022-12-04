package main

import "time"

type StationSchedule struct {
	TrainId   int
	LineId    int
	Arrival   time.Time
	Departure time.Time
	Direction TrainDirection
}

type Station struct {
	ID               int
	Name             string
	Latitude         float64
	Longitude        float64
	TrainWaitingTime time.Duration
	NextSchedule     []StationSchedule
}
