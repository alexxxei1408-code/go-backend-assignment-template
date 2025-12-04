package main

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func newMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New error: %v", err)
	}
	return db, mock
}

func TestInsertUser_Basic(t *testing.T) {
	db, mock := newMockDB(t)
	defer db.Close()

	u := User{ID: 1, Name: "Alice", Email: "a@example.com", Active: true}

	mock.ExpectExec(`INSERT INTO users`).
		WithArgs(u.ID, u.Name, u.Email, u.Active).
		WillReturnResult(sqlmock.NewResult(0, 1))

	if err := InsertUser(context.Background(), db, u); err != nil {
		t.Fatalf("InsertUser error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func TestGetActiveUsers_Basic(t *testing.T) {
	db, mock := newMockDB(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "email", "active"}).
		AddRow(1, "Alice", "a@example.com", true).
		AddRow(2, "Bob", "b@example.com", true)

	mock.ExpectQuery(`SELECT id, name, email, active FROM users WHERE active = TRUE ORDER BY id`).
		WillReturnRows(rows)

	users, err := GetActiveUsers(context.Background(), db)
	if err != nil {
		t.Fatalf("GetActiveUsers error: %v", err)
	}

	if len(users) != 2 {
		t.Fatalf("GetActiveUsers len: got %d, want %d", len(users), 2)
	}

	if users[0].ID != 1 || users[1].ID != 2 {
		t.Fatalf("unexpected order: %+v", users)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func TestDeactivateUser_Basic(t *testing.T) {
	db, mock := newMockDB(t)
	defer db.Close()

	mock.ExpectExec(`UPDATE users SET active = FALSE WHERE id = \?`).
		WithArgs(int64(1)).
		WillReturnResult(sqlmock.NewResult(0, 1))

	if err := DeactivateUser(context.Background(), db, 1); err != nil {
		t.Fatalf("DeactivateUser error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}


