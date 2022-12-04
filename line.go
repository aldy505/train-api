package main

import (
	"train-api/clock"
)

type Line struct {
	ID               int
	Name             string
	OperationalStart clock.Clock
	OperationalStop  clock.Clock
	Stations         []*Station
	Trains           []*Train
}

func MapStationsToLine(lineConfiguration LineConfiguration, stations []*Station, trains []*Train) (Line, error) {
	operationalStart, err := clock.Parse(lineConfiguration.OperationalStart)
	if err != nil {
		return Line{}, err
	}

	operationalStop, err := clock.Parse(lineConfiguration.OperationalStop)
	if err != nil {
		return Line{}, err
	}

	var filteredStations []*Station
	for _, stationId := range lineConfiguration.StationIDs {
		for index, station := range stations {
			if stationId == station.ID {
				filteredStations = append(filteredStations, stations[index])
				stations[index].NextSchedule = append(
					stations[index].NextSchedule,
					StationSchedule{
						LineId:    lineConfiguration.ID,
						Direction: TrainDirectionUp,
					},
					StationSchedule{
						LineId:    lineConfiguration.ID,
						Direction: TrainDirectionDown,
					},
				)
				break
			}
		}
	}

	return Line{
		ID:               lineConfiguration.ID,
		Name:             lineConfiguration.Name,
		OperationalStart: operationalStart,
		OperationalStop:  operationalStop,
		Stations:         filteredStations,
		Trains:           trains,
	}, nil
}
