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

func neighbors(grid gridT, curr locT, canClimbSlope bool) []locT {
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

		if grid[next.i][next.j] == "#" {
			continue
		}

		if grid[next.i][next.j] == "." {
			result = append(result, next)
			continue
		}

		if grid[next.i][next.j] == slope || canClimbSlope {
			result = append(result, next)
		}
	}

	return result
}

func copyMap(visited map[locT]bool) map[locT]bool {
	cp := map[locT]bool{}
	for k, v := range visited {
		cp[k] = v
	}
	return cp
}

func findIntersections(grid gridT) map[locT]bool {
	result := map[locT]bool{}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "#" {
				continue
			}

			curr := locT{i, j}
			if len(neighbors(grid, curr, true)) >= 3 {
				result[curr] = true
			}
		}
	}
	return result
}

type headT struct {
	curr   locT
	nsteps int
}

// distance from one intersection (key) to all other connected intersections (value)
func compress(grid gridT, start, finish locT, canClimb bool) map[locT][]headT {
	intersections := findIntersections(grid)
	intersections[start] = true
	intersections[finish] = true

	result := map[locT][]headT{}

	for intersection := range intersections {
		queue := []headT{{curr: intersection}}
		visited := map[locT]bool{}

		for len(queue) > 0 {
			head := queue[0]
			queue = queue[1:]

			if intersections[head.curr] && head.curr != intersection {
				result[intersection] = append(result[intersection], head)
				continue
			}

			visited[head.curr] = true

			for _, next := range neighbors(grid, head.curr, canClimb) {
				if visited[next] {
					continue
				}

				queue = append(queue, headT{
					curr:   next,
					nsteps: head.nsteps + 1,
				})
			}
		}
	}

	return result
}

func search(grid gridT, start, finish locT, canClimb bool) int {
	compressed := compress(grid, start, finish, canClimb)

	type queueItem struct {
		curr    locT
		visited map[locT]bool
		nsteps  int
	}

	queue := []queueItem{
		{
			curr:    start,
			visited: map[locT]bool{},
			nsteps:  0,
		},
	}

	max := 0

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		if head.curr == finish {
			max = utils.Max(max, head.nsteps)
			continue
		}

		head.visited[head.curr] = true

		for _, next := range compressed[head.curr] {
			if head.visited[next.curr] {
				continue
			}

			queue = append(queue, queueItem{
				curr:    next.curr,
				visited: copyMap(head.visited),
				nsteps:  head.nsteps + next.nsteps,
			})
		}
	}

	return max
}

func part1(grid gridT, start, finish locT) int {
	return search(grid, start, finish, false)
}

func part2(grid gridT, start, finish locT) int {
	defer utils.Timer()()
	return search(grid, start, finish, true)
}

func main() {
	grid, start, finish := parse(utils.Filepath())
	fmt.Println(part1(grid, start, finish))
	fmt.Println(part2(grid, start, finish))
}
