package main

import "net/http"

// HandleHello обрабатывает запросы к /api/hello.
// Реализуйте логику:
// - Обработка GET-запросов с параметром name.
// - Формирование и возврат JSON-ответа.
// - Обработка некорректных методов (например, POST).
func HandleHello(w http.ResponseWriter, r *http.Request) {
	// TODO: Ваша реализация здесь
}

// main оставлен пустым, чтобы сборка успешно проходила.
// В более сложных заданиях сюда можно вынести запуск HTTP-сервера.
func main() {
}


