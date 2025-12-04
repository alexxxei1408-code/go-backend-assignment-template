package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type stubLogger struct {
	lines []string
}

func (s *stubLogger) Log(line string) {
	s.lines = append(s.lines, line)
}

func TestPingHandler_GET(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()

	PingHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("PingHandler status: got %d, want %d", w.Code, http.StatusOK)
	}
	if body := w.Body.String(); body != "pong" {
		t.Fatalf("PingHandler body: got %q, want %q", body, "pong")
	}
}

func TestPingHandler_OtherMethods(t *testing.T) {
	methods := []string{
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
	}

	for _, m := range methods {
		t.Run(m, func(t *testing.T) {
			req := httptest.NewRequest(m, "/ping", nil)
			w := httptest.NewRecorder()

			PingHandler(w, req)

			if w.Code != http.StatusMethodNotAllowed {
				t.Fatalf("PingHandler %s: status got %d, want %d", m, w.Code, http.StatusMethodNotAllowed)
			}
		})
	}
}

func TestLoggingMiddleware_LogsAndCallsNext(t *testing.T) {
	logger := &stubLogger{}

	called := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusTeapot)
	})

	h := LoggingMiddleware(logger, next)
	if h == nil {
		t.Fatalf("LoggingMiddleware returned nil handler")
	}

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)

	if !called {
		t.Fatalf("next handler was not called")
	}
	if len(logger.lines) != 1 || logger.lines[0] != "GET /test" {
		t.Fatalf("unexpected logs: %+v", logger.lines)
	}
	if w.Code != http.StatusTeapot {
		t.Fatalf("status code from next handler: got %d, want %d", w.Code, http.StatusTeapot)
	}
}

func TestAPIKeyAuthMiddleware_AllowsValidKey(t *testing.T) {
	const key = "secret"

	called := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	})

	h := APIKeyAuthMiddleware(key, next)
	if h == nil {
		t.Fatalf("APIKeyAuthMiddleware returned nil handler")
	}

	req := httptest.NewRequest(http.MethodGet, "/secure", nil)
	req.Header.Set("X-API-Key", key)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)

	if !called {
		t.Fatalf("next handler was not called for valid key")
	}
	if w.Code != http.StatusOK {
		t.Fatalf("status for valid key: got %d, want %d", w.Code, http.StatusOK)
	}
}

func TestAPIKeyAuthMiddleware_RejectsInvalidKey(t *testing.T) {
	const key = "secret"

	called := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
	})

	h := APIKeyAuthMiddleware(key, next)

	req := httptest.NewRequest(http.MethodGet, "/secure", nil)
	req.Header.Set("X-API-Key", "wrong")
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)

	if called {
		t.Fatalf("next handler must not be called for invalid key")
	}
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("status for invalid key: got %d, want %d", w.Code, http.StatusUnauthorized)
	}
}


