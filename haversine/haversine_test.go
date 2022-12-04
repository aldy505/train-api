package haversine_test

import (
	"testing"

	"train-api/haversine"
)

var tests = []struct {
	p     haversine.Coordinate
	q     haversine.Coordinate
	outMi float64
	outKm float64
}{
	{
		haversine.Coordinate{Latitude: 22.55, Longitude: 43.12},  // Rio de Janeiro, Brazil
		haversine.Coordinate{Latitude: 13.45, Longitude: 100.28}, // Bangkok, Thailand
		3786.251258825624,
		6094.544408786774,
	},
	{
		haversine.Coordinate{Latitude: 20.10, Longitude: 57.30}, // Port Louis, Mauritius
		haversine.Coordinate{Latitude: 0.57, Longitude: 100.21}, // Padang, Indonesia
		3196.671009759937,
		5145.525771394785,
	},
	{
		haversine.Coordinate{Latitude: 51.45, Longitude: 1.15},  // Oxford, United Kingdom
		haversine.Coordinate{Latitude: 41.54, Longitude: 12.27}, // Vatican, City Vatican City
		863.0311907424888,
		1389.1793118293067,
	},
	{
		haversine.Coordinate{Latitude: 22.34, Longitude: 17.05}, // Windhoek, Namibia
		haversine.Coordinate{Latitude: 51.56, Longitude: 4.29},  // Rotterdam, Netherlands
		2130.8298370015464,
		3429.89310043882,
	},
	{
		haversine.Coordinate{Latitude: 63.24, Longitude: 56.59}, // Esperanza, Argentina
		haversine.Coordinate{Latitude: 8.50, Longitude: 13.14},  // Luanda, Angola
		4346.398369403186,
		6996.18595539861,
	},
	{
		haversine.Coordinate{Latitude: 90.00, Longitude: 0.00}, // North/South Poles
		haversine.Coordinate{Latitude: 48.51, Longitude: 2.21}, // Paris,  France
		2866.1346681303867,
		4613.477506482742,
	},
	{
		haversine.Coordinate{Latitude: 45.04, Longitude: 7.42},  // Turin, Italy
		haversine.Coordinate{Latitude: 3.09, Longitude: 101.42}, // Kuala Lumpur, Malaysia
		6261.05275709582,
		10078.111954385415,
	},
}

func TestHaversineDistance(t *testing.T) {
	for _, input := range tests {
		mi, km := haversine.Distance(input.p, input.q)

		if input.outMi != mi || input.outKm != km {
			t.Errorf("fail: want %v %v -> %v %v got %v %v",
				input.p,
				input.q,
				input.outMi,
				input.outKm,
				mi,
				km,
			)
		}
	}
}
