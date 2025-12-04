package main

import (
	"context"
	"database/sql"
	"net/http"
)

// Stats описывает агрегированную статистику по пользователям.
type Stats struct {
	ActiveUsers   int64 `json:"active_users"`
	InactiveUsers int64 `json:"inactive_users"`
}

// LoadUserStats загружает статистику по активным/неактивным пользователям из БД.
//
// Требования:
//   - Необходимо выполнить два запроса:
//     SELECT COUNT(*) FROM users WHERE active = TRUE
//     SELECT COUNT(*) FROM users WHERE active = FALSE
//   - Рекомендуется выполнять запросы конкурентно с синхронизацией через WaitGroup
//     и передачей ошибок через каналы.
func LoadUserStats(ctx context.Context, db *sql.DB) (Stats, error) {
	// TODO: реализуйте загрузку статистики из БД
	return Stats{}, nil
}

// StatsHandler создаёт http.Handler, который отвечает на GET /stats JSON-объектом Stats.
//
// Требования:
//   - Для обработки запроса нужно использовать LoadUserStats с контекстом запроса.
//   - При успешной загрузке статистики возвращать 200 OK и JSON-тело.
//   - При ошибке БД возвращать 500 Internal Server Error.
//   - Content-Type ответа должен начинаться с "application/json".
func StatsHandler(db *sql.DB) http.Handler {
	// TODO: реализуйте хэндлер с использованием LoadUserStats
	return nil
}

func main() {}


