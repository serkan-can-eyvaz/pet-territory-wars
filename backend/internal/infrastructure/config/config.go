package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	defaultEnvironment             = "local"
	defaultHTTPPort                = 8080
	defaultLogLevel                = "debug"
	defaultWalkWorkerConcurrency   = 2
	defaultOutboxWorkerConcurrency = 1
	defaultHTTPReadTimeout         = 5 * time.Second
	defaultHTTPWriteTimeout        = 10 * time.Second
	defaultHTTPShutdownTimeout     = 10 * time.Second
)

// Config contains the backend's technical runtime configuration.
type Config struct {
	Environment             string
	HTTPPort                int
	DatabaseURL             string
	LogLevel                string
	WalkWorkerConcurrency   int
	OutboxWorkerConcurrency int
	HTTPReadTimeout         time.Duration
	HTTPWriteTimeout        time.Duration
	HTTPShutdownTimeout     time.Duration
}

// Load reads configuration from environment variables and validates it.
func Load() (Config, error) {
	config := Config{
		Environment:             valueOrDefault("APP_ENV", defaultEnvironment),
		HTTPPort:                defaultHTTPPort,
		DatabaseURL:             strings.TrimSpace(os.Getenv("DATABASE_URL")),
		LogLevel:                valueOrDefault("LOG_LEVEL", defaultLogLevel),
		WalkWorkerConcurrency:   defaultWalkWorkerConcurrency,
		OutboxWorkerConcurrency: defaultOutboxWorkerConcurrency,
		HTTPReadTimeout:         defaultHTTPReadTimeout,
		HTTPWriteTimeout:        defaultHTTPWriteTimeout,
		HTTPShutdownTimeout:     defaultHTTPShutdownTimeout,
	}

	var err error

	config.HTTPPort, err = intValue("HTTP_PORT", defaultHTTPPort)
	if err != nil {
		return Config{}, err
	}

	config.WalkWorkerConcurrency, err = intValue(
		"WALK_WORKER_CONCURRENCY",
		defaultWalkWorkerConcurrency,
	)
	if err != nil {
		return Config{}, err
	}

	config.OutboxWorkerConcurrency, err = intValue(
		"OUTBOX_WORKER_CONCURRENCY",
		defaultOutboxWorkerConcurrency,
	)
	if err != nil {
		return Config{}, err
	}

	config.HTTPReadTimeout, err = durationValue("HTTP_READ_TIMEOUT", defaultHTTPReadTimeout)
	if err != nil {
		return Config{}, err
	}

	config.HTTPWriteTimeout, err = durationValue("HTTP_WRITE_TIMEOUT", defaultHTTPWriteTimeout)
	if err != nil {
		return Config{}, err
	}

	config.HTTPShutdownTimeout, err = durationValue(
		"HTTP_SHUTDOWN_TIMEOUT",
		defaultHTTPShutdownTimeout,
	)
	if err != nil {
		return Config{}, err
	}

	if err := config.validate(); err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c Config) validate() error {
	if !isSupportedEnvironment(c.Environment) {
		return fmt.Errorf("APP_ENV must be one of: local, staging, production, test")
	}

	if c.DatabaseURL == "" {
		return fmt.Errorf("DATABASE_URL must not be empty")
	}

	if c.HTTPPort < 1 || c.HTTPPort > 65535 {
		return fmt.Errorf("HTTP_PORT must be between 1 and 65535")
	}

	if c.WalkWorkerConcurrency < 0 {
		return fmt.Errorf("WALK_WORKER_CONCURRENCY must not be negative")
	}

	if c.OutboxWorkerConcurrency < 0 {
		return fmt.Errorf("OUTBOX_WORKER_CONCURRENCY must not be negative")
	}

	return nil
}

func durationValue(key string, defaultValue time.Duration) (time.Duration, error) {
	rawValue := os.Getenv(key)
	if rawValue == "" {
		return defaultValue, nil
	}

	value, err := time.ParseDuration(rawValue)
	if err != nil {
		return 0, fmt.Errorf("parse %s: %w", key, err)
	}

	return value, nil
}

func intValue(key string, defaultValue int) (int, error) {
	rawValue := os.Getenv(key)
	if rawValue == "" {
		return defaultValue, nil
	}

	value, err := strconv.Atoi(rawValue)
	if err != nil {
		return 0, fmt.Errorf("parse %s: %w", key, err)
	}

	return value, nil
}

func isSupportedEnvironment(environment string) bool {
	switch environment {
	case "local", "staging", "production", "test":
		return true
	default:
		return false
	}
}

func valueOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
