package main

import (
	"regexp"
	"strconv"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

func parse(filepath string) ([]int, []int) {
	lines := utils.ReadLines(filepath)

	re := regexp.MustCompile(`\d+`)

	timesStr := re.FindAllString(lines[0], -1)
	times := []int{}
	for _, s := range timesStr {
		i, _ := strconv.Atoi(s)
		times = append(times, i)
	}

	distsStr := re.FindAllString(lines[1], -1)
	dists := []int{}
	for _, s := range distsStr {
		i, _ := strconv.Atoi(s)
		dists = append(dists, i)
	}

	return times, dists
}

func position(holdTime, maxTime int) int {
	a := 1
	return holdTime * a * (maxTime - holdTime)
}

func part1(times, dists []int) int {
	allNumberOfWays := []int{}

	for i, maxTime := range times {
		numberOfWays := 0
		for holdTime := 1; holdTime <= maxTime; holdTime++ {
			if position(holdTime, maxTime) > dists[i] {
				numberOfWays++
			}
		}
		allNumberOfWays = append(allNumberOfWays, numberOfWays)
	}

	result := 1

	for _, v := range allNumberOfWays {
		result *= v
	}

	return result
}

func main() {
	times, dists := parse(utils.Filepath())
	p1 := part1(times, dists)
	p2 := 0

	utils.PrintSolution(p1, p2)
}
