package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestDecodeUsers_Basic(t *testing.T) {
	data := `
[
  {"id": 1, "name": "Alice", "email": "a@example.com", "active": true},
  {"id": 2, "name": "Bob", "email": "b@example.com", "active": false}
]`

	users, err := DecodeUsers(strings.NewReader(data))
	if err != nil {
		t.Fatalf("DecodeUsers error: %v", err)
	}

	if len(users) != 2 {
		t.Fatalf("DecodeUsers len: got %d, want %d", len(users), 2)
	}

	if users[0].ID != 1 || users[0].Name != "Alice" || !users[0].Active {
		t.Fatalf("DecodeUsers first element mismatch: %+v", users[0])
	}
}

func TestDecodeUsers_InvalidJSON(t *testing.T) {
	data := `{"id": 1, "name": "broken"}`
	if _, err := DecodeUsers(strings.NewReader(data)); err == nil {
		t.Fatalf("DecodeUsers: expected error for non-array JSON")
	}
}

func TestEncodeActiveUsers_FiltersAndSorts(t *testing.T) {
	users := []User{
		{ID: 2, Name: "Bob", Email: "b@example.com", Active: true},
		{ID: 1, Name: "Alice", Email: "a@example.com", Active: false},
		{ID: 3, Name: "Carol", Email: "c@example.com", Active: true},
	}

	var buf bytes.Buffer
	if err := EncodeActiveUsers(&buf, users); err != nil {
		t.Fatalf("EncodeActiveUsers error: %v", err)
	}

	var decoded []User
	if err := json.Unmarshal(buf.Bytes(), &decoded); err != nil {
		t.Fatalf("result is not valid JSON: %v", err)
	}

	if len(decoded) != 2 {
		t.Fatalf("expected 2 active users, got %d", len(decoded))
	}

	if decoded[0].ID != 2 || decoded[1].ID != 3 {
		t.Fatalf("expected users with IDs 2 and 3, got %v and %v", decoded[0].ID, decoded[1].ID)
	}
}


