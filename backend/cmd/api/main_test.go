package main

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestServe_ShutsDownOnCanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	server := &http.Server{
		Addr:    "127.0.0.1:0",
		Handler: http.NewServeMux(),
	}

	if err := serve(ctx, server, time.Second); err != nil {
		t.Fatalf("serve() error = %v, want nil", err)
	}
}

func TestServe_ReturnsUnexpectedServerError(t *testing.T) {
	server := &http.Server{
		Addr:    ":invalid",
		Handler: http.NewServeMux(),
	}

	if err := serve(context.Background(), server, time.Second); err == nil {
		t.Fatal("serve() error = nil, want unexpected server error")
	}
}
