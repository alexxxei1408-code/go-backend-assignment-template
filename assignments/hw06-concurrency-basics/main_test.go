package main

import "testing"

func TestSafeCounter_Sequential(t *testing.T) {
	c := NewSafeCounter()

	c.Inc("a")
	c.Inc("a")
	c.Inc("b")

	if got := c.Value("a"); got != 2 {
		t.Fatalf("Value(\"a\"): got %d, want %d", got, 2)
	}
	if got := c.Value("b"); got != 1 {
		t.Fatalf("Value(\"b\"): got %d, want %d", got, 1)
	}
	if got := c.Value("c"); got != 0 {
		t.Fatalf("Value(\"c\") for unknown key: got %d, want %d", got, 0)
	}
	if got := c.Total(); got != 3 {
		t.Fatalf("Total: got %d, want %d", got, 3)
	}
}

func TestSumConcurrently_Basic(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	if got := SumConcurrently(nums, 2); got != 15 {
		t.Fatalf("SumConcurrently: got %d, want %d", got, 15)
	}
}

func TestSumConcurrently_WorkersEdgeCases(t *testing.T) {
	nums := []int{10, -5, 3}

	if got := SumConcurrently(nums, 0); got != 8 {
		t.Fatalf("workers=0: got %d, want %d", got, 8)
	}
	if got := SumConcurrently(nums, -1); got != 8 {
		t.Fatalf("workers<0: got %d, want %d", got, 8)
	}
}


