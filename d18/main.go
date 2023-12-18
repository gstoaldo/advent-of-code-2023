package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

type stepT struct {
	dir    string
	length int
	color  string
}

type edgeT struct {
	x0, y0, x1, y1 int
	color          string
}

type coordT struct{ x, y int }

func parse(filepath string) (result []stepT) {
	re := regexp.MustCompile(`(\w).(\d+)..(#\w+)`)

	for _, line := range utils.ReadLines(filepath) {
		matches := re.FindStringSubmatch(line)
		length, _ := strconv.Atoi(matches[2])

		result = append(result, stepT{
			dir:    matches[1],
			length: length,
			color:  matches[3],
		})
	}

	return result
}

var dirToDeltas = map[string]struct{ dx, dy int }{
	"U": {0, -1},
	"D": {0, 1},
	"R": {1, 0},
	"L": {-1, 0},
}

func edges(steps []stepT) (result []edgeT) {
	x, y := 0, 0
	for _, step := range steps {
		x1 := x + dirToDeltas[step.dir].dx*step.length
		y1 := y + dirToDeltas[step.dir].dy*step.length

		result = append(result, edgeT{
			x0:    x,
			x1:    x1,
			y0:    y,
			y1:    y1,
			color: step.color,
		})

		x, y = x1, y1
	}

	return result
}

func polygonArea(edges []edgeT) int {
	coords := []coordT{}
	for _, edge := range edges {
		coords = append(coords, coordT{edge.x0, edge.y0})
	}

	result := 0
	for i := 1; i < len(coords); i++ {
		result += (coords[i-1].x + coords[i].x) * (coords[i-1].y - coords[i].y)
	}

	return utils.Abs(result) / 2
}

func digArea(steps []stepT) int {
	perimeter := 0
	for _, step := range steps {
		perimeter += step.length
	}

	return perimeter/2 + polygonArea(edges(steps)) + 1
}

func part1(steps []stepT) int {
	return digArea(steps)
}

func main() {
	steps := parse(utils.Filepath())
	fmt.Println(part1(steps))
}
