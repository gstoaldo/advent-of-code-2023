package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

var example1 = utils.ReadLines("example1.txt")
var example2 = utils.ReadLines("example2.txt")

func TestCalibrationValue(t *testing.T) {
	testcases := []struct {
		line     string
		expected int
	}{
		{example1[0], 12},
		{example1[1], 38},
		{example1[2], 15},
		{example1[3], 77},
	}

	for _, tc := range testcases {
		got := calibrationValue(tc.line)

		if got != tc.expected {
			t.Fatalf("expected: %v, got: %v", tc.expected, got)
		}
	}
}

func TestTranslate(t *testing.T) {
	testcases := []struct {
		line     string
		expected string
	}{
		{example2[0], "219"},
		{example2[1], "823"},
		{example2[2], "123"},
		{"twoone", "21"},
		{"twone", "21"},
	}

	for _, tc := range testcases {
		got := translate(tc.line)

		if got != tc.expected {
			t.Fatalf("expected: %v, got: %v", tc.expected, got)
		}
	}
}
