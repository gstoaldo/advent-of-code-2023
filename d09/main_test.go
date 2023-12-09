package main

import "testing"

var example1 = parse("example1.txt")

func Test_extrapolate(t *testing.T) {
	testcases := []struct {
		sequence []int
		expected int
	}{
		{example1[0], 18},
		{example1[1], 28},
		{example1[2], 68},
	}

	for _, tc := range testcases {
		got := extrapolate(tc.sequence, false)

		if got != tc.expected {
			t.Fatalf("expected: %v, got: %v", tc.expected, got)
		}
	}
}

func Test_extrapolate_backwards(t *testing.T) {
	testcases := []struct {
		sequence []int
		expected int
	}{
		{example1[0], -3},
		{example1[1], 0},
		{example1[2], 5},
	}

	for _, tc := range testcases {
		got := extrapolate(tc.sequence, true)

		if got != tc.expected {
			t.Fatalf("expected: %v, got: %v", tc.expected, got)
		}
	}
}
