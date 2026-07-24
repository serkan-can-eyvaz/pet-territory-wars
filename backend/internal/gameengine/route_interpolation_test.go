package gameengine

import (
	"testing"
	"time"
)

func TestInterpolateSegment_ShortSegmentPreservesEndpoints(t *testing.T) {
	t.Parallel()

	segment := interpolationSegment(interpolationPoint(1, 10, 20), interpolationPoint(2, 30, 40), 50)
	points := interpolateSegment(segment, 100)

	assertInterpolatedPoints(t, points, []interpolatedPoint{
		{Latitude: 10, Longitude: 20},
		{Latitude: 30, Longitude: 40},
	})
}

func TestInterpolateSegment_LongSegmentAddsOrderedIntermediatePoints(t *testing.T) {
	t.Parallel()

	segment := interpolationSegment(interpolationPoint(1, 0, 0), interpolationPoint(2, 10, 20), 250)
	points := interpolateSegment(segment, 100)

	assertInterpolatedPoints(t, points, []interpolatedPoint{
		{Latitude: 0, Longitude: 0},
		{Latitude: 4, Longitude: 8},
		{Latitude: 8, Longitude: 16},
		{Latitude: 10, Longitude: 20},
	})
}

func TestInterpolateRoute_SkipsInvalidSegmentsWithoutConnectingValidParts(t *testing.T) {
	t.Parallel()

	first := interpolationSegment(interpolationPoint(1, 0, 0), interpolationPoint(2, 0, 1), 100)
	invalid := interpolationSegment(interpolationPoint(2, 0, 1), interpolationPoint(3, 0, 2), 100)
	invalid.IsValid = false
	last := interpolationSegment(interpolationPoint(3, 0, 2), interpolationPoint(4, 0, 3), 100)

	points := interpolateRoute([]RouteSegment{first, invalid, last}, MovementRules{InterpolationMeters: 50})

	assertInterpolatedPoints(t, points, []interpolatedPoint{
		{Latitude: 0, Longitude: 0},
		{Latitude: 0, Longitude: 0.5},
		{Latitude: 0, Longitude: 1},
		{Latitude: 0, Longitude: 2},
		{Latitude: 0, Longitude: 2.5},
		{Latitude: 0, Longitude: 3},
	})
}

func TestInterpolateRoute_DeduplicatesOnlyDirectSourceContinuation(t *testing.T) {
	t.Parallel()

	first := interpolationSegment(interpolationPoint(1, 0, 0), interpolationPoint(2, 0, 1), 100)
	directContinuation := interpolationSegment(interpolationPoint(2, 0, 1), interpolationPoint(3, 0, 2), 100)
	notDirectContinuation := interpolationSegment(interpolationPoint(4, 0, 2), interpolationPoint(5, 0, 3), 100)

	points := interpolateRoute(
		[]RouteSegment{first, directContinuation, notDirectContinuation},
		MovementRules{InterpolationMeters: 100},
	)

	assertInterpolatedPoints(t, points, []interpolatedPoint{
		{Latitude: 0, Longitude: 0},
		{Latitude: 0, Longitude: 1},
		{Latitude: 0, Longitude: 2},
		{Latitude: 0, Longitude: 2},
		{Latitude: 0, Longitude: 3},
	})
}

func TestInterpolateRoute_IsDeterministicAndDoesNotMutateSegments(t *testing.T) {
	t.Parallel()

	segments := []RouteSegment{
		interpolationSegment(interpolationPoint(1, 0, 0), interpolationPoint(2, 0, 1), 100),
		interpolationSegment(interpolationPoint(2, 0, 1), interpolationPoint(3, 0, 2), 100),
	}
	original := append([]RouteSegment(nil), segments...)
	rules := MovementRules{InterpolationMeters: 25}

	first := interpolateRoute(segments, rules)
	second := interpolateRoute(segments, rules)
	assertInterpolatedPoints(t, second, first)
	for index := range segments {
		if segments[index] != original[index] {
			t.Fatalf("segment %d was mutated", index)
		}
	}
}

func interpolationPoint(sequence int, latitude float64, longitude float64) LocationPoint {
	return LocationPoint{
		Sequence:   sequence,
		Latitude:   latitude,
		Longitude:  longitude,
		RecordedAt: time.Date(2026, time.January, 1, 12, 0, sequence, 0, time.UTC),
	}
}

func interpolationSegment(from LocationPoint, to LocationPoint, distanceMeters float64) RouteSegment {
	return RouteSegment{
		From:           from,
		To:             to,
		DistanceMeters: distanceMeters,
		IsValid:        true,
		InvalidReason:  SegmentInvalidReasonNone,
	}
}

func assertInterpolatedPoints(t *testing.T, got []interpolatedPoint, want []interpolatedPoint) {
	t.Helper()

	if len(got) != len(want) {
		t.Fatalf("point count = %d, want %d (%v)", len(got), len(want), want)
	}
	for index := range want {
		if got[index] != want[index] {
			t.Fatalf("point %d = %#v, want %#v", index, got[index], want[index])
		}
	}
}
