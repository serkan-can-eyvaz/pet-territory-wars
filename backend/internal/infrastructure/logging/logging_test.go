package logging

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestNewWritesJSONWithRequiredFields(t *testing.T) {
	var output bytes.Buffer

	logger, err := New(&output, "api", "info")
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	logger.Info("api starting")

	var record map[string]any
	if err := json.Unmarshal(output.Bytes(), &record); err != nil {
		t.Fatalf("decode JSON log: %v", err)
	}

	for _, key := range []string{"time", "level", "msg", "service"} {
		if _, ok := record[key]; !ok {
			t.Errorf("log record missing %q", key)
		}
	}
	if record["level"] != "INFO" {
		t.Errorf("level = %q, want %q", record["level"], "INFO")
	}
	if record["msg"] != "api starting" {
		t.Errorf("msg = %q, want %q", record["msg"], "api starting")
	}
	if record["service"] != "api" {
		t.Errorf("service = %q, want %q", record["service"], "api")
	}
}

func TestNewAppliesConfiguredLevel(t *testing.T) {
	var output bytes.Buffer

	logger, err := New(&output, "api", "warn")
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	logger.Debug("debug message")
	logger.Info("info message")
	if output.Len() != 0 {
		t.Errorf("output contains records below warn level: %q", output.String())
	}

	logger.Warn("warn message")
	if output.Len() == 0 {
		t.Error("output is empty after warn log")
	}
}

func TestNewRejectsUnsupportedLevel(t *testing.T) {
	logger, err := New(&bytes.Buffer{}, "api", "verbose")
	if err == nil {
		t.Fatal("New() error = nil, want error")
	}
	if logger != nil {
		t.Error("New() logger is not nil on error")
	}
}
