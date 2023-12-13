package main

import "testing"

func Test_find_axis(t *testing.T) {
	example := parse("example1.txt")

	testcases := []struct {
		pattern   patternT
		expectedV int
		expectedH int
	}{
		{example[0], 5, 0},
		{example[1], 0, 4},
	}

	for _, tc := range testcases {
		gotV := findVAxis(tc.pattern)
		gotH := findHAxis(tc.pattern)

		if gotV != tc.expectedV || gotH != tc.expectedH {
			t.Fatalf("expected: %v, %v, got: %v, %v", tc.expectedV, tc.expectedH, gotV, gotH)
		}
	}
}
