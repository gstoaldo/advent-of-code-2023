package main

import (
	"reflect"
	"testing"
)

var example = parse("example1.txt")

func Test_expasion(t *testing.T) {

	rows, cols := rowAndColToExpand(example)
	expectedRows := []int{3, 7}
	expectedCols := []int{2, 5, 8}

	if !reflect.DeepEqual(rows, expectedRows) {
		t.Fatalf("expected: %v, got: %v", expectedRows, rows)
	}

	if !reflect.DeepEqual(cols, expectedCols) {
		t.Fatalf("expected: %v, got: %v", expectedCols, cols)
	}
}

func Test_distance(t *testing.T) {
	testcases := []struct {
		galaxyA  pos
		galaxyB  pos
		expected int
	}{
		{pos{5, 1}, pos{9, 4}, 9},
		{pos{2, 0}, pos{6, 9}, 17},
	}

	rowsToExpand, colsToExpand := rowAndColToExpand(example)

	for _, tc := range testcases {
		got := dist(tc.galaxyA, tc.galaxyB, rowsToExpand, colsToExpand, 2)

		if got != tc.expected {
			t.Fatalf("expected: %v, got: %v", tc.expected, got)
		}
	}
}
