package gameengine

import "github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/domain/id"

// HexResolver resolves coordinates to deterministic H3 hex identifiers.
// The same latitude, longitude, and resolution must always resolve to the same id.HexID.
type HexResolver interface {
	ResolveHex(
		latitude float64,
		longitude float64,
		resolution int,
	) (id.HexID, error)
}

func resolveInterpolationHexes(
	points []interpolatedPoint,
	resolution int,
	resolver HexResolver,
) ([]id.HexID, error) {
	hexIDs := make([]id.HexID, 0, len(points))
	for _, point := range points {
		hexID, err := resolver.ResolveHex(point.Latitude, point.Longitude, resolution)
		if err != nil {
			return nil, err
		}
		hexIDs = append(hexIDs, hexID)
	}

	return hexIDs, nil
}
