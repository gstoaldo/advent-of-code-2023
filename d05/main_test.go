package main

import (
	"testing"
)

var seeds, maps = parse("example1.txt")

func Test_value_is_in_range(t *testing.T) {
	testcases := []struct {
		mapRange mapRange
		source   int
		expected bool
	}{
		{mapRange{50, 98, 2}, 98, true},
		{mapRange{50, 98, 2}, 99, true},
		{mapRange{50, 98, 2}, 100, false},
	}

	for _, tc := range testcases {
		got := isInRange(tc.source, tc.mapRange)

		if got != tc.expected {
			t.Fatalf("expected: %v, got: %v", tc.expected, got)
		}
	}
}

func Test_convert(t *testing.T) {
	testcases := []struct {
		mapRanges []mapRange
		source    int
		expected  int
	}{
		{maps[0], 79, 81},
		{maps[1], 81, 81},
		{maps[2], 81, 81},
		{maps[3], 81, 74},
		{maps[4], 74, 78},
		{maps[5], 78, 78},
		{maps[6], 78, 82},
	}

	for _, tc := range testcases {
		got := convert(tc.source, tc.mapRanges)

		if got != tc.expected {
			t.Fatalf("expected: %v, got: %v", tc.expected, got)
		}
	}
}
