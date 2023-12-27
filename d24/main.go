package main

import (
	"fmt"
	"regexp"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

type stoneT struct {
	x0, y0, z0 int
	vx, vy, vz int
}

// line slope
func (r stoneT) m() float64 {
	return float64(r.vy) / float64(r.vx)
}

// line equation
func (r stoneT) f(x float64) float64 {
	return float64(r.y0) + r.m()*(x-float64(r.x0))
}

func (r stoneT) time(x float64) float64 {
	return (x - float64(r.x0)) / float64(r.vx)
}

// fa(x) = fb(x)
func intersectionX(a, b stoneT) float64 {
	return (float64(b.y0-a.y0) + a.m()*float64(a.x0) - b.m()*float64(b.x0)) / (a.m() - b.m())
}

func parse(filepath string) []stoneT {
	result := []stoneT{}
	re := regexp.MustCompile(`-?\d+`)
	for _, line := range utils.ReadLines(filepath) {
		matches := re.FindAllString(line, -1)

		result = append(result, stoneT{
			x0: utils.ToInt(matches[0]),
			y0: utils.ToInt(matches[1]),
			z0: utils.ToInt(matches[2]),
			vx: utils.ToInt(matches[3]),
			vy: utils.ToInt(matches[4]),
			vz: utils.ToInt(matches[5]),
		})
	}

	return result
}

func futureCross(stoneA, stoneB stoneT) bool {
	min, max := float64(200_000_000_000_000), float64(400_000_000_000_000)
	// min, max := float64(7), float64(27) // example

	// paths are parallel
	if stoneA.m() == stoneB.m() {
		return false
	}

	x := intersectionX(stoneA, stoneB)
	y := stoneA.f(x)

	// paths crossed in the past
	if stoneA.time(x) < 0 || stoneB.time(x) < 0 {
		return false
	}

	return x >= min && x <= max && y >= min && y <= max
}

func part1(stones []stoneT) int {
	defer utils.Timer()()
	count := 0
	for i, stoneA := range stones {
		for _, stoneB := range stones[i+1:] {
			if futureCross(stoneA, stoneB) {
				count++
			}
		}
	}
	return count
}

func main() {
	stones := parse(utils.Filepath())
	fmt.Println(part1(stones))
}
