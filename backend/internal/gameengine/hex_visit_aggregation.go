package gameengine

import (
	"sort"
	"time"

	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/domain/id"
)

type hexRouteSample struct {
	HexID                          id.HexID
	RecordedAt                     time.Time
	SourceSegmentIndex             int
	DistanceFromSegmentStartMeters float64
}

func aggregateVisitedHexes(samples []hexRouteSample) []VisitedHex {
	visitsByHexID := make(map[id.HexID]*VisitedHex, len(samples))
	presenceDurations := make(map[id.HexID]time.Duration, len(samples))

	for index, sample := range samples {
		visit, exists := visitsByHexID[sample.HexID]
		if !exists {
			visit = &VisitedHex{
				HexID:          sample.HexID,
				FirstEnteredAt: sample.RecordedAt,
				LastExitedAt:   sample.RecordedAt,
			}
			visitsByHexID[sample.HexID] = visit
		}

		if sample.RecordedAt.After(visit.LastExitedAt) {
			visit.LastExitedAt = sample.RecordedAt
		}

		if index == 0 ||
			samples[index-1].SourceSegmentIndex != sample.SourceSegmentIndex ||
			samples[index-1].HexID != sample.HexID {
			visit.EntryCount++
		}

		if index == 0 || samples[index-1].SourceSegmentIndex != sample.SourceSegmentIndex {
			continue
		}

		previousSample := samples[index-1]
		previousVisit := visitsByHexID[previousSample.HexID]
		previousVisit.DistanceMeters += sample.DistanceFromSegmentStartMeters - previousSample.DistanceFromSegmentStartMeters
		presenceDurations[previousSample.HexID] += sample.RecordedAt.Sub(previousSample.RecordedAt)
		if sample.RecordedAt.After(previousVisit.LastExitedAt) {
			previousVisit.LastExitedAt = sample.RecordedAt
		}
	}

	visits := make([]VisitedHex, 0, len(visitsByHexID))
	for hexID, visit := range visitsByHexID {
		visit.PresenceSeconds = int(presenceDurations[hexID].Seconds())
		visits = append(visits, *visit)
	}

	sort.Slice(visits, func(left int, right int) bool {
		if visits[left].FirstEnteredAt.Equal(visits[right].FirstEnteredAt) {
			return visits[left].HexID.String() < visits[right].HexID.String()
		}
		return visits[left].FirstEnteredAt.Before(visits[right].FirstEnteredAt)
	})

	return visits
}
