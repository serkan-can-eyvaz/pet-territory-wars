package gameengine

import "time"

func calculateWalkValidation(segments []RouteSegment, rules WalkRules) (ValidationResult, WalkMetrics) {
	metrics := WalkMetrics{}
	validSegmentCount := 0
	validDuration := time.Duration(0)
	hasInvalidSegment := false
	hasMockLocation := false

	for _, segment := range segments {
		metrics.TotalDistanceMeters += segment.DistanceMeters
		if !segment.IsValid {
			hasInvalidSegment = true
			if segment.InvalidReason == SegmentInvalidReasonMockLocation {
				hasMockLocation = true
			}
			continue
		}

		validSegmentCount++
		metrics.ValidDistanceMeters += segment.DistanceMeters
		validDuration += segment.Duration
	}
	metrics.ValidDurationSeconds = int(validDuration.Seconds())

	if metrics.TotalDistanceMeters > 0 {
		metrics.ValidRouteRatio = metrics.ValidDistanceMeters / metrics.TotalDistanceMeters
	}

	rejectionReasons := make([]WalkRejectionReason, 0, 5)
	if metrics.ValidDurationSeconds < rules.MinDurationSeconds {
		rejectionReasons = append(rejectionReasons, WalkRejectionReasonTooShortDuration)
	}
	if metrics.ValidDistanceMeters < rules.MinDistanceMeters {
		rejectionReasons = append(rejectionReasons, WalkRejectionReasonTooShortDistance)
	}
	if metrics.ValidRouteRatio < rules.MinValidRouteRatio {
		rejectionReasons = append(rejectionReasons, WalkRejectionReasonLowValidRouteRatio)
	}
	if validSegmentCount == 0 {
		rejectionReasons = append(rejectionReasons, WalkRejectionReasonNoValidSegments)
	}
	if hasMockLocation {
		rejectionReasons = append(rejectionReasons, WalkRejectionReasonMockLocationDetected)
	}

	status := WalkValidationStatusValid
	if validSegmentCount == 0 ||
		metrics.ValidDurationSeconds < rules.MinDurationSeconds ||
		metrics.ValidDistanceMeters < rules.MinDistanceMeters ||
		metrics.ValidRouteRatio < rules.MinValidRouteRatio {
		status = WalkValidationStatusInvalid
	} else if hasInvalidSegment {
		status = WalkValidationStatusPartiallyValid
	}

	return ValidationResult{
		Status:           status,
		RejectionReasons: rejectionReasons,
	}, metrics
}
