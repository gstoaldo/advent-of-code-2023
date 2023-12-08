package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

type nodesT map[string][]string

var lrToIndex = map[string]int{"L": 0, "R": 1}

func parse(filepath string) (string, nodesT) {
	file := utils.ReadFile(filepath)

	chunks := strings.Split(file, "\n\n")

	instructions := chunks[0]

	nodes := nodesT{}
	re := regexp.MustCompile(`\w+`)
	for _, line := range strings.Split(chunks[1], "\n") {
		matches := re.FindAllString(line, -1)
		nodes[matches[0]] = []string{matches[1], matches[2]}
	}

	return instructions, nodes
}

func part1(instructions string, nodes nodesT) int {
	nsteps := 0
	current := "AAA"

	for current != "ZZZ" {
		direction := string(instructions[nsteps%len(instructions)])
		current = nodes[current][lrToIndex[direction]]

		nsteps++
	}

	return nsteps
}

func main() {
	instructions, nodes := parse(utils.Filepath())
	fmt.Println(part1(instructions, nodes))
}
