package main

import (
	"fmt"
	"regexp"
	"strings"

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

func pipeLoop(input inputT, startPos pos) map[pos]bool {
	visited := map[pos]bool{}
	visited[startPos] = true

	current := startNeighbors(input, startPos)[0]
	for current != startPos {
		visited[current] = true

		for _, n := range neighbors(input, current) {
			if visited[n] && n != startPos {
				continue
			}

			current = n
			break
		}
	}

	return visited
}

func part1(input inputT) int {
	return len(pipeLoop(input, findStart(input))) / 2
}

func filterRowPipesToTheRight(current pos, input inputT, loop map[pos]bool) string {
	result := ""
	for j := current.j; j < len(input[0]); j++ {
		if loop[pos{current.i, j}] {
			result += string(input[current.i][j])
		}
	}

	return result
}

func replaceTurns(pipesInRow string) string {
	exps := []struct {
		re  *regexp.Regexp
		new string
	}{
		{regexp.MustCompile(`L-*J`), "||"},
		{regexp.MustCompile(`F-*7`), "||"},
		{regexp.MustCompile(`F-*J`), "|"},
		{regexp.MustCompile(`L-*7`), "|"},
	}

	result := pipesInRow

	for _, exp := range exps {
		result = exp.re.ReplaceAllString(result, exp.new)
	}

	return result
}

func isInside(input inputT, loop map[pos]bool, current pos) bool {
	// A tile is enclosed by the loop if, going in any direction, the number of
	// pipes crossed is odd.

	// I pick a tile (current) and look all the tiles to the right. Then replace
	// U turns by "||" and Z turns by "|" (check image).
	replaced := replaceTurns(filterRowPipesToTheRight(current, input, loop))

	return len(replaced)%2 != 0
}

func part2(input inputT, startShape string) int {
	count := 0

	start := findStart(input)
	_pipeLoop := pipeLoop(input, start)

	input[start.i] = strings.Replace(input[start.i], "S", startShape, -1)

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			current := pos{i, j}
			if _pipeLoop[current] {
				continue
			}

			if isInside(input, _pipeLoop, current) {
				count++
			}
		}
	}

	return count
}

func main() {
	input := parse(utils.Filepath())
	fmt.Println(part1(input))

	// TODO: find starting position pipe shape
	// example 4, S = F
	// input, S = L
	fmt.Println(part2(input, "L"))
}
