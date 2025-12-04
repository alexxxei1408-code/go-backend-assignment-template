package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func newStatsMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New error: %v", err)
	}
	return db, mock
}

func TestLoadUserStats_Basic(t *testing.T) {
	db, mock := newStatsMockDB(t)
	defer db.Close()

	activeRows := sqlmock.NewRows([]string{"count"}).AddRow(5)
	inactiveRows := sqlmock.NewRows([]string{"count"}).AddRow(2)

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM users WHERE active = TRUE`).
		WillReturnRows(activeRows)
	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM users WHERE active = FALSE`).
		WillReturnRows(inactiveRows)

	stats, err := LoadUserStats(context.Background(), db)
	if err != nil {
		t.Fatalf("LoadUserStats error: %v", err)
	}

	if stats.ActiveUsers != 5 || stats.InactiveUsers != 2 {
		t.Fatalf("LoadUserStats result: %+v", stats)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func TestStatsHandler_OK(t *testing.T) {
	db, mock := newStatsMockDB(t)
	defer db.Close()

	activeRows := sqlmock.NewRows([]string{"count"}).AddRow(3)
	inactiveRows := sqlmock.NewRows([]string{"count"}).AddRow(1)

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM users WHERE active = TRUE`).
		WillReturnRows(activeRows)
	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM users WHERE active = FALSE`).
		WillReturnRows(inactiveRows)

	h := StatsHandler(db)
	if h == nil {
		t.Fatalf("StatsHandler returned nil")
	}

	req := httptest.NewRequest(http.MethodGet, "/stats", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status: got %d, want %d", w.Code, http.StatusOK)
	}

	ct := w.Header().Get("Content-Type")
	if ct == "" || ct[:16] != "application/json" {
		t.Fatalf("unexpected Content-Type: %q", ct)
	}

	var stats Stats
	if err := json.Unmarshal(w.Body.Bytes(), &stats); err != nil {
		t.Fatalf("response is not valid JSON: %v", err)
	}
	if stats.ActiveUsers != 3 || stats.InactiveUsers != 1 {
		t.Fatalf("unexpected stats: %+v", stats)
	}
}


