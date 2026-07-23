package id

import (
	"encoding/json"
	"strings"
	"testing"
)

const validUUID = "7b0f04af-37dd-4a69-9bb1-3d8f0c59c8f8"

func TestUUIDIDConstructors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		constructor func(string) (string, error)
	}{
		{
			name: "player",
			constructor: func(value string) (string, error) {
				id, err := NewPlayerID(value)
				return id.String(), err
			},
		},
		{
			name: "pet",
			constructor: func(value string) (string, error) {
				id, err := NewPetID(value)
				return id.String(), err
			},
		},
		{
			name: "city",
			constructor: func(value string) (string, error) {
				id, err := NewCityID(value)
				return id.String(), err
			},
		},
		{
			name: "walk",
			constructor: func(value string) (string, error) {
				id, err := NewWalkID(value)
				return id.String(), err
			},
		},
		{
			name: "season",
			constructor: func(value string) (string, error) {
				id, err := NewSeasonID(value)
				return id.String(), err
			},
		},
		{
			name: "event",
			constructor: func(value string) (string, error) {
				id, err := NewEventID(value)
				return id.String(), err
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actual, err := test.constructor(validUUID)
			if err != nil {
				t.Fatalf("construct ID: %v", err)
			}
			if actual != validUUID {
				t.Errorf("ID = %q, want %q", actual, validUUID)
			}
		})
	}
}

func TestUUIDIDConstructorsRejectInvalidValues(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		constructor func(string) (string, error)
	}{
		{
			name: "player",
			constructor: func(value string) (string, error) {
				id, err := NewPlayerID(value)
				return id.String(), err
			},
		},
		{
			name: "pet",
			constructor: func(value string) (string, error) {
				id, err := NewPetID(value)
				return id.String(), err
			},
		},
		{
			name: "city",
			constructor: func(value string) (string, error) {
				id, err := NewCityID(value)
				return id.String(), err
			},
		},
		{
			name: "walk",
			constructor: func(value string) (string, error) {
				id, err := NewWalkID(value)
				return id.String(), err
			},
		},
		{
			name: "season",
			constructor: func(value string) (string, error) {
				id, err := NewSeasonID(value)
				return id.String(), err
			},
		},
		{
			name: "event",
			constructor: func(value string) (string, error) {
				id, err := NewEventID(value)
				return id.String(), err
			},
		},
	}

	for _, value := range []string{"not-a-uuid", "00000000-0000-0000-0000-000000000000"} {
		for _, test := range tests {
			t.Run(test.name+"/"+value, func(t *testing.T) {
				t.Parallel()

				_, err := test.constructor(value)
				if err == nil {
					t.Fatal("expected an error")
				}
			})
		}
	}
}

func TestHexID(t *testing.T) {
	t.Parallel()

	id := NewHexID(^uint64(0))
	if id != HexID(^uint64(0)) {
		t.Errorf("ID = %d, want %d", id, ^uint64(0))
	}

	encoded, err := json.Marshal(id)
	if err != nil {
		t.Fatalf("marshal HexID: %v", err)
	}
	if actual, want := string(encoded), `"18446744073709551615"`; actual != want {
		t.Errorf("JSON = %s, want %s", actual, want)
	}
}

func TestVersionConstructors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		constructor func(string) (string, error)
	}{
		{
			name: "rule set",
			constructor: func(value string) (string, error) {
				version, err := NewRuleSetVersion(value)
				return version.String(), err
			},
		},
		{
			name: "engine",
			constructor: func(value string) (string, error) {
				version, err := NewEngineVersion(value)
				return version.String(), err
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name+"/accepts non-semantic version", func(t *testing.T) {
			t.Parallel()

			actual, err := test.constructor("release candidate")
			if err != nil {
				t.Fatalf("construct version: %v", err)
			}
			if actual != "release candidate" {
				t.Errorf("version = %q, want %q", actual, "release candidate")
			}
		})

		for _, value := range []string{"", " \t\n "} {
			t.Run(test.name+"/rejects blank", func(t *testing.T) {
				t.Parallel()

				_, err := test.constructor(value)
				if err == nil {
					t.Fatal("expected an error")
				}
				if !strings.Contains(err.Error(), "must not be blank") {
					t.Errorf("error = %q, want blank-value error", err)
				}
			})
		}
	}
}

func TestUUIDIDsMarshalAsJSONStrings(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		value any
	}{
		{
			name:  "player",
			value: mustPlayerID(t),
		},
		{
			name:  "pet",
			value: mustPetID(t),
		},
		{
			name:  "city",
			value: mustCityID(t),
		},
		{
			name:  "walk",
			value: mustWalkID(t),
		},
		{
			name:  "season",
			value: mustSeasonID(t),
		},
		{
			name:  "event",
			value: mustEventID(t),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			encoded, err := json.Marshal(test.value)
			if err != nil {
				t.Fatalf("marshal ID: %v", err)
			}
			if actual, want := string(encoded), `"`+validUUID+`"`; actual != want {
				t.Errorf("JSON = %s, want %s", actual, want)
			}
		})
	}
}

func mustPlayerID(t *testing.T) PlayerID {
	t.Helper()
	id, err := NewPlayerID(validUUID)
	if err != nil {
		t.Fatalf("construct PlayerID: %v", err)
	}
	return id
}

func mustPetID(t *testing.T) PetID {
	t.Helper()
	id, err := NewPetID(validUUID)
	if err != nil {
		t.Fatalf("construct PetID: %v", err)
	}
	return id
}

func mustCityID(t *testing.T) CityID {
	t.Helper()
	id, err := NewCityID(validUUID)
	if err != nil {
		t.Fatalf("construct CityID: %v", err)
	}
	return id
}

func mustWalkID(t *testing.T) WalkID {
	t.Helper()
	id, err := NewWalkID(validUUID)
	if err != nil {
		t.Fatalf("construct WalkID: %v", err)
	}
	return id
}

func mustSeasonID(t *testing.T) SeasonID {
	t.Helper()
	id, err := NewSeasonID(validUUID)
	if err != nil {
		t.Fatalf("construct SeasonID: %v", err)
	}
	return id
}

func mustEventID(t *testing.T) EventID {
	t.Helper()
	id, err := NewEventID(validUUID)
	if err != nil {
		t.Fatalf("construct EventID: %v", err)
	}
	return id
}
