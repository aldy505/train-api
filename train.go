package main

import "sync"

type TrainDirection string

const (
	TrainDirectionUp   TrainDirection = "UP"
	TrainDirectionDown TrainDirection = "DOWN"
)

type Train struct {
	ID                   int
	LineID               int
	CurrentLatitude      float64
	CurrentLongitude     float64
	Direction            TrainDirection
	StoppingAtStation    bool
	TotalPassengers      int32
	OutOfService         bool
	DestinationStationId int

	sync.Mutex
}
