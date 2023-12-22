package main

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

type pointT struct {
	x, y, z int
}

type brickT struct {
	name   string
	p0, p1 pointT
}

func parse(filepath string) []brickT {
	result := []brickT{}
	re := regexp.MustCompile(`\d+`)

	for i, line := range utils.ReadLines(filepath) {
		matches := re.FindAllString(line, -1)

		p0 := pointT{
			x: utils.ToInt(matches[0]),
			y: utils.ToInt(matches[1]),
			z: utils.ToInt(matches[2]),
		}

		p1 := pointT{
			x: utils.ToInt(matches[3]),
			y: utils.ToInt(matches[4]),
			z: utils.ToInt(matches[5]),
		}

		result = append(result, brickT{fmt.Sprintf("%v", i+1), p0, p1})
	}

	return result
}

func printStack(stack []brickT) {
	cubes := []pointT{}

	for _, b := range stack {
		cubes = append(cubes, generateCubes(b)...)
	}

	W, H := 0, 0

	for _, c := range cubes {
		W = utils.Max(W, c.x)
		H = utils.Max(H, c.z)
	}

	grid := [][]string{}
	for z := 0; z <= H; z++ {
		row := []string{}
		for x := 0; x <= W; x++ {
			if z == 0 {
				row = append(row, "-")
			} else {
				row = append(row, ".")
			}

		}
		grid = append(grid, row)
	}

	for _, c := range cubes {
		grid[c.z][c.x] = "#"
	}

	for i := range grid {
		fmt.Println(grid[len(grid)-1-i])
	}

	fmt.Println()
}

func sortByZ(bricks []brickT) []brickT {
	cp := append([]brickT{}, bricks...)

	sort.Slice(cp, func(i, j int) bool {
		return utils.Min(cp[i].p0.z, cp[i].p1.z) < utils.Min(cp[j].p0.z, cp[j].p1.z)
	})

	return cp
}

func generateCubes(brick brickT) []pointT {
	delta := pointT{brick.p1.x - brick.p0.x, brick.p1.y - brick.p0.y, brick.p1.z - brick.p0.z}
	delta.x = delta.x / utils.Max(utils.Abs(delta.x), 1)
	delta.y = delta.y / utils.Max(utils.Abs(delta.y), 1)
	delta.z = delta.z / utils.Max(utils.Abs(delta.z), 1)

	result := []pointT{}
	curr := brick.p0
	for curr != brick.p1 {
		result = append(result, curr)
		curr = pointT{curr.x + delta.x, curr.y + delta.y, curr.z + delta.z}
	}
	result = append(result, curr)

	return result
}

func collides(brick1 brickT, brick2 brickT) bool {
	brick1Cubes := generateCubes(brick1)
	brick2Cubes := generateCubes(brick2)

	brick2CubesMap := map[pointT]bool{}
	for _, b := range brick2Cubes {
		brick2CubesMap[b] = true
	}

	for _, b := range brick1Cubes {
		if brick2CubesMap[b] {
			return true
		}
	}

	return false
}

func getCollisions(brick brickT, stack []brickT) (bool, []brickT) {
	result := []brickT{}
	if brick.p0.z == 0 || brick.p1.z == 0 {
		return true, result
	}

	for _, b := range stack {
		if collides(brick, b) {
			result = append(result, b)
		}
	}

	return len(result) > 0, result
}

func moveDown(brick brickT) brickT {
	p0 := brick.p0
	p0.z--
	p1 := brick.p1
	p1.z--

	return brickT{brick.name, p0, p1}
}

func simulate(bricks []brickT) ([]brickT, map[brickT][]brickT) {
	queue := sortByZ(bricks)
	stack := []brickT{}

	brickToSupports := map[brickT][]brickT{}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		fmt.Printf("queue: %v\n", len(queue))

		for {
			moved := moveDown(curr)
			doCollides, collisions := getCollisions(moved, stack)

			if doCollides {
				brickToSupports[curr] = collisions
				stack = append(stack, curr)
				break
			} else {
				curr = moved
			}
		}
	}

	// printStack(append(queue, stack...))
	return stack, brickToSupports
}

func getNotSafeToDesintegrateBricks(bricks []brickT, supportsMap map[brickT][]brickT) map[brickT]bool {
	notSafe := map[brickT]bool{}

	for _, supports := range supportsMap {
		if len(supports) == 1 {
			notSafe[supports[0]] = true
		}
	}

	return notSafe
}

func chainReaction(bricks []brickT, brickToSupports map[brickT][]brickT, start brickT) int {
	chain := map[brickT]bool{}

	supportToBricks := map[brickT][]brickT{}
	for b, supports := range brickToSupports {
		for _, s := range supports {
			supportToBricks[s] = append(supportToBricks[s], b)
		}
	}

	chain[start] = true
	queue := append([]brickT{}, supportToBricks[start]...)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		supports := brickToSupports[curr]
		allSupportsInChain := true
		for _, s := range supports {
			if !chain[s] {
				allSupportsInChain = false
			}
		}

		if allSupportsInChain {
			chain[curr] = true
			queue = append(queue, supportToBricks[curr]...)
		}
	}

	return len(chain) - 1
}

func part1(bricks []brickT) int {
	_, supports := simulate(bricks)
	notSafe := getNotSafeToDesintegrateBricks(bricks, supports)

	return len(bricks) - len(notSafe)
}

func part2(bricks []brickT) int {
	defer utils.Timer()()

	sum := 0
	stack, supports := simulate(bricks)
	notSafe := getNotSafeToDesintegrateBricks(bricks, supports)

	i := 0
	for brick := range notSafe {
		fmt.Printf("chain: %v/%v\n", i, len(notSafe))
		sum += chainReaction(stack, supports, brick)
		i++
	}

	return sum

}

func main() {
	bricks := parse(utils.Filepath())
	fmt.Println(part1(bricks))
	fmt.Println(part2(bricks))
}
