package main

import (
	"fmt"
	"strings"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

type gridT [][]string
type locT struct{ i, j int }

func parse(filepath string) (gridT, locT, locT) {
	grid := gridT{}

	for _, line := range utils.ReadLines(filepath) {
		grid = append(grid, strings.Split(line, ""))
	}

	i := 0
	var start locT
	for j := range grid[i] {
		if grid[i][j] == "." {
			start = locT{i, j}
		}
	}

	i = len(grid) - 1
	var finish locT
	for j := range grid[i] {
		if grid[i][j] == "." {
			finish = locT{i, j}
		}
	}

	return grid, start, finish
}

func neighbors(grid gridT, curr locT) []locT {
	maxI, maxJ := len(grid)-1, len(grid[0])-1
	deltas := map[locT]string{
		{1, 0}:  "v",
		{-1, 0}: "^",
		{0, 1}:  ">",
		{0, -1}: "<",
	}

	result := []locT{}
	for d, slope := range deltas {
		next := locT{curr.i + d.i, curr.j + d.j}

		if next.i < 0 || next.i > maxI || next.j < 0 || next.j > maxJ {
			continue
		}

		if grid[next.i][next.j] == slope || grid[next.i][next.j] == "." {
			result = append(result, next)
		}
	}

	return result
}

func search(grid gridT, start, finish locT) int {
	type headT struct {
		curr   locT
		prev   locT
		nsteps int
	}

	paths := []headT{}

	queue := []headT{{
		curr: start,
	}}

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		if head.curr == finish {
			paths = append(paths, head)
			continue
		}

		for _, next := range neighbors(grid, head.curr) {
			if next.i == head.prev.i && next.j == head.prev.j {
				continue
			}

			queue = append(queue, headT{
				curr:   next,
				prev:   head.curr,
				nsteps: head.nsteps + 1,
			})
		}
	}

	max := 0
	for _, p := range paths {
		max = utils.Max(max, p.nsteps)
	}

	return max
}

func part1(grid gridT, start, finish locT) int {
	return search(grid, start, finish)
}

func main() {
	grid, start, finish := parse(utils.Filepath())
	fmt.Println(part1(grid, start, finish))
}
