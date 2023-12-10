package main

import (
	"fmt"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

type inputT []string
type pos struct{ i, j int }

var shapeNeighbors = map[rune][][]int{
	'S': {{-1, 0}, {0, 1}, {1, 0}, {0, -1}},
	'|': {{-1, 0}, {1, 0}},
	'-': {{0, 1}, {0, -1}},
	'L': {{-1, 0}, {0, 1}},
	'J': {{-1, 0}, {0, -1}},
	'7': {{1, 0}, {0, -1}},
	'F': {{1, 0}, {0, 1}},
}

func parse(filepath string) inputT {
	return utils.ReadLines(filepath)
}

func findStart(input inputT) pos {
	for i, line := range input {
		for j, c := range line {
			if c == 'S' {
				return pos{i, j}
			}
		}
	}
	panic("starting position not found")
}

func neighbors(input inputT, current pos) []pos {
	result := []pos{}
	for _, n := range shapeNeighbors[rune(input[current.i][current.j])] {
		ni, nj := current.i+n[0], current.j+n[1]

		if ni < 0 || ni >= len(input) || nj < 0 || nj >= len(input[0]) {
			continue
		}

		result = append(result, pos{ni, nj})
	}
	return result
}

func startNeighbors(input inputT, startPos pos) []pos {
	result := []pos{}

	for _, startNeighbor := range neighbors(input, startPos) {
		if input[startNeighbor.i][startNeighbor.j] == '.' {
			continue
		}
		for _, n := range neighbors(input, startNeighbor) {
			if n == startPos {
				result = append(result, startNeighbor)
			}
		}
	}

	return result
}

func pipeLength(input inputT, startPos pos) int {
	visited := map[pos]bool{}
	path := []pos{}

	visited[startPos] = true
	path = append(path, startPos)

	current := startNeighbors(input, startPos)[0]

	for current != startPos {
		visited[current] = true
		path = append(path, current)

		for _, n := range neighbors(input, current) {
			if visited[n] && n != startPos {
				continue
			}

			current = n
			break
		}

	}

	return len(path)
}

func part1(input inputT) int {
	return pipeLength(input, findStart(input)) / 2
}

func main() {
	input := parse(utils.Filepath())
	fmt.Println(part1(input))
}
