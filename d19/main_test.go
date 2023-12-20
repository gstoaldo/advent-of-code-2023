package main

import (
	"reflect"
	"testing"
)

func Test_split_ranges_less_than(t *testing.T) {
	rule := []string{"s", "<", "2770"}
	ranges := [][]int{{0, 4000}, {0, 4000}, {0, 4000}, {0, 4000}}

	gotA, gotB := splitRanges(ranges, rule)

	expectedA := [][]int{{0, 4000}, {0, 4000}, {0, 4000}, {0, 2769}}
	expectedB := [][]int{{0, 4000}, {0, 4000}, {0, 4000}, {2771, 4000}}

	if !reflect.DeepEqual(gotA, expectedA) || !reflect.DeepEqual(gotB, expectedB) {
		t.Fail()
	}
}

func Test_split_ranges_greater_than(t *testing.T) {
	rule := []string{"s", ">", "2770"}
	ranges := [][]int{{0, 4000}, {0, 4000}, {0, 4000}, {0, 4000}}

	gotA, gotB := splitRanges(ranges, rule)

	expectedA := [][]int{{0, 4000}, {0, 4000}, {0, 4000}, {2771, 4000}}
	expectedB := [][]int{{0, 4000}, {0, 4000}, {0, 4000}, {0, 2769}}

	if !reflect.DeepEqual(gotA, expectedA) || !reflect.DeepEqual(gotB, expectedB) {
		t.Fail()
	}
}
