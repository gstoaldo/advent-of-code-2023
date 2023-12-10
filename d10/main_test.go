package main

import (
	"reflect"
	"testing"
)

var example1 = parse("example1.txt")

func Test_start_neighbors(t *testing.T) {
	start := findStart(example1)
	got := startNeighbors(example1, start)

	expected := []pos{{1, 2}, {2, 1}}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected: %v, got: %v", expected, got)
	}
}
