package gameengine

import (
	"errors"
	"testing"

	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/domain/id"
)

func TestResolveInterpolationHexes_EmptyInput(t *testing.T) {
	t.Parallel()

	resolver := &fakeHexResolver{}
	hexIDs, err := resolveInterpolationHexes(nil, 9, resolver)
	if err != nil {
		t.Fatalf("resolve interpolation hexes: %v", err)
	}
	if len(hexIDs) != 0 {
		t.Fatalf("hex ID count = %d, want 0", len(hexIDs))
	}
	if len(resolver.calls) != 0 {
		t.Fatalf("resolver calls = %d, want 0", len(resolver.calls))
	}
}

func TestResolveInterpolationHexes_ResolvesPointsInInputOrder(t *testing.T) {
	t.Parallel()

	points := []interpolatedPoint{
		{Latitude: 41.0082, Longitude: 28.9784},
		{Latitude: 41.0151, Longitude: 28.9795},
		{Latitude: 41.0220, Longitude: 28.9806},
	}
	resolver := &fakeHexResolver{hexIDs: []id.HexID{101, 202, 303}}

	hexIDs, err := resolveInterpolationHexes(points, 9, resolver)
	if err != nil {
		t.Fatalf("resolve interpolation hexes: %v", err)
	}
	assertHexIDs(t, hexIDs, []id.HexID{101, 202, 303})
	assertResolverCalls(t, resolver.calls, []resolverCall{
		{latitude: 41.0082, longitude: 28.9784, resolution: 9},
		{latitude: 41.0151, longitude: 28.9795, resolution: 9},
		{latitude: 41.0220, longitude: 28.9806, resolution: 9},
	})
}

func TestResolveInterpolationHexes_SinglePointDoesNotMutateInput(t *testing.T) {
	t.Parallel()

	points := []interpolatedPoint{{Latitude: 41.0082, Longitude: 28.9784}}
	original := append([]interpolatedPoint(nil), points...)
	resolver := &fakeHexResolver{hexIDs: []id.HexID{101}}

	hexIDs, err := resolveInterpolationHexes(points, 8, resolver)
	if err != nil {
		t.Fatalf("resolve interpolation hexes: %v", err)
	}
	assertHexIDs(t, hexIDs, []id.HexID{101})
	if points[0] != original[0] {
		t.Fatal("input point was mutated")
	}
}

func TestResolveInterpolationHexes_StopsAtFirstResolverError(t *testing.T) {
	t.Parallel()

	resolverError := errors.New("resolver unavailable")
	resolver := &fakeHexResolver{
		hexIDs: []id.HexID{101},
		errAt:  2,
		err:    resolverError,
	}
	points := []interpolatedPoint{
		{Latitude: 41.0082, Longitude: 28.9784},
		{Latitude: 41.0151, Longitude: 28.9795},
		{Latitude: 41.0220, Longitude: 28.9806},
	}

	hexIDs, err := resolveInterpolationHexes(points, 9, resolver)
	if !errors.Is(err, resolverError) {
		t.Fatalf("error = %v, want resolver error", err)
	}
	if hexIDs != nil {
		t.Fatalf("hex IDs = %v, want no partial result", hexIDs)
	}
	assertResolverCalls(t, resolver.calls, []resolverCall{
		{latitude: 41.0082, longitude: 28.9784, resolution: 9},
		{latitude: 41.0151, longitude: 28.9795, resolution: 9},
	})
}

func TestResolveInterpolationHexes_IsDeterministic(t *testing.T) {
	t.Parallel()

	points := []interpolatedPoint{
		{Latitude: 41.0082, Longitude: 28.9784},
		{Latitude: 41.0151, Longitude: 28.9795},
	}
	rulesResolution := 9

	first, err := resolveInterpolationHexes(points, rulesResolution, &fakeHexResolver{hexIDs: []id.HexID{101, 202}})
	if err != nil {
		t.Fatalf("first resolution: %v", err)
	}
	second, err := resolveInterpolationHexes(points, rulesResolution, &fakeHexResolver{hexIDs: []id.HexID{101, 202}})
	if err != nil {
		t.Fatalf("second resolution: %v", err)
	}
	assertHexIDs(t, second, first)
}

type resolverCall struct {
	latitude   float64
	longitude  float64
	resolution int
}

type fakeHexResolver struct {
	hexIDs []id.HexID
	errAt  int
	err    error
	calls  []resolverCall
}

func (resolver *fakeHexResolver) ResolveHex(latitude float64, longitude float64, resolution int) (id.HexID, error) {
	resolver.calls = append(resolver.calls, resolverCall{
		latitude:   latitude,
		longitude:  longitude,
		resolution: resolution,
	})
	if resolver.errAt == len(resolver.calls) {
		return 0, resolver.err
	}

	return resolver.hexIDs[len(resolver.calls)-1], nil
}

func assertHexIDs(t *testing.T, got []id.HexID, want []id.HexID) {
	t.Helper()

	if len(got) != len(want) {
		t.Fatalf("hex ID count = %d, want %d", len(got), len(want))
	}
	for index := range want {
		if got[index] != want[index] {
			t.Fatalf("hex ID %d = %d, want %d", index, got[index], want[index])
		}
	}
}

func assertResolverCalls(t *testing.T, got []resolverCall, want []resolverCall) {
	t.Helper()

	if len(got) != len(want) {
		t.Fatalf("resolver call count = %d, want %d", len(got), len(want))
	}
	for index := range want {
		if got[index] != want[index] {
			t.Fatalf("resolver call %d = %#v, want %#v", index, got[index], want[index])
		}
	}
}
