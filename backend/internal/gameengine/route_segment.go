package gameengine

import (
	"math"
	"time"
)

const earthRadiusMeters = 6371000.0

// SegmentInvalidReason describes why a route segment cannot be used by the engine.
type SegmentInvalidReason string

const (
	SegmentInvalidReasonNone                SegmentInvalidReason = ""
	SegmentInvalidReasonNonPositiveDuration SegmentInvalidReason = "NON_POSITIVE_DURATION"
	SegmentInvalidReasonMockLocation        SegmentInvalidReason = "MOCK_LOCATION"
	SegmentInvalidReasonLowAccuracy         SegmentInvalidReason = "LOW_ACCURACY"
	SegmentInvalidReasonLocationJump        SegmentInvalidReason = "LOCATION_JUMP"
	SegmentInvalidReasonImpossibleSpeed     SegmentInvalidReason = "IMPOSSIBLE_SPEED"
	SegmentInvalidReasonOutsideCity         SegmentInvalidReason = "OUTSIDE_CITY"
)

// RouteSegment represents one adjacent pair of recorded route points.
type RouteSegment struct {
	From           LocationPoint
	To             LocationPoint
	DistanceMeters float64
	Duration       time.Duration
	SpeedMPS       float64
	IsValid        bool
	InvalidReason  SegmentInvalidReason
}

func buildRouteSegments(route []LocationPoint, rules WalkRules) []RouteSegment {
	if len(route) < 2 {
		return nil
	}

	segments := make([]RouteSegment, 0, len(route)-1)
	for index := 0; index < len(route)-1; index++ {
		segments = append(segments, calculateRouteSegment(route[index], route[index+1], rules))
	}

	return segments
}

func calculateRouteSegment(from LocationPoint, to LocationPoint, rules WalkRules) RouteSegment {
	duration := to.RecordedAt.Sub(from.RecordedAt)
	distanceMeters := geoDistanceMeters(from.Latitude, from.Longitude, to.Latitude, to.Longitude)
	speedMPS := 0.0
	if duration > 0 {
		speedMPS = distanceMeters / duration.Seconds()
	}

	invalidReason := SegmentInvalidReasonNone
	switch {
	case duration <= 0:
		invalidReason = SegmentInvalidReasonNonPositiveDuration
	case from.IsMockLocation || to.IsMockLocation:
		invalidReason = SegmentInvalidReasonMockLocation
	case from.AccuracyMeters > rules.MaxAccuracyMeters || to.AccuracyMeters > rules.MaxAccuracyMeters:
		invalidReason = SegmentInvalidReasonLowAccuracy
	case distanceMeters > rules.MaxJumpMeters:
		invalidReason = SegmentInvalidReasonLocationJump
	case speedMPS > rules.MaxSpeedMPS:
		invalidReason = SegmentInvalidReasonImpossibleSpeed
	}

	return RouteSegment{
		From:           from,
		To:             to,
		DistanceMeters: distanceMeters,
		Duration:       duration,
		SpeedMPS:       speedMPS,
		IsValid:        invalidReason == SegmentInvalidReasonNone,
		InvalidReason:  invalidReason,
	}
}

func geoDistanceMeters(fromLatitude float64, fromLongitude float64, toLatitude float64, toLongitude float64) float64 {
	lat1 := fromLatitude * math.Pi / 180
	lat2 := toLatitude * math.Pi / 180
	deltaLat := (toLatitude - fromLatitude) * math.Pi / 180
	deltaLon := (toLongitude - fromLongitude) * math.Pi / 180

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	a = math.Max(0, math.Min(1, a))

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return earthRadiusMeters * c
}
