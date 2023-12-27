package main

import "testing"

func equal(a, b float64) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}

	return diff < 0.001
}

func Test_intersection(t *testing.T) {
	a := stoneT{19, 13, 30, -2, 1, -2}
	b := stoneT{18, 19, 22, -1, -1, -2}

	gotX := intersectionX(a, b)
	expectedX := 14.333

	gotY := a.f(gotX)
	expectedY := 15.333

	if !equal(gotX, expectedX) {
		t.Errorf("expected: %v, got: %v", expectedX, gotX)
	}

	if !equal(gotY, expectedY) {
		t.Errorf("expected: %v, got: %v", expectedY, gotY)
	}
}
