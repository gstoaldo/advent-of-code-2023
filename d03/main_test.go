package main

import "testing"

func Test_coordinate_is_inside_rectangle(t *testing.T) {
	testcases := []struct {
		c, topLeft, bottomRight coord
		expected                bool
	}{
		{coord{1, 1}, coord{0, 0}, coord{2, 2}, true},
		{coord{1, 1}, coord{2, 2}, coord{4, 4}, false},
	}

	for _, tc := range testcases {
		got := inside(tc.c, tc.topLeft, tc.bottomRight)

		if got != tc.expected {
			t.Fatalf("expected: %v, got: %v", tc.expected, got)
		}
	}
}
