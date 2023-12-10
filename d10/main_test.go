package main

import (
	"reflect"
	"testing"
)

func Test_start_neighbors(t *testing.T) {
	example1 := parse("example1.txt")
	start := findStart(example1)
	got := startNeighbors(example1, start)

	expected := []pos{{1, 2}, {2, 1}}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected: %v, got: %v", expected, got)
	}
}

func Test_filter_row(t *testing.T) {
	example := parse("example3.txt")
	start := findStart(example)
	current := pos{3, 3}
	got := filterRowPipes(current, example, pipeLoop(example, start))

	expected := "||"

	if got != expected {
		t.Fatalf("expected: %v, got: %v", expected, got)
	}
}

func Test_replace_turns(t *testing.T) {
	testcases := []struct {
		row      string
		expected string
	}{
		{"F-J", "|"},
		{"LJ||LJL-7", "|||||||"},
	}

	for _, tc := range testcases {
		got := replaceTurns(tc.row)

		if got != tc.expected {
			t.Fatalf("expected: %v, got: %v", tc.expected, got)
		}
	}
}

func Test_is_inside(t *testing.T) {
	testcases := []struct {
		input    inputT
		current  pos
		expected bool
	}{
		{parse("example3.txt"), pos{3, 3}, false},
		{parse("example3.txt"), pos{3, 4}, false},
		{parse("example3.txt"), pos{6, 2}, true},
		{parse("example3.txt"), pos{6, 3}, true},
		{parse("example4.txt"), pos{4, 7}, true},
	}

	for _, tc := range testcases {
		got := isInside(tc.input, pipeLoop(tc.input, findStart(tc.input)), tc.current)

		if got != tc.expected {
			t.Fatalf("expected: %v, got: %v", tc.expected, got)
		}
	}
}
