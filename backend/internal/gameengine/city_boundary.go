package gameengine

import "github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/domain/id"

// GeoCoordinate represents a geographic coordinate in a city boundary ring.
type GeoCoordinate struct {
	Latitude  float64
	Longitude float64
}

// CityBoundary represents a city's single ordered outer polygon ring.
type CityBoundary struct {
	CityID    id.CityID
	OuterRing []GeoCoordinate
}
