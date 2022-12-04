package haversine

import (
	"math"
)

const (
	earthRadiusMi = 3958 // radius of the earth in miles.
	earthRadiusKm = 6371 // radius of the earth in kilometers.
)

// Coordinate represents a geographic coordinate.
type Coordinate struct {
	Latitude  float64
	Longitude float64
}

func NewCoordinate(latitude float64, longitude float64) Coordinate {
	return Coordinate{latitude, longitude}
}

// degreesToRadians converts from degrees to radians.
func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

// Distance calculates the shortest path between two coordinates on the surface
// of the Earth. This function returns two units of measure, the first is the
// distance in miles, the second is the distance in kilometers.
func Distance(p, q Coordinate) (mi, km float64) {
	lat1 := degreesToRadians(p.Latitude)
	lon1 := degreesToRadians(p.Longitude)
	lat2 := degreesToRadians(q.Latitude)
	lon2 := degreesToRadians(q.Longitude)

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	mi = c * earthRadiusMi
	km = c * earthRadiusKm

	return mi, km
}
