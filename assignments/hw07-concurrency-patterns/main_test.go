package main

import (
	"sort"
	"testing"
	"time"
)

func TestRunWorkerPool_Basic(t *testing.T) {
	jobs := []Job{
		{ID: 2, Input: 3},
		{ID: 1, Input: 2},
		{ID: 3, Input: 4},
	}

	fn := func(x int) int { return x * x }

	results := RunWorkerPool(jobs, 2, fn)
	if len(results) != len(jobs) {
		t.Fatalf("RunWorkerPool len: got %d, want %d", len(results), len(jobs))
	}

	for i, r := range results {
		if r.JobID != i+1 {
			t.Fatalf("results[%d].JobID = %d, want %d", i, r.JobID, i+1)
		}
		want := (i + 1) * (i + 1)
		if r.Output != want {
			t.Fatalf("result for job %d: got %d, want %d", r.JobID, r.Output, want)
		}
	}
}

func TestRunWorkerPool_WorkersEdgeCases(t *testing.T) {
	jobs := []Job{
		{ID: 1, Input: 1},
		{ID: 2, Input: 2},
	}
	fn := func(x int) int { return x + 1 }

	r1 := RunWorkerPool(jobs, 0, fn)
	r2 := RunWorkerPool(jobs, -5, fn)

	if len(r1) != len(jobs) || len(r2) != len(jobs) {
		t.Fatalf("unexpected len for workers<=0")
	}
}

func TestFanIn_Basic(t *testing.T) {
	ch1 := make(chan int, 3)
	ch2 := make(chan int, 3)

	for _, v := range []int{1, 2, 3} {
		ch1 <- v
	}
	for _, v := range []int{4, 5} {
		ch2 <- v
	}

	close(ch1)
	close(ch2)

	out := FanIn(ch1, ch2)

	var got []int
	timeout := time.After(2 * time.Second)

loop:
	for {
		select {
		case v, ok := <-out:
			if !ok {
				break loop
			}
			got = append(got, v)
		case <-timeout:
			t.Fatal("FanIn timeout waiting for data")
		}
	}

	sort.Ints(got)
	want := []int{1, 2, 3, 4, 5}
	if len(got) != len(want) {
		t.Fatalf("FanIn len: got %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("FanIn item %d: got %d, want %d", i, got[i], want[i])
		}
	}
}


