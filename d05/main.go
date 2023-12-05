package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

type seed int
type mapRange []int

func parse(filepath string) ([]seed, [][]mapRange) {
	file := utils.ReadFile(filepath)

	re := regexp.MustCompile(`\d+`)

	chunks := strings.Split(file, "\n\n")

	seeds := []seed{}
	seedsStr := re.FindAllString(chunks[0], -1)
	for _, vStr := range seedsStr {
		vInt, _ := strconv.Atoi(vStr)
		seeds = append(seeds, seed(vInt))
	}

	maps := [][]mapRange{}
	for _, chunk := range chunks[1:] {
		mapRanges := []mapRange{}
		for _, line := range strings.Split(chunk, "\n")[1:] {
			mapRange := mapRange{}
			numbersStr := re.FindAllString(line, -1)
			for _, vStr := range numbersStr {
				vInt, _ := strconv.Atoi(vStr)
				mapRange = append(mapRange, vInt)
			}
			mapRanges = append(mapRanges, mapRange)
		}

		maps = append(maps, mapRanges)
	}

	return seeds, maps
}

func isInRange(n int, mapRange mapRange) bool {
	return n >= mapRange[1] && n < mapRange[1]+mapRange[2]
}

func convert(n int, mapRanges []mapRange) int {
	for _, mapRange := range mapRanges {
		if isInRange(n, mapRange) {
			return n - mapRange[1] + mapRange[0]
		}
	}

	return n
}

func part1(seeds []seed, maps [][]mapRange) int {
	min := math.MaxInt
	for _, seed := range seeds {
		source := int(seed)
		for _, mapRanges := range maps {
			source = convert(source, mapRanges)
		}
		min = utils.Min(min, source)
	}

	return min
}

func main() {
	seeds, maps := parse(utils.Filepath())
	p1 := part1(seeds, maps)
	p2 := 0

	utils.PrintSolution(p1, p2)
}
