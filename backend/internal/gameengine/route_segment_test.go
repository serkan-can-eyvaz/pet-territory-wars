package gameengine

import (
	"math"
	"testing"
	"time"
)

func TestSegmentInvalidReason_CanonicalValues(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		reason SegmentInvalidReason
		value  string
	}{
		{name: "none", reason: SegmentInvalidReasonNone, value: ""},
		{name: "non-positive duration", reason: SegmentInvalidReasonNonPositiveDuration, value: "NON_POSITIVE_DURATION"},
		{name: "mock location", reason: SegmentInvalidReasonMockLocation, value: "MOCK_LOCATION"},
		{name: "low accuracy", reason: SegmentInvalidReasonLowAccuracy, value: "LOW_ACCURACY"},
		{name: "location jump", reason: SegmentInvalidReasonLocationJump, value: "LOCATION_JUMP"},
		{name: "impossible speed", reason: SegmentInvalidReasonImpossibleSpeed, value: "IMPOSSIBLE_SPEED"},
		{name: "outside city", reason: SegmentInvalidReasonOutsideCity, value: "OUTSIDE_CITY"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if string(test.reason) != test.value {
				t.Fatalf("reason = %q, want %q", test.reason, test.value)
			}
		})
	}
}

func TestBuildRouteSegments_ConstructsOrderedSegmentsWithoutMutatingInput(t *testing.T) {
	t.Parallel()

	base := time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC)
	points := []LocationPoint{
		{Sequence: 3, Latitude: 41.0, Longitude: 29.0, RecordedAt: base, AccuracyMeters: 5},
		{Sequence: 1, Latitude: 41.0, Longitude: 29.001, RecordedAt: base.Add(time.Minute), AccuracyMeters: 5},
		{Sequence: 2, Latitude: 41.0, Longitude: 29.002, RecordedAt: base.Add(2 * time.Minute), AccuracyMeters: 5},
	}
	original := append([]LocationPoint(nil), points...)
	rules := testWalkRules()

	if segments := buildRouteSegments(nil, rules); len(segments) != 0 {
		t.Fatalf("zero-point segments = %d, want 0", len(segments))
	}
	if segments := buildRouteSegments(points[:1], rules); len(segments) != 0 {
		t.Fatalf("one-point segments = %d, want 0", len(segments))
	}

	segments := buildRouteSegments(points, rules)
	if len(segments) != 2 {
		t.Fatalf("segment count = %d, want 2", len(segments))
	}
	if segments[0].From != points[0] || segments[0].To != points[1] {
		t.Fatal("first segment does not preserve the first adjacent point pair")
	}
	if segments[1].From != points[1] || segments[1].To != points[2] {
		t.Fatal("second segment does not preserve the second adjacent point pair")
	}
	for index := range points {
		if points[index] != original[index] {
			t.Fatalf("route point %d was mutated", index)
		}
	}
}

func TestGeoDistanceMeters_DeterministicHaversineDistance(t *testing.T) {
	t.Parallel()

	if distance := geoDistanceMeters(41, 29, 41, 29); distance != 0 {
		t.Fatalf("identical-coordinate distance = %v, want 0", distance)
	}

	forward := geoDistanceMeters(0, 0, 0, 0.001)
	if math.Abs(forward-111.19492664455875) > 0.000001 {
		t.Fatalf("short-pair distance = %.12f, want approximately 111.194926644559", forward)
	}
	if reverse := geoDistanceMeters(0, 0.001, 0, 0); math.Abs(forward-reverse) > 0.000000001 {
		t.Fatalf("reversed distance = %.12f, want %.12f", reverse, forward)
	}
	if repeated := geoDistanceMeters(0, 0, 0, 0.001); repeated != forward {
		t.Fatalf("repeated distance = %.12f, want %.12f", repeated, forward)
	}
}

func TestCalculateRouteSegment_DurationAndSpeed(t *testing.T) {
	t.Parallel()

	base := time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC)
	from := LocationPoint{Latitude: 0, Longitude: 0, RecordedAt: base, AccuracyMeters: 5}
	to := LocationPoint{Latitude: 0, Longitude: 0.001, RecordedAt: base.Add(2 * time.Second), AccuracyMeters: 5}
	rules := testWalkRules()

	segment := calculateRouteSegment(from, to, rules)
	if segment.Duration != 2*time.Second {
		t.Fatalf("duration = %s, want 2s", segment.Duration)
	}
	if want := segment.DistanceMeters / segment.Duration.Seconds(); segment.SpeedMPS != want {
		t.Fatalf("speed = %.12f, want %.12f", segment.SpeedMPS, want)
	}

	zeroDistance := calculateRouteSegment(from, LocationPoint{Latitude: 0, Longitude: 0, RecordedAt: base.Add(time.Second), AccuracyMeters: 5}, rules)
	if zeroDistance.DistanceMeters != 0 || zeroDistance.SpeedMPS != 0 || !zeroDistance.IsValid {
		t.Fatalf("zero-distance segment = %#v, want valid segment with zero distance and speed", zeroDistance)
	}

	for _, duration := range []time.Duration{0, -time.Second} {
		segment := calculateRouteSegment(from, LocationPoint{Latitude: 0, Longitude: 0.001, RecordedAt: base.Add(duration), AccuracyMeters: 5}, rules)
		if segment.SpeedMPS != 0 || segment.InvalidReason != SegmentInvalidReasonNonPositiveDuration || segment.IsValid {
			t.Fatalf("duration %s produced %#v", duration, segment)
		}
	}
}

func TestCalculateRouteSegment_InvalidReasonsAndThresholds(t *testing.T) {
	t.Parallel()

	base := time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC)
	from := LocationPoint{Latitude: 0, Longitude: 0, RecordedAt: base, AccuracyMeters: 5}
	to := LocationPoint{Latitude: 0, Longitude: 0.001, RecordedAt: base.Add(time.Second), AccuracyMeters: 5}
	distance := geoDistanceMeters(from.Latitude, from.Longitude, to.Latitude, to.Longitude)
	speed := distance / time.Second.Seconds()

	tests := []struct {
		name       string
		from       LocationPoint
		to         LocationPoint
		rules      WalkRules
		wantReason SegmentInvalidReason
	}{
		{name: "valid", from: from, to: to, rules: testWalkRules(), wantReason: SegmentInvalidReasonNone},
		{name: "mock from", from: LocationPoint{Latitude: 0, Longitude: 0, RecordedAt: base, AccuracyMeters: 5, IsMockLocation: true}, to: to, rules: testWalkRules(), wantReason: SegmentInvalidReasonMockLocation},
		{name: "mock to", from: from, to: LocationPoint{Latitude: 0, Longitude: 0.001, RecordedAt: base.Add(time.Second), AccuracyMeters: 5, IsMockLocation: true}, rules: testWalkRules(), wantReason: SegmentInvalidReasonMockLocation},
		{name: "low accuracy from", from: LocationPoint{Latitude: 0, Longitude: 0, RecordedAt: base, AccuracyMeters: 11}, to: to, rules: testWalkRules(), wantReason: SegmentInvalidReasonLowAccuracy},
		{name: "low accuracy to", from: from, to: LocationPoint{Latitude: 0, Longitude: 0.001, RecordedAt: base.Add(time.Second), AccuracyMeters: 11}, rules: testWalkRules(), wantReason: SegmentInvalidReasonLowAccuracy},
		{name: "accuracy exact threshold", from: from, to: LocationPoint{Latitude: 0, Longitude: 0.001, RecordedAt: base.Add(time.Second), AccuracyMeters: 10}, rules: testWalkRules(), wantReason: SegmentInvalidReasonNone},
		{name: "location jump", from: from, to: to, rules: WalkRules{MaxAccuracyMeters: 10, MaxJumpMeters: distance - 0.001, MaxSpeedMPS: speed + 1}, wantReason: SegmentInvalidReasonLocationJump},
		{name: "jump exact threshold", from: from, to: to, rules: WalkRules{MaxAccuracyMeters: 10, MaxJumpMeters: distance, MaxSpeedMPS: speed + 1}, wantReason: SegmentInvalidReasonNone},
		{name: "impossible speed", from: from, to: to, rules: WalkRules{MaxAccuracyMeters: 10, MaxJumpMeters: distance + 1, MaxSpeedMPS: speed - 0.001}, wantReason: SegmentInvalidReasonImpossibleSpeed},
		{name: "speed exact threshold", from: from, to: to, rules: WalkRules{MaxAccuracyMeters: 10, MaxJumpMeters: distance + 1, MaxSpeedMPS: speed}, wantReason: SegmentInvalidReasonNone},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			segment := calculateRouteSegment(test.from, test.to, test.rules)
			if segment.InvalidReason != test.wantReason {
				t.Fatalf("invalid reason = %q, want %q", segment.InvalidReason, test.wantReason)
			}
			if segment.IsValid != (test.wantReason == SegmentInvalidReasonNone) {
				t.Fatalf("IsValid = %t, want %t", segment.IsValid, test.wantReason == SegmentInvalidReasonNone)
			}
		})
	}
}

func TestCalculateRouteSegment_InvalidReasonPrecedence(t *testing.T) {
	t.Parallel()

	base := time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC)
	from := LocationPoint{Latitude: 0, Longitude: 0, RecordedAt: base, AccuracyMeters: 5}
	to := LocationPoint{Latitude: 0, Longitude: 0.001, RecordedAt: base.Add(time.Second), AccuracyMeters: 5}

	tests := []struct {
		name       string
		from       LocationPoint
		to         LocationPoint
		rules      WalkRules
		wantReason SegmentInvalidReason
	}{
		{
			name:       "non-positive duration beats mock location",
			from:       LocationPoint{Latitude: 0, Longitude: 0, RecordedAt: base, AccuracyMeters: 100, IsMockLocation: true},
			to:         LocationPoint{Latitude: 0, Longitude: 0.001, RecordedAt: base, AccuracyMeters: 100, IsMockLocation: true},
			rules:      WalkRules{MaxAccuracyMeters: 1, MaxJumpMeters: 1, MaxSpeedMPS: 1},
			wantReason: SegmentInvalidReasonNonPositiveDuration,
		},
		{
			name:       "mock location beats low accuracy",
			from:       LocationPoint{Latitude: 0, Longitude: 0, RecordedAt: base, AccuracyMeters: 100, IsMockLocation: true},
			to:         to,
			rules:      WalkRules{MaxAccuracyMeters: 1, MaxJumpMeters: 1, MaxSpeedMPS: 1},
			wantReason: SegmentInvalidReasonMockLocation,
		},
		{
			name:       "low accuracy beats location jump",
			from:       LocationPoint{Latitude: 0, Longitude: 0, RecordedAt: base, AccuracyMeters: 100},
			to:         to,
			rules:      WalkRules{MaxAccuracyMeters: 1, MaxJumpMeters: 1, MaxSpeedMPS: 1},
			wantReason: SegmentInvalidReasonLowAccuracy,
		},
		{
			name:       "location jump beats impossible speed",
			from:       from,
			to:         to,
			rules:      WalkRules{MaxAccuracyMeters: 10, MaxJumpMeters: 1, MaxSpeedMPS: 1},
			wantReason: SegmentInvalidReasonLocationJump,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			segment := calculateRouteSegment(test.from, test.to, test.rules)
			if segment.InvalidReason != test.wantReason {
				t.Fatalf("invalid reason = %q, want %q", segment.InvalidReason, test.wantReason)
			}
		})
	}
}

func testWalkRules() WalkRules {
	return WalkRules{
		MaxSpeedMPS:       1000,
		MaxAccuracyMeters: 10,
		MaxJumpMeters:     1000,
	}
}
