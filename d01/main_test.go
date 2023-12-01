package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

var example = utils.ReadLines("example.txt")

func TestCalibrationValue(t *testing.T) {
	testcases := []struct {
		line     string
		expected int
	}{
		{example[0], 12},
		{example[1], 38},
		{example[2], 15},
		{example[3], 77},
	}

	for _, tc := range testcases {
		got := calibrationValue(tc.line)

		if got != tc.expected {
			t.Fatalf("expected: %v, got: %v", tc.expected, got)
		}
	}
}
