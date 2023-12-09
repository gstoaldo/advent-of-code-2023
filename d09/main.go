package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

func parse(filepath string) [][]int {
	input := [][]int{}

	for _, line := range utils.ReadLines(filepath) {
		history := []int{}
		for _, nStr := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(nStr)
			history = append(history, n)
		}
		input = append(input, history)
	}

	return input
}

func allZeros(sequence []int) bool {
	for _, n := range sequence {
		if n != 0 {
			return false
		}
	}
	return true
}

func sequenceDiff(sequence []int) []int {
	result := []int{}

	for i := 1; i < len(sequence); i++ {
		result = append(result, sequence[i]-sequence[i-1])
	}

	return result
}

func extrapolate(sequence []int, backwards bool) int {
	if allZeros(sequence) {
		return 0
	}

	nextValue := extrapolate(sequenceDiff(sequence), backwards)

	if backwards {
		return sequence[0] - nextValue
	}

	return sequence[len(sequence)-1] + nextValue
}

func sumExtrapolation(input [][]int, backwards bool) int {
	sum := 0

	for _, sequence := range input {
		sum += extrapolate(sequence, backwards)
	}

	return sum
}

func part1(input [][]int) int {
	return sumExtrapolation(input, false)
}

func part2(input [][]int) int {
	return sumExtrapolation(input, true)
}

func main() {
	input := parse(utils.Filepath())
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
