package main

import (
	"fmt"
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

func part1(input [][]string) int {
	tilted := tiltNorth(input)
	// for _, row := range tilted {
	// 	fmt.Println(row)
	// }

	return totalLoad(tilted)
}

func main() {
	input := parse(utils.Filepath())
	fmt.Println(part1(input))
}
