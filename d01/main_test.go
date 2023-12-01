package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

var example1 = utils.ReadLines("example1.txt")
var example2 = utils.ReadLines("example2.txt")

func TestCalibrationValueP1(t *testing.T) {
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
		got := calibrationValueP1(tc.line)

		if got != tc.expected {
			t.Fatalf("expected: %v, got: %v", tc.expected, got)
		}
	}
}

func TestCalibrationValueP2(t *testing.T) {
	testcases := []struct {
		line     string
		expected int
	}{
		{example2[0], 29},
		{example2[1], 83},
		{example2[2], 13},
		{example2[3], 24},
		{example2[4], 42},
		{example2[5], 14},
		{example2[6], 76},
		{"rphtbkncs4nznsix", 46},
		{"nineeighttworhtvxdtxp8twoneh", 91},
		{"63eightsixgdsdqqxzzsbnkt782", 62},
		{"three6fivefoursixgtzfzbkhmnplfm", 36},
		{"fourgtwopbjbcvgtwo3one", 41},
		{"5three1", 51},
		{"nine671seventwotwonejkf", 91},
		{"twone", 21},
		{"21", 21},
		{"shrzvdcghblt21", 21},
	}

	for _, tc := range testcases {
		got := calibrationValueP2(tc.line)

		if got != tc.expected {
			t.Fatalf("expected: %v, got: %v", tc.expected, got)
		}
	}
}
