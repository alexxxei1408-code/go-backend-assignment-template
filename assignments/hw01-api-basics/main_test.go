package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHelloGET(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		wantCode int
		wantBody string
	}{
		{
			name:     "with name",
			query:    "name=Go",
			wantCode: http.StatusOK,
			wantBody: `{"message":"Hello, Go!"}`,
		},
		{
			name:     "without name",
			query:    "",
			wantCode: http.StatusOK,
			wantBody: `{"message":"Hello, World!"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api/hello?"+tt.query, nil)
			w := httptest.NewRecorder()

			HandleHello(w, req)

			if w.Code != tt.wantCode {
				t.Errorf("код ответа: получили %d, хотели %d", w.Code, tt.wantCode)
			}

			if got := w.Body.String(); got != tt.wantBody {
				t.Errorf("тело ответа: получили %s, хотели %s", got, tt.wantBody)
			}
		})
	}
}

func TestHandleHelloPOST(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/hello", nil)
	w := httptest.NewRecorder()

	HandleHello(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("ожидали 405, получили %d", w.Code)
	}
}


