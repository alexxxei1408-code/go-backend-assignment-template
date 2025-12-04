package main

import "io"

// User описывает пользователя в JSON-модели.
type User struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Active bool   `json:"active"`
}

// DecodeUsers читает из r JSON-массив пользователей и возвращает его в виде слайса.
//
// Требования:
//   - Ожидается JSON-массив объектов вида {"id": ..., "name": ..., "email": ..., "active": ...}.
//   - Неизвестные поля должны игнорироваться.
//   - При некорректном JSON должна возвращаться ошибка.
func DecodeUsers(r io.Reader) ([]User, error) {
	// TODO: реализуйте функцию с использованием encoding/json
	return nil, nil
}

// EncodeActiveUsers записывает в w JSON-массив только активных пользователей.
//
// Требования:
//   - В результирующий JSON попадают только записи с Active == true.
//   - Пользователи должны быть отсортированы по возрастанию ID.
//   - Формат вывода — корректный JSON-массив (пробелы и переносы строк значения не имеют).
func EncodeActiveUsers(w io.Writer, users []User) error {
	// TODO: реализуйте функцию с использованием encoding/json
	return nil
}

func main() {}


