package main

import "testing"

func Test_find_axis(t *testing.T) {
	example := parse("example1.txt")

	testcases := []struct {
		pattern   patternT
		fix       bool
		expectedV int
		expectedH int
	}{
		{example[0], false, 5, 0},
		{example[1], false, 0, 4},
		{example[0], true, 0, 3},
		{example[1], true, 0, 1},
	}

	for _, tc := range testcases {
		t.Run("", func(t *testing.T) {
			gotV := findVAxis(tc.pattern, tc.fix)
			gotH := findHAxis(tc.pattern, tc.fix)

			if gotV != tc.expectedV || gotH != tc.expectedH {
				t.Fatalf("expected: %v, %v, got: %v, %v", tc.expectedV, tc.expectedH, gotV, gotH)
			}
		})
	}
}
