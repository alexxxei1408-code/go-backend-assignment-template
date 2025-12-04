package main

import (
	"bytes"
	"io"
	"testing"
)

func TestLimitReader_ReadLessThanLimit(t *testing.T) {
	src := bytes.NewBufferString("hello")
	lr := NewLimitReader(src, 3)

	data, err := io.ReadAll(lr)
	if err != nil {
		t.Fatalf("ReadAll error: %v", err)
	}

	if string(data) != "hel" {
		t.Fatalf("LimitReader: got %q, want %q", string(data), "hel")
	}
}

func TestLimitReader_ReadInChunks(t *testing.T) {
	src := bytes.NewBufferString("abcdef")
	lr := NewLimitReader(src, 4)

	buf := make([]byte, 2)

	n1, err1 := lr.Read(buf)
	if err1 != nil && err1 != nil {
		t.Fatalf("unexpected error on first read: %v", err1)
	}
	if n1 != 2 || string(buf[:n1]) != "ab" {
		t.Fatalf("first read: n=%d, data=%q", n1, string(buf[:n1]))
	}

	n2, err2 := lr.Read(buf)
	if err2 != nil && err2 != nil {
		t.Fatalf("unexpected error on second read: %v", err2)
	}
	if n2 != 2 || string(buf[:n2]) != "cd" {
		t.Fatalf("second read: n=%d, data=%q", n2, string(buf[:n2]))
	}

	n3, err3 := lr.Read(buf)
	if n3 != 0 || err3 != io.EOF {
		t.Fatalf("expected EOF after limit, got n=%d, err=%v", n3, err3)
	}
}

func TestLimitReader_ZeroLimit(t *testing.T) {
	src := bytes.NewBufferString("data")
	lr := NewLimitReader(src, 0)

	buf := make([]byte, 10)
	n, err := lr.Read(buf)
	if n != 0 || err != io.EOF {
		t.Fatalf("zero limit: expected EOF immediately, got n=%d, err=%v", n, err)
	}
}

func TestCountingWriter_CountsBytes(t *testing.T) {
	var buf bytes.Buffer
	cw := NewCountingWriter(&buf)

	n1, err1 := cw.Write([]byte("hi"))
	if err1 != nil || n1 != 2 {
		t.Fatalf("first Write: n=%d err=%v", n1, err1)
	}

	n2, err2 := cw.Write([]byte("!"))
	if err2 != nil || n2 != 1 {
		t.Fatalf("second Write: n=%d err=%v", n2, err2)
	}

	if cw.N != 3 {
		t.Fatalf("CountingWriter.N: got %d, want %d", cw.N, 3)
	}
	if buf.String() != "hi!" {
		t.Fatalf("underlying buffer: got %q, want %q", buf.String(), "hi!")
	}
}


