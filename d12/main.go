package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

func parse(filepath string) (records []string, sizesList [][]int) {
	lines := utils.ReadLines(filepath)

	for _, line := range lines {
		chunks := strings.Split(line, " ")

		records = append(records, chunks[0])

		sizes := []int{}
		for _, nStr := range strings.Split(chunks[1], ",") {
			n, _ := strconv.Atoi(nStr)
			sizes = append(sizes, n)
		}

		sizesList = append(sizesList, sizes)
	}

	return records, sizesList
}

func hash(sizes []int) string {
	return fmt.Sprintf("%v", sizes)
}

func listCombinations(nGroups int, maxSum int) [][]int {
	combinations := [][]int{}
	visited := map[string]bool{}

	first := []int{}
	for i := 0; i < nGroups; i++ {
		first = append(first, 1)
	}
	first[0] = 0

	combinations = append(combinations, first)
	queue := append([][]int{}, first)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for i := 0; i < nGroups; i++ {
			new := append([]int{}, current...)
			new[i]++

			if visited[hash(new)] {
				continue
			}

			if utils.Sum(new) <= maxSum {
				combinations = append(combinations, new)
				queue = append(queue, new)
				visited[hash(new)] = true
			}
		}
	}

	return combinations
}

func combinationsIsValid(record string, sizes []int, combination []int) bool {
	p := ""

	for i := range sizes {
		for x := 0; x < combination[i]; x++ {
			p += "."
		}

		for x := 0; x < sizes[i]; x++ {
			p += "#"
		}
	}

	for len(p) < len(record) {
		p += "."
	}

	for i := range record {
		if record[i] != p[i] && record[i] != '?' {
			return false
		}
	}

	return true
}

func countCombinations(record string, sizes []int) int {
	sum := 0
	combinations := listCombinations(len(sizes), len(record)-utils.Sum(sizes))

	for _, c := range combinations {
		if combinationsIsValid(record, sizes, c) {
			sum++
		}
	}

	return sum
}

func part1(records []string, sizes [][]int) int {
	sum := 0
	for i := range records {
		sum += countCombinations(records[i], sizes[i])
	}

	return sum
}

func main() {
	records, sizes := parse(utils.Filepath())
	fmt.Println(part1(records, sizes))
}
