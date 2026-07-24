package gameengine

import (
	"testing"
	"time"

	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/domain/id"
)

func TestAggregateVisitedHexes_EmptyInput(t *testing.T) {
	t.Parallel()

	visits := aggregateVisitedHexes(nil)
	if len(visits) != 0 {
		t.Fatalf("visit count = %d, want 0", len(visits))
	}
}

func TestAggregateVisitedHexes_SingleSample(t *testing.T) {
	t.Parallel()

	enteredAt := time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC)
	visits := aggregateVisitedHexes([]hexRouteSample{
		hexSample(1, enteredAt, 0, 0),
	})

	assertVisitedHexes(t, visits, []VisitedHex{{
		HexID:          1,
		EntryCount:     1,
		FirstEnteredAt: enteredAt,
		LastExitedAt:   enteredAt,
	}})
}

func TestAggregateVisitedHexes_AttributesIntervalsToStartingHex(t *testing.T) {
	t.Parallel()

	startedAt := time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC)
	visits := aggregateVisitedHexes([]hexRouteSample{
		hexSample(1, startedAt, 0, 0),
		hexSample(1, startedAt.Add(10*time.Second), 0, 10),
		hexSample(2, startedAt.Add(20*time.Second), 0, 20),
		hexSample(1, startedAt.Add(30*time.Second), 0, 30),
	})

	assertVisitedHexes(t, visits, []VisitedHex{
		{
			HexID:           1,
			DistanceMeters:  20,
			PresenceSeconds: 20,
			EntryCount:      2,
			FirstEnteredAt:  startedAt,
			LastExitedAt:    startedAt.Add(30 * time.Second),
		},
		{
			HexID:           2,
			DistanceMeters:  10,
			PresenceSeconds: 10,
			EntryCount:      1,
			FirstEnteredAt:  startedAt.Add(20 * time.Second),
			LastExitedAt:    startedAt.Add(30 * time.Second),
		},
	})
}

func TestAggregateVisitedHexes_DoesNotConnectSourceSegments(t *testing.T) {
	t.Parallel()

	startedAt := time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC)
	visits := aggregateVisitedHexes([]hexRouteSample{
		hexSample(1, startedAt, 0, 0),
		hexSample(1, startedAt.Add(10*time.Second), 0, 10),
		hexSample(1, startedAt.Add(20*time.Second), 1, 0),
		hexSample(1, startedAt.Add(30*time.Second), 1, 10),
	})

	assertVisitedHexes(t, visits, []VisitedHex{{
		HexID:           1,
		DistanceMeters:  20,
		PresenceSeconds: 20,
		EntryCount:      2,
		FirstEnteredAt:  startedAt,
		LastExitedAt:    startedAt.Add(30 * time.Second),
	}})
}

func TestAggregateVisitedHexes_SortsByFirstEntryThenHexIDString(t *testing.T) {
	t.Parallel()

	enteredAt := time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC)
	visits := aggregateVisitedHexes([]hexRouteSample{
		hexSample(2, enteredAt, 0, 0),
		hexSample(10, enteredAt, 1, 0),
	})

	assertVisitedHexes(t, visits, []VisitedHex{
		{HexID: 10, EntryCount: 1, FirstEnteredAt: enteredAt, LastExitedAt: enteredAt},
		{HexID: 2, EntryCount: 1, FirstEnteredAt: enteredAt, LastExitedAt: enteredAt},
	})
}

func TestAggregateVisitedHexes_IsDeterministicAndDoesNotMutateInput(t *testing.T) {
	t.Parallel()

	startedAt := time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC)
	samples := []hexRouteSample{
		hexSample(1, startedAt, 0, 0),
		hexSample(2, startedAt.Add(10*time.Second), 0, 10),
		hexSample(1, startedAt.Add(20*time.Second), 0, 20),
	}
	original := append([]hexRouteSample(nil), samples...)

	first := aggregateVisitedHexes(samples)
	second := aggregateVisitedHexes(samples)
	assertVisitedHexes(t, second, first)
	for index := range samples {
		if samples[index] != original[index] {
			t.Fatalf("sample %d was mutated", index)
		}
	}
}

func hexSample(hexID id.HexID, recordedAt time.Time, sourceSegmentIndex int, distanceMeters float64) hexRouteSample {
	return hexRouteSample{
		HexID:                          hexID,
		RecordedAt:                     recordedAt,
		SourceSegmentIndex:             sourceSegmentIndex,
		DistanceFromSegmentStartMeters: distanceMeters,
	}
}

func assertVisitedHexes(t *testing.T, got []VisitedHex, want []VisitedHex) {
	t.Helper()

	if len(got) != len(want) {
		t.Fatalf("visit count = %d, want %d", len(got), len(want))
	}
	for index := range want {
		if got[index] != want[index] {
			t.Fatalf("visit %d = %#v, want %#v", index, got[index], want[index])
		}
	}
}
