package main

import "testing"

func Test_position(t *testing.T) {
	testcases := []struct {
		holdTime int
		maxTime  int
		expected int
	}{
		{1, 7, 6},
		{2, 7, 10},
		{3, 7, 12},
		{4, 7, 12},
	}

	for _, tc := range testcases {
		got := position(tc.holdTime, tc.maxTime)

		if got != tc.expected {
			t.Fatalf("expected: %v, got: %v", tc.expected, got)
		}
	}
}
