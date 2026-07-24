package gameengine

type interpolatedPoint struct {
	Latitude  float64
	Longitude float64
}

func interpolateRoute(segments []RouteSegment, rules MovementRules) []interpolatedPoint {
	var points []interpolatedPoint
	var previousValidSegment RouteSegment
	hasPreviousValidSegment := false

	for _, segment := range segments {
		if !segment.IsValid {
			hasPreviousValidSegment = false
			continue
		}

		segmentPoints := interpolateSegment(segment, rules.InterpolationMeters)
		if hasPreviousValidSegment && previousValidSegment.To == segment.From {
			points = append(points, segmentPoints[1:]...)
		} else {
			points = append(points, segmentPoints...)
		}

		previousValidSegment = segment
		hasPreviousValidSegment = true
	}

	return points
}

func interpolateSegment(segment RouteSegment, interpolationMeters float64) []interpolatedPoint {
	points := []interpolatedPoint{
		{Latitude: segment.From.Latitude, Longitude: segment.From.Longitude},
	}

	for index := 1; ; index++ {
		distanceMeters := float64(index) * interpolationMeters
		if distanceMeters >= segment.DistanceMeters {
			break
		}

		ratio := distanceMeters / segment.DistanceMeters
		points = append(points, interpolatedPoint{
			Latitude:  segment.From.Latitude + (segment.To.Latitude-segment.From.Latitude)*ratio,
			Longitude: segment.From.Longitude + (segment.To.Longitude-segment.From.Longitude)*ratio,
		})
	}

	return append(points, interpolatedPoint{
		Latitude:  segment.To.Latitude,
		Longitude: segment.To.Longitude,
	})
}
