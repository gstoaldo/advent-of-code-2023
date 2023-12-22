package main

import (
	"reflect"
	"testing"
)

func Test_generateCubes(t *testing.T) {
	brick := brickT{pointT{0, 0, 4}, pointT{0, 2, 4}}

	got := generateCubes(brick)

	expected := []pointT{{0, 0, 4}, {0, 1, 4}, {0, 2, 4}}

	if !reflect.DeepEqual(got, expected) {
		t.Fail()
	}
}

func Test_do_collides(t *testing.T) {
	brick1 := brickT{pointT{0, 1, 1}, pointT{2, 1, 1}}
	brick2 := brickT{pointT{1, 0, 1}, pointT{1, 2, 1}}

	if !collides(brick1, brick2) {
		t.Fail()
	}
}

func Test_do_not_collides(t *testing.T) {
	brick1 := brickT{pointT{0, 1, 1}, pointT{2, 1, 1}}
	brick2 := brickT{pointT{1, 2, 1}, pointT{1, 4, 1}}

	if collides(brick1, brick2) {
		t.Fail()
	}
}
