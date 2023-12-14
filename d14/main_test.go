package main

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	original := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
	}

	got := rotate(original)

	expected := [][]string{
		{"4", "1"},
		{"5", "2"},
		{"6", "3"},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Fail()
	}
}

func Test_find_sequence(t *testing.T) {
	list := []int{0, 0, 0, 0, 1, 2, 3, 1, 2, 3, 0, 0, 0}

	gotIndex, gotLength := findSequence(list)
	expectedIndex, expectedLength := 4, 3

	if gotIndex != expectedIndex || gotLength != expectedLength {
		t.Fail()
	}
}
