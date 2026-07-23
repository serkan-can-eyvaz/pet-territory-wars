// Package id defines immutable domain identifier value objects.
package id

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// PlayerID identifies a player.
type PlayerID uuid.UUID

// NewPlayerID constructs a PlayerID from a non-nil UUID string.
func NewPlayerID(value string) (PlayerID, error) {
	parsed, err := parseUUID(value, "player ID")
	if err != nil {
		return PlayerID{}, err
	}

	return PlayerID(parsed), nil
}

// String returns the canonical UUID representation.
func (id PlayerID) String() string {
	return uuid.UUID(id).String()
}

// MarshalJSON encodes the ID as a JSON string.
func (id PlayerID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

// PetID identifies a pet.
type PetID uuid.UUID

// NewPetID constructs a PetID from a non-nil UUID string.
func NewPetID(value string) (PetID, error) {
	parsed, err := parseUUID(value, "pet ID")
	if err != nil {
		return PetID{}, err
	}

	return PetID(parsed), nil
}

// String returns the canonical UUID representation.
func (id PetID) String() string {
	return uuid.UUID(id).String()
}

// MarshalJSON encodes the ID as a JSON string.
func (id PetID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

// CityID identifies a city.
type CityID uuid.UUID

// NewCityID constructs a CityID from a non-nil UUID string.
func NewCityID(value string) (CityID, error) {
	parsed, err := parseUUID(value, "city ID")
	if err != nil {
		return CityID{}, err
	}

	return CityID(parsed), nil
}

// String returns the canonical UUID representation.
func (id CityID) String() string {
	return uuid.UUID(id).String()
}

// MarshalJSON encodes the ID as a JSON string.
func (id CityID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

// WalkID identifies a walk.
type WalkID uuid.UUID

// NewWalkID constructs a WalkID from a non-nil UUID string.
func NewWalkID(value string) (WalkID, error) {
	parsed, err := parseUUID(value, "walk ID")
	if err != nil {
		return WalkID{}, err
	}

	return WalkID(parsed), nil
}

// String returns the canonical UUID representation.
func (id WalkID) String() string {
	return uuid.UUID(id).String()
}

// MarshalJSON encodes the ID as a JSON string.
func (id WalkID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

// SeasonID identifies a season.
type SeasonID uuid.UUID

// NewSeasonID constructs a SeasonID from a non-nil UUID string.
func NewSeasonID(value string) (SeasonID, error) {
	parsed, err := parseUUID(value, "season ID")
	if err != nil {
		return SeasonID{}, err
	}

	return SeasonID(parsed), nil
}

// String returns the canonical UUID representation.
func (id SeasonID) String() string {
	return uuid.UUID(id).String()
}

// MarshalJSON encodes the ID as a JSON string.
func (id SeasonID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

// EventID identifies a domain event.
type EventID uuid.UUID

// NewEventID constructs an EventID from a non-nil UUID string.
func NewEventID(value string) (EventID, error) {
	parsed, err := parseUUID(value, "event ID")
	if err != nil {
		return EventID{}, err
	}

	return EventID(parsed), nil
}

// String returns the canonical UUID representation.
func (id EventID) String() string {
	return uuid.UUID(id).String()
}

// MarshalJSON encodes the ID as a JSON string.
func (id EventID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

// HexID identifies a territory hexagon without exposing an H3 dependency.
type HexID uint64

// NewHexID constructs a HexID from its domain value.
func NewHexID(value uint64) HexID {
	return HexID(value)
}

// String returns the decimal representation used in JSON.
func (id HexID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

// MarshalJSON encodes the ID as a JSON string to avoid numeric precision loss.
func (id HexID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

// RuleSetVersion identifies an immutable version of a rule set.
type RuleSetVersion string

// NewRuleSetVersion constructs a RuleSetVersion from a non-blank value.
func NewRuleSetVersion(value string) (RuleSetVersion, error) {
	if strings.TrimSpace(value) == "" {
		return "", fmt.Errorf("rule set version must not be blank")
	}

	return RuleSetVersion(value), nil
}

// String returns the version value.
func (version RuleSetVersion) String() string {
	return string(version)
}

// MarshalJSON encodes the version as a JSON string.
func (version RuleSetVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(version.String())
}

// EngineVersion identifies an immutable version of the game engine.
type EngineVersion string

// NewEngineVersion constructs an EngineVersion from a non-blank value.
func NewEngineVersion(value string) (EngineVersion, error) {
	if strings.TrimSpace(value) == "" {
		return "", fmt.Errorf("engine version must not be blank")
	}

	return EngineVersion(value), nil
}

// String returns the version value.
func (version EngineVersion) String() string {
	return string(version)
}

// MarshalJSON encodes the version as a JSON string.
func (version EngineVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(version.String())
}

func parseUUID(value, name string) (uuid.UUID, error) {
	parsed, err := uuid.Parse(value)
	if err != nil {
		return uuid.Nil, fmt.Errorf("parse %s: %w", name, err)
	}
	if parsed == uuid.Nil {
		return uuid.Nil, fmt.Errorf("%s must not be nil", name)
	}

	return parsed, nil
}
