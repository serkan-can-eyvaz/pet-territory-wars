package gameengine

import (
	"testing"
	"time"

	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/domain/id"
)

func TestResolveWalkInputStoresDomainState(t *testing.T) {
	walkID := mustWalkID(t)
	playerID := mustPlayerID(t)
	cityID := mustCityID(t)
	engineVersion := mustEngineVersion(t)
	hexID := id.NewHexID(0x8928308280fffff)
	startedAt := time.Date(2026, time.July, 24, 10, 0, 0, 0, time.UTC)
	endedAt := startedAt.Add(10 * time.Minute)
	evaluatedAt := endedAt.Add(time.Minute)
	lastUpdatedAt := startedAt.Add(-24 * time.Hour)
	route := []LocationPoint{
		{
			Sequence:       1,
			Latitude:       41.0082,
			Longitude:      28.9784,
			RecordedAt:     startedAt,
			AccuracyMeters: 5,
			IsMockLocation: false,
		},
	}
	boundary := CityBoundary{
		CityID: cityID,
		OuterRing: []GeoCoordinate{
			{Latitude: 41.0082, Longitude: 28.9784},
		},
	}
	existingHexes := map[id.HexID]HexState{
		hexID: {
			HexID:         hexID,
			OwnerID:       &playerID,
			Dominance:     60,
			LastUpdatedAt: lastUpdatedAt,
			Version:       12,
		},
	}

	input := ResolveWalkInput{
		WalkID:        walkID,
		PlayerID:      playerID,
		CityID:        cityID,
		StartedAt:     startedAt,
		EndedAt:       endedAt,
		EvaluatedAt:   evaluatedAt,
		Route:         route,
		ExistingHexes: existingHexes,
		Boundary:      boundary,
		EngineVersion: engineVersion,
	}

	if input.WalkID != walkID || input.PlayerID != playerID || input.CityID != cityID {
		t.Error("ResolveWalkInput did not preserve domain IDs")
	}
	if input.EngineVersion != engineVersion {
		t.Error("ResolveWalkInput did not preserve EngineVersion")
	}
	if input.StartedAt != startedAt || input.EndedAt != endedAt || input.EvaluatedAt != evaluatedAt {
		t.Error("ResolveWalkInput did not preserve timestamps")
	}
	if input.Boundary.CityID != boundary.CityID || len(input.Boundary.OuterRing) != len(boundary.OuterRing) {
		t.Error("ResolveWalkInput did not preserve CityBoundary")
	}
	for index, coordinate := range boundary.OuterRing {
		if input.Boundary.OuterRing[index] != coordinate {
			t.Errorf("Boundary.OuterRing[%d] = %+v, want %+v", index, input.Boundary.OuterRing[index], coordinate)
		}
	}
	if len(input.Route) != 1 || input.Route[0] != route[0] {
		t.Errorf("Route = %+v, want %+v", input.Route, route)
	}
	if actual, ok := input.ExistingHexes[hexID]; !ok || actual != existingHexes[hexID] {
		t.Errorf("ExistingHexes[%s] = %+v, want %+v", hexID, actual, existingHexes[hexID])
	}
}

func TestLocationPointStoresValues(t *testing.T) {
	recordedAt := time.Date(2026, time.July, 24, 10, 0, 0, 0, time.UTC)
	point := LocationPoint{
		Sequence:       2,
		Latitude:       41.015,
		Longitude:      28.985,
		RecordedAt:     recordedAt,
		AccuracyMeters: 8.5,
		IsMockLocation: true,
	}

	if point.Sequence != 2 || point.Latitude != 41.015 || point.Longitude != 28.985 {
		t.Errorf("LocationPoint = %+v, want configured values", point)
	}
	if point.RecordedAt != recordedAt || point.AccuracyMeters != 8.5 || !point.IsMockLocation {
		t.Errorf("LocationPoint = %+v, want configured values", point)
	}
}

func TestHexStateStoresOwnedAndUnownedValues(t *testing.T) {
	playerID := mustPlayerID(t)
	hexID := id.NewHexID(0x8928308280fffff)
	lastUpdatedAt := time.Date(2026, time.July, 24, 10, 0, 0, 0, time.UTC)

	owned := HexState{
		HexID:         hexID,
		OwnerID:       &playerID,
		Dominance:     60,
		LastUpdatedAt: lastUpdatedAt,
		Version:       3,
	}
	unowned := HexState{HexID: hexID}

	if owned.OwnerID == nil || *owned.OwnerID != playerID {
		t.Error("owned HexState did not preserve OwnerID")
	}
	if owned.Dominance != 60 || owned.LastUpdatedAt != lastUpdatedAt || owned.Version != 3 {
		t.Errorf("owned HexState = %+v, want configured values", owned)
	}
	if unowned.OwnerID != nil {
		t.Errorf("unowned OwnerID = %v, want nil", unowned.OwnerID)
	}
}

func mustWalkID(t *testing.T) id.WalkID {
	t.Helper()

	value, err := id.NewWalkID("7b0f04af-37dd-4a69-9bb1-3d8f0c59c8f8")
	if err != nil {
		t.Fatalf("create walk ID: %v", err)
	}

	return value
}

func mustPlayerID(t *testing.T) id.PlayerID {
	t.Helper()

	value, err := id.NewPlayerID("c55e1241-73df-46f2-b10d-c7b56e967024")
	if err != nil {
		t.Fatalf("create player ID: %v", err)
	}

	return value
}

func mustCityID(t *testing.T) id.CityID {
	t.Helper()

	value, err := id.NewCityID("1276c9f1-97a9-48a8-81d2-063d80a221d5")
	if err != nil {
		t.Fatalf("create city ID: %v", err)
	}

	return value
}

func mustEngineVersion(t *testing.T) id.EngineVersion {
	t.Helper()

	value, err := id.NewEngineVersion("1.0.0")
	if err != nil {
		t.Fatalf("create engine version: %v", err)
	}

	return value
}
