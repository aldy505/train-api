package main

import (
	"errors"
	"os"
	"time"

	"github.com/pelletier/go-toml/v2"
)

type StationConfigurationFile struct {
	Station []struct {
		ID               int
		Name             string
		Latitude         float64
		Longitude        float64
		TrainWaitingTime string
	} `toml:"station"`
}

func ParseStationConfiguration() ([]*Station, error) {
	files, err := os.ReadDir("configurations/stations")
	if err != nil {
		return []*Station{}, err
	}

	var stations []*Station

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		f, err := os.Open("configurations/stations/" + file.Name())
		if err != nil {
			return []*Station{}, err
		}

		var station StationConfigurationFile
		err = toml.NewDecoder(f).Decode(&station)
		if err != nil {
			return []*Station{}, err
		}

		err = f.Close()
		if err != nil && !errors.Is(err, os.ErrClosed) {
			return []*Station{}, err
		}

		for _, s := range station.Station {
			parsedWaitingTime, err := time.ParseDuration(s.TrainWaitingTime)
			if err != nil {
				return []*Station{}, err
			}

			stations = append(stations, &Station{
				ID:               s.ID,
				Name:             s.Name,
				Latitude:         s.Latitude,
				Longitude:        s.Longitude,
				TrainWaitingTime: parsedWaitingTime,
				NextSchedule:     []StationSchedule{},
			})
		}
	}

	return stations, nil
}

type LineConfigurationFile struct {
	Line []LineConfiguration `toml:"line"`
}

type LineConfiguration struct {
	ID               int
	Name             string
	StationIDs       []int
	OperationalStart string
	OperationalStop  string
	Trains           int
}

func ParseRoutesConfiguration() ([]LineConfiguration, error) {
	files, err := os.ReadDir("configurations/routes")
	if err != nil {
		return []LineConfiguration{}, err
	}

	var lines []LineConfiguration

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		f, err := os.Open("configurations/routes/" + file.Name())
		if err != nil {
			return []LineConfiguration{}, err
		}

		var line LineConfigurationFile
		err = toml.NewDecoder(f).Decode(&line)
		if err != nil {
			return []LineConfiguration{}, err
		}

		err = f.Close()
		if err != nil && !errors.Is(err, os.ErrClosed) {
			return []LineConfiguration{}, err
		}

		lines = append(lines, line.Line...)
	}

	return lines, nil
}

func GenerateTrains(amount int, startingId int, lineId int, startingLatitude float64, startingLongitude float64) []*Train {
	var trains []*Train
	var currentId = startingId

	for i := 0; i < amount; i++ {
		trains = append(trains, &Train{
			ID:                currentId,
			LineID:            lineId,
			CurrentLatitude:   startingLatitude,
			CurrentLongitude:  startingLongitude,
			Direction:         TrainDirectionUp,
			StoppingAtStation: false,
			TotalPassengers:   0,
			OutOfService:      true,
		})

		currentId += 1
	}

	return trains
}
