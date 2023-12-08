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

func next(current string, nsteps int, instructions string, nodes nodesT) string {
	direction := string(instructions[nsteps%len(instructions)])
	return nodes[current][lrToIndex[direction]]
}

func part1(instructions string, nodes nodesT) int {
	nsteps := 0
	current := "AAA"

	for current != "ZZZ" {
		current = next(current, nsteps, instructions, nodes)
		nsteps++
	}

	return nsteps
}

func startList(nodes nodesT) []string {
	result := []string{}
	for n := range nodes {
		if string(n[2]) == "A" {
			result = append(result, n)
		}
	}
	return result
}

func isExit(node string) bool {
	return node[2] == 'Z'
}

func calcFrequency(node string, instructions string, nodes nodesT) int {
	nsteps := 0
	current := node

	for nsteps < 50_000 {
		current = next(current, nsteps, instructions, nodes)

		if isExit(current) {
			return nsteps + 1
		}

		nsteps++
	}

	panic(fmt.Sprintf("frequency not found for node %v", node))
}

func part2(instructions string, nodes nodesT) []int {
	multiples := []int{}
	for _, c := range startList(nodes) {
		multiples = append(multiples, calcFrequency(c, instructions, nodes))
	}

	return multiples
}

func main() {
	instructions, nodes := parse(utils.Filepath())
	fmt.Println(part1(instructions, nodes))
	fmt.Println(part2(instructions, nodes)) // TODO: calculate least commom multiple
}
