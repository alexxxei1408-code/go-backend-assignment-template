package main

import "io"

// LimitReader ограничивает количество байт, читаемых из базового io.Reader.
// После того как будет прочитано N байт, все последующие вызовы Read должны
// возвращать 0, io.EOF.
type LimitReader struct {
	R io.Reader
	N int64
}

// NewLimitReader создаёт LimitReader с указанным лимитом.
func NewLimitReader(r io.Reader, n int64) *LimitReader {
	// TODO: инициализируйте структуру
	return nil
}

// Read читает не более чем N оставшихся байт из базового Reader.
func (lr *LimitReader) Read(p []byte) (int, error) {
	// TODO: реализуйте логику ограничения чтения
	return 0, nil
}

// CountingWriter оборачивает io.Writer и считает количество успешно записанных байт.
type CountingWriter struct {
	W io.Writer
	N int64 // количество записанных байт
}

// NewCountingWriter создаёт CountingWriter, оборачивающий переданный Writer.
func NewCountingWriter(w io.Writer) *CountingWriter {
	// TODO: инициализируйте структуру
	return nil
}

// Write записывает данные в базовый Writer и увеличивает счётчик N
// на количество фактически записанных байт.
func (cw *CountingWriter) Write(p []byte) (int, error) {
	// TODO: реализуйте подсчёт записанных байт
	return 0, nil
}

func main() {}


