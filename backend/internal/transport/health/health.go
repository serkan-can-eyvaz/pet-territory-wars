// Package health provides HTTP health endpoints.
package health

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	livePath  = "/health/live"
	readyPath = "/health/ready"
)

// DatabasePinger is the database availability check used by readiness.
type DatabasePinger interface {
	Ping(ctx context.Context) error
}

type handler struct {
	databasePinger DatabasePinger
}

type response struct {
	Status string `json:"status"`
}

// RegisterRoutes registers the health endpoints on mux.
func RegisterRoutes(mux *http.ServeMux, databasePinger DatabasePinger) {
	handler := handler{databasePinger: databasePinger}

	mux.HandleFunc(livePath, handler.live)
	mux.HandleFunc(readyPath, handler.ready)
}

func (h handler) live(w http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}

	writeResponse(w, http.StatusOK, response{Status: "ok"})
}

func (h handler) ready(w http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}

	if h.databasePinger == nil || h.databasePinger.Ping(request.Context()) != nil {
		writeResponse(w, http.StatusServiceUnavailable, response{Status: "not_ready"})
		return
	}

	writeResponse(w, http.StatusOK, response{Status: "ready"})
}

func methodNotAllowed(w http.ResponseWriter) {
	w.Header().Set("Allow", http.MethodGet)
	writeResponse(w, http.StatusMethodNotAllowed, response{Status: "method_not_allowed"})
}

func writeResponse(w http.ResponseWriter, statusCode int, response response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(response)
}
