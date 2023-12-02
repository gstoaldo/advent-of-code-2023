package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

type inputT [][][]int // RGB

var colorToIndex = map[string]int{
	"red":   0,
	"green": 1,
	"blue":  2,
}

func parse(lines []string) inputT {
	result := inputT{}

	re := regexp.MustCompile(`(\d+) (red|green|blue)`)

	for _, line := range lines {
		game := [][]int{}
		sets := strings.Split(line, ";")

		for _, set := range sets {
			matches := re.FindAllStringSubmatch(set, -1)

			colors := [3]int{}

			for _, match := range matches {
				number, _ := strconv.Atoi(match[1])
				colors[colorToIndex[match[2]]] = number
			}

			game = append(game, colors[:])
		}

		result = append(result, game)
	}

	return result
}

func gameIsValid(game [][]int, config []int) bool {
	valid := true
	for _, set := range game {
		for colorID := range set {
			valid = valid && set[colorID] <= config[colorID]
		}
	}

	return valid
}

func part1(input inputT) int {
	result := 0
	config := []int{12, 13, 14}

	for i, game := range input {
		gameID := i + 1
		valid := gameIsValid(game, config)

		if valid {
			result += gameID
		}
	}

	return result
}

func gameMinConfig(game [][]int) []int {
	result := [3]int{}

	for _, set := range game {
		for colorID := range set {
			result[colorID] = utils.Max(result[colorID], set[colorID])
		}
	}

	return result[:]
}

func power(set []int) int {
	result := 1
	for _, color := range set {
		result *= color
	}

	return result
}

func part2(input inputT) int {
	result := 0

	for _, game := range input {
		power := power(gameMinConfig(game))
		result += power
	}

	return result
}

func main() {
	input := parse(utils.ReadLines(utils.Filepath()))

	p1 := part1(input)
	p2 := part2(input)

	fmt.Printf("Part 1: %v\nPart 2: %v\n", p1, p2)
}
