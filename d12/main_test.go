package main

import "testing"

func Test_list_combinations(t *testing.T) {
	testcases := []struct {
		nGroups   int
		maxLength int
		expected  int
	}{
		{3, 6, 83},
	}

	for _, tc := range testcases {
		got := listCombinations(tc.nGroups, tc.maxLength)

		if len(got) != tc.expected {
			t.Fatalf("expected: %v, got: %v", tc.expected, len(got))
		}
	}
}

func Test_combinations_is_valid(t *testing.T) {
	testcases := []struct {
		record      string
		sizes       []int
		combination []int
		expected    bool
	}{
		{"?###????????", []int{3, 2, 1}, []int{0, 1, 1}, false},
		{"?###????????", []int{3, 2, 1}, []int{1, 1, 1}, true},
		{"???.###", []int{1, 1, 3}, []int{0, 1, 1}, true},
	}

	for _, tc := range testcases {
		got := combinationsIsValid(tc.record, tc.sizes, tc.combination)

		if got != tc.expected {
			t.Fatalf("expected: %v, got: %v", tc.expected, got)
		}
	}
}

func Test_count_combinations(t *testing.T) {
	var records, sizes = parse("example1.txt")

	testcases := []struct {
		record   string
		sizes    []int
		expected int
	}{
		{records[0], sizes[0], 1},
		{records[1], sizes[1], 4},
		{records[2], sizes[2], 1},
		{records[3], sizes[3], 1},
		{records[4], sizes[4], 4},
		{records[5], sizes[5], 10},
	}

	for _, tc := range testcases {
		t.Run("", func(t *testing.T) {
			got := countCombinations(tc.record, tc.sizes)

			if got != tc.expected {
				t.Fatalf("expected: %v, got: %v", tc.expected, got)
			}
		})
	}
}
