package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type HttpDependency struct {
	lines []*Line
}

func NewHttpRouter(lines []*Line) *chi.Mux {
	dependency := &HttpDependency{lines: lines}

	r := chi.NewRouter()

	r.Get("/", dependency.Index)
	r.Get("/stations", dependency.GetStations)
	r.Get("/trains", dependency.GetTrains)
	r.Get("/lines", dependency.GetLines)

	return r
}

func (h *HttpDependency) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`# Train API

This is a public API meant for be used publicly for frontends.
This is also a good project for backend developers that can only
code CRUD stuff. Now, you have to think in states, as there are
no database to begin with.

I intentionally uses real train stations and coordinates from those
that exists in Jakarta to mimic a real train arrival/departure
behavior. Please do not use this API and advertise it
as if this is a real train schedule. It's definitely not.
BUT, if you are advertising the API as a real train schedule,
please stop. Don't do that.

THIS API IS MEANT FOR EDUCATIONAL PURPOSE.
I AM NOT RESPONSIBLE FOR ANY ABUSE IN THE NAME OF USING THIS PROGRAM.
NOR I DO NOT GAIN ANY MONEY FROM CREATING THIS PROGRAM.

The endpoints are pretty simple, you are most welcome to open a PR
and submit a new endpoint in case you need them:

	### Get all lines
	GET https://train.reinaldyrafli.com/lines
	Accept: application/json

	### Get lines by ID
	GET https://train.reinaldyrafli.com/lines?id=1
	Accept: application/json

	### Get all stations
	GET https://train.reinaldyrafli.com/stations
	Accept: application/json

	### Get station by ID
	GET https://train.reinaldyrafli.com/stations?id=1
	Accept: application/json

	### Get all trains
	GET https://train.reinaldyrafli.com/trains
	Accept: application/json

	### Get trains by ID
	GET https://train.reinaldyrafli.com/trains?id=100
	Accept: application/json

## License

	Copyright 2022 Reinaldy Rafli

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
`))
}

func (h *HttpDependency) GetStations(w http.ResponseWriter, r *http.Request) {
	var stations []*Station
	for _, line := range h.lines {
		stations = append(stations, line.Stations...)
	}

	if r.URL.Query().Has("id") {
		byID := r.URL.Query().Get("id")
		var foundIndex int = -1
		for i, station := range stations {
			if strconv.Itoa(station.ID) == byID {
				foundIndex = i
				break
			}
		}

		if foundIndex == -1 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}

		marshalled, err := json.Marshal([]*Station{stations[foundIndex]})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(marshalled)
		return
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

	if r.URL.Query().Has("id") {
		byID := r.URL.Query().Get("id")
		var foundIndex int = -1
		for i, train := range trains {
			if strconv.Itoa(train.ID) == byID {
				foundIndex = i
				break
			}
		}

		if foundIndex == -1 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}

		marshalled, err := json.Marshal([]*Train{trains[foundIndex]})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(marshalled)
		return
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
	if r.URL.Query().Has("id") {
		byID := r.URL.Query().Get("id")
		var foundIndex int = -1
		for i, line := range h.lines {
			if strconv.Itoa(line.ID) == byID {
				foundIndex = i
				break
			}
		}

		if foundIndex == -1 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}

		marshalled, err := json.Marshal([]*Line{h.lines[foundIndex]})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(marshalled)
		return
	}

	marshalled, err := json.Marshal(h.lines)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshalled)
}
