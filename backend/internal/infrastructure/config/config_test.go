package config

import (
	"strings"
	"testing"
	"time"
)

func TestLoadDefaults(t *testing.T) {
	clearConfigurationEnvironment(t)
	t.Setenv("DATABASE_URL", "postgres://pet_app:local_password@localhost:5432/pet_territory?sslmode=disable")

	config, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if config.Environment != "local" {
		t.Errorf("Environment = %q, want %q", config.Environment, "local")
	}
	if config.HTTPPort != 8080 {
		t.Errorf("HTTPPort = %d, want %d", config.HTTPPort, 8080)
	}
	if config.LogLevel != "debug" {
		t.Errorf("LogLevel = %q, want %q", config.LogLevel, "debug")
	}
	if config.WalkWorkerConcurrency != 2 {
		t.Errorf("WalkWorkerConcurrency = %d, want %d", config.WalkWorkerConcurrency, 2)
	}
	if config.OutboxWorkerConcurrency != 1 {
		t.Errorf("OutboxWorkerConcurrency = %d, want %d", config.OutboxWorkerConcurrency, 1)
	}
	if config.HTTPReadTimeout != 5*time.Second {
		t.Errorf("HTTPReadTimeout = %s, want %s", config.HTTPReadTimeout, 5*time.Second)
	}
	if config.HTTPWriteTimeout != 10*time.Second {
		t.Errorf("HTTPWriteTimeout = %s, want %s", config.HTTPWriteTimeout, 10*time.Second)
	}
	if config.HTTPShutdownTimeout != 10*time.Second {
		t.Errorf("HTTPShutdownTimeout = %s, want %s", config.HTTPShutdownTimeout, 10*time.Second)
	}
}

func TestLoadUsesEnvironmentValues(t *testing.T) {
	clearConfigurationEnvironment(t)
	t.Setenv("APP_ENV", "staging")
	t.Setenv("HTTP_PORT", "9090")
	t.Setenv("DATABASE_URL", "postgres://example")
	t.Setenv("LOG_LEVEL", "info")
	t.Setenv("WALK_WORKER_CONCURRENCY", "3")
	t.Setenv("OUTBOX_WORKER_CONCURRENCY", "4")
	t.Setenv("HTTP_READ_TIMEOUT", "6s")
	t.Setenv("HTTP_WRITE_TIMEOUT", "11s")
	t.Setenv("HTTP_SHUTDOWN_TIMEOUT", "12s")

	config, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if config.Environment != "staging" {
		t.Errorf("Environment = %q, want %q", config.Environment, "staging")
	}
	if config.HTTPPort != 9090 {
		t.Errorf("HTTPPort = %d, want %d", config.HTTPPort, 9090)
	}
	if config.DatabaseURL != "postgres://example" {
		t.Errorf("DatabaseURL = %q, want %q", config.DatabaseURL, "postgres://example")
	}
	if config.LogLevel != "info" {
		t.Errorf("LogLevel = %q, want %q", config.LogLevel, "info")
	}
	if config.WalkWorkerConcurrency != 3 {
		t.Errorf("WalkWorkerConcurrency = %d, want %d", config.WalkWorkerConcurrency, 3)
	}
	if config.OutboxWorkerConcurrency != 4 {
		t.Errorf("OutboxWorkerConcurrency = %d, want %d", config.OutboxWorkerConcurrency, 4)
	}
	if config.HTTPReadTimeout != 6*time.Second {
		t.Errorf("HTTPReadTimeout = %s, want %s", config.HTTPReadTimeout, 6*time.Second)
	}
	if config.HTTPWriteTimeout != 11*time.Second {
		t.Errorf("HTTPWriteTimeout = %s, want %s", config.HTTPWriteTimeout, 11*time.Second)
	}
	if config.HTTPShutdownTimeout != 12*time.Second {
		t.Errorf("HTTPShutdownTimeout = %s, want %s", config.HTTPShutdownTimeout, 12*time.Second)
	}
}

func TestLoadRejectsInvalidConfiguration(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		value     string
		errString string
	}{
		{
			name:      "unsupported environment",
			key:       "APP_ENV",
			value:     "development",
			errString: "APP_ENV",
		},
		{
			name:      "empty database URL",
			key:       "DATABASE_URL",
			value:     "",
			errString: "DATABASE_URL",
		},
		{
			name:      "invalid HTTP port",
			key:       "HTTP_PORT",
			value:     "invalid",
			errString: "HTTP_PORT",
		},
		{
			name:      "out of range HTTP port",
			key:       "HTTP_PORT",
			value:     "65536",
			errString: "HTTP_PORT",
		},
		{
			name:      "zero HTTP port",
			key:       "HTTP_PORT",
			value:     "0",
			errString: "HTTP_PORT",
		},
		{
			name:      "negative walk worker concurrency",
			key:       "WALK_WORKER_CONCURRENCY",
			value:     "-1",
			errString: "WALK_WORKER_CONCURRENCY",
		},
		{
			name:      "negative outbox worker concurrency",
			key:       "OUTBOX_WORKER_CONCURRENCY",
			value:     "-1",
			errString: "OUTBOX_WORKER_CONCURRENCY",
		},
		{
			name:      "invalid read timeout",
			key:       "HTTP_READ_TIMEOUT",
			value:     "invalid",
			errString: "HTTP_READ_TIMEOUT",
		},
		{
			name:      "invalid write timeout",
			key:       "HTTP_WRITE_TIMEOUT",
			value:     "invalid",
			errString: "HTTP_WRITE_TIMEOUT",
		},
		{
			name:      "invalid shutdown timeout",
			key:       "HTTP_SHUTDOWN_TIMEOUT",
			value:     "invalid",
			errString: "HTTP_SHUTDOWN_TIMEOUT",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			clearConfigurationEnvironment(t)
			t.Setenv("DATABASE_URL", "postgres://example")
			t.Setenv(test.key, test.value)

			_, err := Load()
			if err == nil {
				t.Fatal("Load() error = nil, want error")
			}
			if !strings.Contains(err.Error(), test.errString) {
				t.Errorf("Load() error = %q, want error containing %q", err, test.errString)
			}
		})
	}
}

func clearConfigurationEnvironment(t *testing.T) {
	t.Helper()

	for _, key := range []string{
		"APP_ENV",
		"HTTP_PORT",
		"DATABASE_URL",
		"LOG_LEVEL",
		"WALK_WORKER_CONCURRENCY",
		"OUTBOX_WORKER_CONCURRENCY",
		"HTTP_READ_TIMEOUT",
		"HTTP_WRITE_TIMEOUT",
		"HTTP_SHUTDOWN_TIMEOUT",
	} {
		t.Setenv(key, "")
	}
}
