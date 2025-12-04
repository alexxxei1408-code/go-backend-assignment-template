package main

import (
	"reflect"
	"testing"
)

func TestTopUsersBySpending_Basic(t *testing.T) {
	purchases := []Purchase{
		{UserID: 1, Category: "books", Amount: 100},
		{UserID: 2, Category: "games", Amount: 500},
		{UserID: 1, Category: "games", Amount: 200},
		{UserID: 2, Category: "books", Amount: 100},
		{UserID: 3, Category: "food", Amount: 50},
	}

	got := TopUsersBySpending(purchases, 2)

	want := []UserStats{
		{UserID: 2, TotalAmount: 600, Categories: 2},
		{UserID: 1, TotalAmount: 300, Categories: 2},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("TopUsersBySpending: got %+v, want %+v", got, want)
	}
}

func TestTopUsersBySpending_TopNGreaterThanUsers(t *testing.T) {
	purchases := []Purchase{
		{UserID: 1, Category: "books", Amount: 100},
		{UserID: 2, Category: "games", Amount: 200},
	}

	got := TopUsersBySpending(purchases, 10)

	want := []UserStats{
		{UserID: 2, TotalAmount: 200, Categories: 1},
		{UserID: 1, TotalAmount: 100, Categories: 1},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("TopUsersBySpending: got %+v, want %+v", got, want)
	}
}

func TestTopUsersBySpending_TopNZeroOrNegative(t *testing.T) {
	purchases := []Purchase{
		{UserID: 1, Category: "books", Amount: 100},
	}

	if got := TopUsersBySpending(purchases, 0); len(got) != 0 {
		t.Fatalf("expected empty slice for topN=0, got: %+v", got)
	}
	if got := TopUsersBySpending(purchases, -1); len(got) != 0 {
		t.Fatalf("expected empty slice for topN<0, got: %+v", got)
	}
}

func TestCategoryTotals_Basic(t *testing.T) {
	purchases := []Purchase{
		{UserID: 1, Category: "books", Amount: 100},
		{UserID: 2, Category: "games", Amount: 200},
		{UserID: 3, Category: "books", Amount: 50},
	}

	got := CategoryTotals(purchases)

	want := map[string]int64{
		"books": 150,
		"games": 200,
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("CategoryTotals: got %+v, want %+v", got, want)
	}
}

func TestCategoryTotals_EmptyInput(t *testing.T) {
	got := CategoryTotals(nil)
	if got == nil {
		t.Fatalf("CategoryTotals: expected non-nil map for nil input")
	}
	if len(got) != 0 {
		t.Fatalf("CategoryTotals: expected empty map for empty input, got: %+v", got)
	}
}
