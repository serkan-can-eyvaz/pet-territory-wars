package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/infrastructure/config"
	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/infrastructure/database"
	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/infrastructure/logging"
	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/transport/health"
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

	logger, err := logging.New(os.Stdout, "api", configuration.LogLevel)
	if err != nil {
		return fmt.Errorf("create logger: %w", err)
	}
	logger.Info("api starting")

	pool, err := database.Open(ctx, configuration.DatabaseURL)
	if err != nil {
		return fmt.Errorf("open database pool: %w", err)
	}
	defer pool.Close()

	mux := http.NewServeMux()
	health.RegisterRoutes(mux, pool)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", configuration.HTTPPort),
		Handler:      mux,
		ReadTimeout:  configuration.HTTPReadTimeout,
		WriteTimeout: configuration.HTTPWriteTimeout,
	}

	serverErrors := make(chan error, 1)
	go func() {
		serverErrors <- server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("serve HTTP: %w", err)
		}
		return nil
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			configuration.HTTPShutdownTimeout,
		)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("shutdown HTTP server: %w", err)
		}

		if err := <-serverErrors; err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("serve HTTP: %w", err)
		}

		return nil
	}
}
