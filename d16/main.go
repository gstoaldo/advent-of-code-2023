package main

import (
	"fmt"
	"strings"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

type gridT [][]string
type headT struct{ i, j, vi, vj int }

func parse(filepath string) gridT {
	result := [][]string{}

	for _, line := range utils.ReadLines(filepath) {
		result = append(result, strings.Split(line, ""))
	}

	return result
}

func bump(curr string, head headT) []headT {
	f := 1
	if curr == "/" {
		f = -1
	}

	return []headT{
		{
			head.i, head.j, head.vj * f, head.vi * f,
		},
	}
}

func split(curr string, head headT) []headT {
	if curr == "-" && head.vi == 0 {
		return []headT{head}
	}

	if curr == "|" && head.vj == 0 {
		return []headT{head}
	}

	return []headT{
		{
			head.i, head.j, head.vj, head.vi,
		},
		{
			head.i, head.j, head.vj * -1, head.vi * -1,
		},
	}
}

func inGrid(grid gridT, head headT) bool {
	return head.i >= 0 && head.i < len(grid) && head.j >= 0 && head.j < len(grid[0])
}

func step(grid gridT, head headT) []headT {
	heads := []headT{}
	curr := grid[head.i][head.j]

	if curr == "|" || curr == "-" {
		heads = split(curr, head)
	}

	if curr == "\\" || curr == "/" {
		heads = bump(curr, head)
	}

	if curr == "." {
		heads = append(heads, head)
	}

	nextHeads := []headT{}
	for _, head := range heads {
		next := headT{head.i + head.vi, head.j + head.vj, head.vi, head.vj}

		if inGrid(grid, next) {
			nextHeads = append(nextHeads, next)
		}
	}

	return nextHeads
}

func run(grid gridT) map[headT]bool {
	visited := map[headT]bool{}
	queue := []headT{{0, 0, 0, 1}}

	for len(queue) > 0 {
		// printGrid(grid, visited)
		head := queue[0]
		queue = queue[1:]
		visited[head] = true

		for _, nextHead := range step(grid, head) {
			if !visited[nextHead] {
				queue = append(queue, nextHead)
			}
		}
	}

	return visited
}

func countEnergizedTiles(visited map[headT]bool) int {
	type pos struct{ i, j int }

	tilesVisited := map[pos]bool{}

	for k := range visited {
		tilesVisited[pos{k.i, k.j}] = true
	}

	return len(tilesVisited)
}

func part1(grid gridT) int {
	return countEnergizedTiles(run(grid))
}

func printGrid(grid gridT, visited map[headT]bool) {
	visitedGrid := map[struct{ i, j int }]bool{}
	for k := range visited {
		visitedGrid[struct{ i, j int }{k.i, k.j}] = true
	}

	for i, l := range grid {
		for j, v := range l {
			if visitedGrid[struct{ i, j int }{i, j}] {
				fmt.Printf("X")
			} else {
				fmt.Printf("%v", v)
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n\n")
}

func main() {
	grid := parse(utils.Filepath())
	fmt.Println(part1(grid))
}
