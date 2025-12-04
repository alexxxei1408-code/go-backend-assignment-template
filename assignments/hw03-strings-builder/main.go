package main

// LogRecord описывает одну запись лога приложения.
type LogRecord struct {
	Level   string
	Message string
}

// FormatLog форматирует слайс записей лога в единую строку.
//
// Требования:
//   - Каждая запись должна быть на отдельной строке.
//   - Формат строки: "<LEVEL>: <Message>\n", где LEVEL — значение Level в верхнем регистре.
//   - При пустом слайсе должна возвращаться пустая строка.
//   - Для конкатенации строк рекомендуется использовать strings.Builder.
func FormatLog(records []LogRecord) string {
	// TODO: реализуйте функцию, эффективно собирая строку
	return ""
}

func main() {}


