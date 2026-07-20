// Package logging provides structured logger construction.
package logging

import (
	"fmt"
	"io"
	"log/slog"
	"strings"
)

// New creates a JSON logger that writes to writer with the configured level.
func New(writer io.Writer, service, level string) (*slog.Logger, error) {
	slogLevel, err := parseLevel(level)
	if err != nil {
		return nil, err
	}

	handler := slog.NewJSONHandler(writer, &slog.HandlerOptions{Level: slogLevel})
	logger := slog.New(handler)

	return logger.With(slog.String("service", service)), nil
}

func parseLevel(level string) (slog.Level, error) {
	switch strings.ToLower(strings.TrimSpace(level)) {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	default:
		return 0, fmt.Errorf("unsupported log level %q", level)
	}
}
