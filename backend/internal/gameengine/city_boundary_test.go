package gameengine

import (
	"testing"

	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/domain/id"
)

func TestCityBoundaryStoresCityAndOrderedOuterRing(t *testing.T) {
	cityID, err := id.NewCityID("7b0f04af-37dd-4a69-9bb1-3d8f0c59c8f8")
	if err != nil {
		t.Fatalf("create city ID: %v", err)
	}

	ring := []GeoCoordinate{
		{Latitude: 41.0082, Longitude: 28.9784},
		{Latitude: 41.0100, Longitude: 28.9800},
		{Latitude: 41.0120, Longitude: 28.9760},
	}
	boundary := CityBoundary{
		CityID:    cityID,
		OuterRing: ring,
	}

	if boundary.CityID != cityID {
		t.Errorf("CityID = %q, want %q", boundary.CityID, cityID)
	}
	if len(boundary.OuterRing) != len(ring) {
		t.Fatalf("OuterRing length = %d, want %d", len(boundary.OuterRing), len(ring))
	}
	for index, coordinate := range ring {
		if boundary.OuterRing[index] != coordinate {
			t.Errorf("OuterRing[%d] = %+v, want %+v", index, boundary.OuterRing[index], coordinate)
		}
	}
}
