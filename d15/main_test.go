package main

import "testing"

func Test_hash(t *testing.T) {
	expected := 52
	got := hash("HASH")

	if got != expected {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}
