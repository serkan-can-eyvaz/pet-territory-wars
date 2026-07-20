package health

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLive(t *testing.T) {
	mux := http.NewServeMux()
	RegisterRoutes(mux, fakeDatabasePinger{})

	response := performRequest(t, mux, http.MethodGet, livePath)

	assertResponse(t, response, http.StatusOK, "ok")
}

func TestLiveRejectsNonGETMethods(t *testing.T) {
	mux := http.NewServeMux()
	RegisterRoutes(mux, fakeDatabasePinger{})

	response := performRequest(t, mux, http.MethodPost, livePath)

	assertResponse(t, response, http.StatusMethodNotAllowed, "method_not_allowed")
	if allow := response.Header().Get("Allow"); allow != http.MethodGet {
		t.Errorf("Allow = %q, want %q", allow, http.MethodGet)
	}
}

func TestReady(t *testing.T) {
	mux := http.NewServeMux()
	RegisterRoutes(mux, fakeDatabasePinger{})

	response := performRequest(t, mux, http.MethodGet, readyPath)

	assertResponse(t, response, http.StatusOK, "ready")
}

func TestReadyReturnsServiceUnavailableWithoutDatabaseDetails(t *testing.T) {
	databaseError := errors.New("database connection refused")
	mux := http.NewServeMux()
	RegisterRoutes(mux, fakeDatabasePinger{err: databaseError})

	response := performRequest(t, mux, http.MethodGet, readyPath)

	assertResponse(t, response, http.StatusServiceUnavailable, "not_ready")
	if strings.Contains(response.Body.String(), databaseError.Error()) {
		t.Error("readiness response exposed the database error")
	}
}

func TestReadyRejectsNonGETMethods(t *testing.T) {
	mux := http.NewServeMux()
	RegisterRoutes(mux, fakeDatabasePinger{})

	response := performRequest(t, mux, http.MethodPost, readyPath)

	assertResponse(t, response, http.StatusMethodNotAllowed, "method_not_allowed")
}

func performRequest(t *testing.T, handler http.Handler, method, path string) *httptest.ResponseRecorder {
	t.Helper()

	request := httptest.NewRequest(method, path, nil)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	return response
}

func assertResponse(t *testing.T, recorder *httptest.ResponseRecorder, statusCode int, status string) {
	t.Helper()

	if recorder.Code != statusCode {
		t.Errorf("status code = %d, want %d", recorder.Code, statusCode)
	}
	if contentType := recorder.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("Content-Type = %q, want %q", contentType, "application/json")
	}

	var body response
	if err := json.Unmarshal(recorder.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if body.Status != status {
		t.Errorf("status = %q, want %q", body.Status, status)
	}
}

type fakeDatabasePinger struct {
	err error
}

func (f fakeDatabasePinger) Ping(context.Context) error {
	return f.err
}
