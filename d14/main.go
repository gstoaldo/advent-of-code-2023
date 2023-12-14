package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

func parse(filepath string) (input [][]string) {
	for _, line := range utils.ReadLines(filepath) {
		input = append(input, strings.Split(line, ""))
	}
	return input
}

func canMoveNorth(input [][]string, i, j int) bool {
	return i > 0 && input[i][j] == "O" && input[i-1][j] == "."
}

func moveNorth(input [][]string, i, j int) {
	input[i][j] = "."
	input[i-1][j] = "O"
}

func tiltNorth(input [][]string) [][]string {
	tilted := [][]string{}
	for _, row := range input {
		tilted = append(tilted, append([]string{}, row...))
	}

	for j := 0; j < len(tilted[0]); j++ {
		for i := 1; i < len(tilted); i++ {
			if canMoveNorth(tilted, i, j) {
				moveNorth(tilted, i, j)
				i = -2
			}
		}
	}

	return tilted
}

func totalLoad(state [][]string) int {
	result := 0

	for i, row := range state {
		for _, v := range row {
			if v == "O" {
				result += len(state) - i
			}
		}
	}

	return result
}

func rotate(state [][]string) [][]string {
	rotated := [][]string{}
	for _ = range state[0] {
		rotated = append(rotated, make([]string, len(state)))
	}

	for i := 0; i < len(rotated); i++ {
		for j := 0; j < len(rotated[0]); j++ {
			w := len(rotated[0])
			rotated[i][j] = state[w-1-j][i]
		}
	}

	return rotated
}

func cycle(state [][]string) [][]string {
	for i := 0; i < 4; i++ {
		state = tiltNorth(state)
		state = rotate(state)
	}

	return state
}

func part1(input [][]string) int {
	return totalLoad(tiltNorth(input))
}

func findSequence(loadHistory []int) (int, int) {
	for L := len(loadHistory) - 1; L > 1; L-- {
		for i := 0; i+2*L < len(loadHistory); i++ {
			if reflect.DeepEqual(loadHistory[i:i+L], loadHistory[i+L:i+2*L]) {
				return i, L
			}
		}
	}

	return 0, 0
}

func part2(input [][]string) int {
	state := input
	loadHistory := []int{}

	ncycles := 1_000_000_000
	sequenceFound := false

	for i := 0; i < ncycles; i++ {
		state = cycle(state)
		loadHistory = append(loadHistory, totalLoad(state))

		_, sequenceLength := findSequence(loadHistory)

		if sequenceLength > 0 && !sequenceFound {
			sequenceFound = true
			N := (ncycles - i) / sequenceLength
			i += N * sequenceLength
		}
	}

	return totalLoad(state)
}

func main() {
	input := parse(utils.Filepath())
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
