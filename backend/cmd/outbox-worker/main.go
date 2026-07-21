package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/infrastructure/config"
	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/infrastructure/database"
	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/infrastructure/logging"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	configuration, err := config.Load()
	if err != nil {
		return fmt.Errorf("load configuration: %w", err)
	}

	logger, err := logging.New(os.Stdout, "outbox-worker", configuration.LogLevel)
	if err != nil {
		return fmt.Errorf("create logger: %w", err)
	}

	pool, err := database.Open(ctx, configuration.DatabaseURL)
	if err != nil {
		return fmt.Errorf("open database pool: %w", err)
	}
	defer pool.Close()

	logger.Info("outbox worker started")
	<-ctx.Done()
	logger.Info("outbox worker shutting down")

	return nil
}
