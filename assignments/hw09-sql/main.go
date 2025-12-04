package main

import (
	"context"
	"database/sql"
)

// User описывает запись в таблице users.
type User struct {
	ID     int64
	Name   string
	Email  string
	Active bool
}

// InsertUser добавляет нового пользователя в таблицу users.
//
// Рекомендуемый SQL:
//   INSERT INTO users (id, name, email, active) VALUES (?, ?, ?, ?)
func InsertUser(ctx context.Context, db *sql.DB, u User) error {
	// TODO: выполните INSERT через ExecContext
	return nil
}

// GetActiveUsers возвращает всех активных пользователей, отсортированных по возрастанию ID.
//
// Рекомендуемый SQL:
//   SELECT id, name, email, active FROM users WHERE active = TRUE ORDER BY id
func GetActiveUsers(ctx context.Context, db *sql.DB) ([]User, error) {
	// TODO: выполните SELECT через QueryContext и соберите результаты в слайс
	return nil, nil
}

// DeactivateUser помечает пользователя как неактивного по его ID.
//
// Рекомендуемый SQL:
//   UPDATE users SET active = FALSE WHERE id = ?
func DeactivateUser(ctx context.Context, db *sql.DB, id int64) error {
	// TODO: выполните UPDATE через ExecContext
	return nil
}

func main() {}


