package main

import (
	"reflect"
	"testing"
)

func Test_handRank(t *testing.T) {
	testcases := []struct {
		hand     string
		expected []int
	}{
		{"AAAAA", []int{7, 13, 13, 13, 13, 13}},
		{"AA8AA", []int{6, 13, 13, 7, 13, 13}},
		{"23332", []int{5, 1, 2, 2, 2, 1}},
		{"TTT98", []int{4, 9, 9, 9, 8, 7}},
		{"23432", []int{3, 1, 2, 3, 2, 1}},
		{"A23A4", []int{2, 13, 1, 2, 13, 3}},
		{"23456", []int{1, 1, 2, 3, 4, 5}},
		{"KTJJT", []int{3, 12, 9, 10, 10, 9}},
		{"KK677", []int{3, 12, 12, 5, 6, 6}},
	}

	for _, tc := range testcases {
		got := handRank(tc.hand, getCardRanksMap(cardRanksP1))

		if !reflect.DeepEqual(got, tc.expected) {
			t.Fatalf("expected: %v, got: %v", tc.expected, got)
		}
	}
}

func Test_is_stronger(t *testing.T) {
	testcases := []struct {
		handA    string
		handB    string
		expected bool
	}{
		{"AAAAA", "AA8AA", true},
		{"AA8AA", "AAAAA", false},
		{"33332", "2AAAA", true},
		{"77888", "77788", true},
		{"KK677", "KTJJT", true},
		{"KTJJT", "KK677", false},
	}

	for _, tc := range testcases {
		got := isStronger(tc.handA, tc.handB)

		if !reflect.DeepEqual(got, tc.expected) {
			t.Fatalf("expected: %v, got: %v", tc.expected, got)
		}
	}
}

func Test_max_possible(t *testing.T) {
	testcases := []struct {
		hand     string
		expected int
	}{
		{"32T3K", 2},
		{"KK677", 3},
		{"T55J5", 6},
		{"KTJJT", 6},
		{"QQQJA", 6},
	}

	for _, tc := range testcases {
		got := maxPossible(tc.hand)

		if !reflect.DeepEqual(got, tc.expected) {
			t.Fatalf("expected: %v, got: %v", tc.expected, got)
		}
	}
}
