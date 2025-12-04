package main

import "net/http"

// Logger описывает минимальный интерфейс логгера,
// используемого в middleware.
type Logger interface {
	Log(line string)
}

// LoggingMiddleware оборачивает http.Handler и логирует каждый запрос.
//
// Требования:
//   - Логируемая строка должна иметь вид "<METHOD> <PATH>".
//   - Логирование должно выполняться до вызова next.ServeHTTP.
func LoggingMiddleware(logger Logger, next http.Handler) http.Handler {
	// TODO: реализуйте middleware
	return nil
}

// APIKeyAuthMiddleware проверяет наличие и корректность ключа доступа в заголовке X-API-Key.
//
// Требования:
//   - Если заголовок X-API-Key совпадает с expectedKey, запрос передаётся дальше.
//   - В противном случае должен возвращаться статус 401 Unauthorized,
//     тело может быть пустым, next вызываться не должен.
func APIKeyAuthMiddleware(expectedKey string, next http.Handler) http.Handler {
	// TODO: реализуйте middleware
	return nil
}

// PingHandler отвечает на GET /ping статусом 200 и телом "pong".
// Для остальных методов должен возвращаться 405 Method Not Allowed.
func PingHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: реализуйте обработчик
}

func main() {}


