package main

import (
	"fmt"
	"strings"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

type patternT []string

func parse(filepath string) (patterns []patternT) {
	for _, block := range strings.Split(utils.ReadFile(filepath), "\n\n") {
		patterns = append(patterns, strings.Split(block, "\n"))
	}
	return patterns
}

func transpose(pattern patternT) patternT {
	transposed := patternT{}

	for j := 0; j < len(pattern[0]); j++ {
		line := make([]rune, len(pattern))
		for i := 0; i < len(pattern); i++ {
			line[i] = rune(pattern[i][j])
		}

		transposed = append(transposed, string(line))
	}

	return transposed
}

func findHAxis(pattern patternT) int {
	for x := 1; x < len(pattern); x++ {
		minLen := utils.Min(len(pattern)-x, x)

		isAxis := true

		for i := 0; i < minLen; i++ {
			if pattern[x-1-i] != pattern[x+i] {
				isAxis = false
				break
			}
		}

		if isAxis {
			return x
		}
	}

	return 0
}

func findVAxis(pattern patternT) int {
	transposed := transpose(pattern)

	return findHAxis(transposed)
}

func part1(patterns []patternT) int {
	sum := 0
	for _, p := range patterns {
		vAxis := findVAxis(p)
		hAxis := findHAxis(p)

		sum += vAxis + 100*hAxis
	}

	return sum
}

func main() {
	patterns := parse(utils.Filepath())
	fmt.Println(part1(patterns))
}
