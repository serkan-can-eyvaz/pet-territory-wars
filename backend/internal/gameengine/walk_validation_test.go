package gameengine

import (
	"testing"
	"time"
)

func TestCalculateWalkValidation_ValidWalk(t *testing.T) {
	t.Parallel()

	validation, metrics := calculateWalkValidation([]RouteSegment{
		validRouteSegment(100, 60*time.Second),
		validRouteSegment(50, 30*time.Second),
	}, WalkRules{
		MinDurationSeconds: 90,
		MinDistanceMeters:  150,
		MinValidRouteRatio: 1,
	})

	if validation.Status != WalkValidationStatusValid {
		t.Fatalf("status = %q, want %q", validation.Status, WalkValidationStatusValid)
	}
	if len(validation.RejectionReasons) != 0 {
		t.Fatalf("rejection reasons = %v, want none", validation.RejectionReasons)
	}
	if metrics.TotalDistanceMeters != 150 || metrics.ValidDistanceMeters != 150 || metrics.ValidDurationSeconds != 90 || metrics.ValidRouteRatio != 1 {
		t.Fatalf("metrics = %#v, want total 150, valid 150, duration 90, ratio 1", metrics)
	}
}

func TestCalculateWalkValidation_InvalidWalkUsesCanonicalRejectionOrder(t *testing.T) {
	t.Parallel()

	validation, metrics := calculateWalkValidation([]RouteSegment{
		validRouteSegment(50, 30*time.Second),
		{
			DistanceMeters: 100,
			Duration:       30 * time.Second,
			IsValid:        false,
			InvalidReason:  SegmentInvalidReasonMockLocation,
		},
	}, WalkRules{
		MinDurationSeconds: 60,
		MinDistanceMeters:  100,
		MinValidRouteRatio: 0.5,
	})

	if validation.Status != WalkValidationStatusInvalid {
		t.Fatalf("status = %q, want %q", validation.Status, WalkValidationStatusInvalid)
	}
	wantReasons := []WalkRejectionReason{
		WalkRejectionReasonTooShortDuration,
		WalkRejectionReasonTooShortDistance,
		WalkRejectionReasonLowValidRouteRatio,
		WalkRejectionReasonMockLocationDetected,
	}
	assertRejectionReasons(t, validation.RejectionReasons, wantReasons)
	if metrics.TotalDistanceMeters != 150 || metrics.ValidDistanceMeters != 50 || metrics.ValidDurationSeconds != 30 || metrics.ValidRouteRatio != 1.0/3.0 {
		t.Fatalf("metrics = %#v, want total 150, valid 50, duration 30, ratio 1/3", metrics)
	}
}

func TestCalculateWalkValidation_PartiallyValidWalk(t *testing.T) {
	t.Parallel()

	validation, metrics := calculateWalkValidation([]RouteSegment{
		validRouteSegment(100, time.Minute),
		{
			DistanceMeters: 100,
			Duration:       time.Minute,
			IsValid:        false,
			InvalidReason:  SegmentInvalidReasonMockLocation,
		},
	}, WalkRules{
		MinDurationSeconds: 60,
		MinDistanceMeters:  100,
		MinValidRouteRatio: 0.5,
	})

	if validation.Status != WalkValidationStatusPartiallyValid {
		t.Fatalf("status = %q, want %q", validation.Status, WalkValidationStatusPartiallyValid)
	}
	assertRejectionReasons(t, validation.RejectionReasons, []WalkRejectionReason{WalkRejectionReasonMockLocationDetected})
	if metrics.TotalDistanceMeters != 200 || metrics.ValidDistanceMeters != 100 || metrics.ValidDurationSeconds != 60 || metrics.ValidRouteRatio != 0.5 {
		t.Fatalf("metrics = %#v, want total 200, valid 100, duration 60, ratio 0.5", metrics)
	}
}

func TestCalculateWalkValidation_NoValidSegments(t *testing.T) {
	t.Parallel()

	validation, metrics := calculateWalkValidation([]RouteSegment{{
		DistanceMeters: 0,
		Duration:       0,
		IsValid:        false,
		InvalidReason:  SegmentInvalidReasonNonPositiveDuration,
	}}, WalkRules{})

	if validation.Status != WalkValidationStatusInvalid {
		t.Fatalf("status = %q, want %q", validation.Status, WalkValidationStatusInvalid)
	}
	assertRejectionReasons(t, validation.RejectionReasons, []WalkRejectionReason{WalkRejectionReasonNoValidSegments})
	if metrics.TotalDistanceMeters != 0 || metrics.ValidDistanceMeters != 0 || metrics.ValidDurationSeconds != 0 || metrics.ValidRouteRatio != 0 {
		t.Fatalf("metrics = %#v, want zero metrics", metrics)
	}
}

func TestCalculateWalkValidation_ZeroTotalDistanceHasZeroRatio(t *testing.T) {
	t.Parallel()

	validation, metrics := calculateWalkValidation([]RouteSegment{
		validRouteSegment(0, 30*time.Second),
	}, WalkRules{
		MinDurationSeconds: 30,
		MinDistanceMeters:  0,
		MinValidRouteRatio: 0,
	})

	if validation.Status != WalkValidationStatusValid {
		t.Fatalf("status = %q, want %q", validation.Status, WalkValidationStatusValid)
	}
	if metrics.ValidRouteRatio != 0 {
		t.Fatalf("valid route ratio = %v, want 0", metrics.ValidRouteRatio)
	}
}

func TestCalculateWalkValidation_IsDeterministic(t *testing.T) {
	t.Parallel()

	segments := []RouteSegment{
		validRouteSegment(100, time.Minute),
		{
			DistanceMeters: 50,
			Duration:       time.Minute,
			IsValid:        false,
			InvalidReason:  SegmentInvalidReasonLocationJump,
		},
	}
	rules := WalkRules{MinDurationSeconds: 60, MinDistanceMeters: 100, MinValidRouteRatio: 0.5}

	firstValidation, firstMetrics := calculateWalkValidation(segments, rules)
	secondValidation, secondMetrics := calculateWalkValidation(segments, rules)
	if firstValidation.Status != secondValidation.Status || firstMetrics != secondMetrics {
		t.Fatalf("repeated calculation changed result: first (%#v, %#v), second (%#v, %#v)", firstValidation, firstMetrics, secondValidation, secondMetrics)
	}
	assertRejectionReasons(t, secondValidation.RejectionReasons, firstValidation.RejectionReasons)
}

func validRouteSegment(distanceMeters float64, duration time.Duration) RouteSegment {
	return RouteSegment{
		DistanceMeters: distanceMeters,
		Duration:       duration,
		IsValid:        true,
		InvalidReason:  SegmentInvalidReasonNone,
	}
}

func assertRejectionReasons(t *testing.T, got []WalkRejectionReason, want []WalkRejectionReason) {
	t.Helper()

	if len(got) != len(want) {
		t.Fatalf("rejection reason count = %d, want %d (%v)", len(got), len(want), want)
	}
	for index := range want {
		if got[index] != want[index] {
			t.Fatalf("rejection reason %d = %q, want %q", index, got[index], want[index])
		}
	}
}
