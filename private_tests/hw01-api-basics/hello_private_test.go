package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// Приватный тест: проверяет корректный Content-Type для успешного ответа.
func TestPrivateHandleHelloContentType(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/hello?name=Go", nil)
	w := httptest.NewRecorder()

	HandleHello(w, req)

	ct := w.Header().Get("Content-Type")
	if !strings.HasPrefix(ct, "application/json") {
		t.Fatalf("unexpected Content-Type: got %q, want prefix %q", ct, "application/json")
	}
}

// Приватный тест: проверяет работу с Unicode-именем.
func TestPrivateHandleHelloUnicodeName(t *testing.T) {
	name := "Голанг"

	q := url.Values{}
	q.Set("name", name)

	req := httptest.NewRequest(http.MethodGet, "/api/hello?"+q.Encode(), nil)
	w := httptest.NewRecorder()

	HandleHello(w, req)

	wantBody := `{"message":"Hello, ` + name + `!"}`
	if w.Code != http.StatusOK {
		t.Fatalf("status code: got %d, want %d", w.Code, http.StatusOK)
	}

	if got := strings.TrimSpace(w.Body.String()); got != wantBody {
		t.Fatalf("response body: got %q, want %q", got, wantBody)
	}
}

// Приватный тест: проверяет, что все методы, кроме GET, запрещены.
func TestPrivateHandleHelloMethodsOtherThanGET(t *testing.T) {
	methods := []string{
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodPatch,
	}

	for _, method := range methods {
		t.Run(method, func(t *testing.T) {
			req := httptest.NewRequest(method, "/api/hello", nil)
			w := httptest.NewRecorder()

			HandleHello(w, req)

			if w.Code != http.StatusMethodNotAllowed {
				t.Fatalf("method %s: status code: got %d, want %d", method, w.Code, http.StatusMethodNotAllowed)
			}
		})
	}
}


