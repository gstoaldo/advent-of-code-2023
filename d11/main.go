package main

import (
	"fmt"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

type inputT []string
type pos struct{ i, j int }

func parse(filepath string) inputT {
	return utils.ReadLines(filepath)
}

func rowAndColToExpand(input inputT) (rows, cols []int) {
	rowsWithGalaxies, colsWithGalaxies := map[int]bool{}, map[int]bool{}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] == '#' {
				rowsWithGalaxies[i] = true
				colsWithGalaxies[j] = true
			}
		}
	}

	for i := 0; i < len(input); i++ {
		if !rowsWithGalaxies[i] {
			rows = append(rows, i)
		}
	}

	for j := 0; j < len(input[0]); j++ {
		if !colsWithGalaxies[j] {
			cols = append(cols, j)
		}
	}

	return rows, cols
}

func listGalaxies(input inputT) (result []pos) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] == '#' {
				result = append(result, pos{i, j})
			}
		}
	}
	return result
}

func countExpasions(a, b int, colsToExpand []int) (result int) {
	start := utils.Min(a, b)
	end := utils.Max(a, b)

	for _, i := range colsToExpand {
		if i > start && i < end {
			result++
		}
	}

	return result
}

func dist(galaxyA, galaxyB pos, rowsToExpand, colsToExpand []int, expansionSize int) int {
	dx := utils.Abs(galaxyB.j - galaxyA.j)
	dy := utils.Abs(galaxyB.i - galaxyA.i)

	dxExp := countExpasions(galaxyB.j, galaxyA.j, colsToExpand)
	dyExp := countExpasions(galaxyB.i, galaxyA.i, rowsToExpand)

	return dx + dy + (dxExp+dyExp)*(expansionSize-1)
}

func sumDists(input inputT, expansionSize int) int {
	sum := 0
	rowsToExpand, colsToExpand := rowAndColToExpand(input)
	galaxies := listGalaxies(input)

	for a := 0; a < len(galaxies); a++ {
		for b := a + 1; b < len(galaxies); b++ {
			sum += dist(galaxies[a], galaxies[b], rowsToExpand, colsToExpand, expansionSize)
		}
	}

	return sum
}

func part1(input inputT) int {
	return sumDists(input, 2)
}

func part2(input inputT) int {
	return sumDists(input, 1000000)
}

func main() {
	input := parse(utils.Filepath())
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
