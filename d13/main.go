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

func findHAxis(pattern patternT, fix bool) int {
	for x := 1; x < len(pattern); x++ {
		fixedCount := 1
		if fix {
			fixedCount = 0
		}

		minLen := utils.Min(len(pattern)-x, x)
		isAxis := true

		for i := 0; i < minLen; i++ {
			a, b := pattern[x-1-i], pattern[x+i]
			fixedCount += countDiff(a, b)

			if a != b && fixedCount > 1 {
				isAxis = false
				break
			}
		}

		if isAxis && fixedCount == 1 {
			return x
		}
	}

	return 0
}

func findVAxis(pattern patternT, fix bool) int {
	return findHAxis(transpose(pattern), fix)
}

func countDiff(a, b string) int {
	if a == b {
		return 0
	}

	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count
}

func summarize(patterns []patternT, fix bool) int {
	sum := 0
	for _, p := range patterns {
		vAxis := findVAxis(p, fix)
		hAxis := findHAxis(p, fix)

		sum += vAxis + 100*hAxis
	}

	return sum
}

func part1(patterns []patternT) int {
	return summarize(patterns, false)
}

func part2(patterns []patternT) int {
	return summarize(patterns, true)
}

func main() {
	patterns := parse(utils.Filepath())
	fmt.Println(part1(patterns))
	fmt.Println(part2(patterns))
}
