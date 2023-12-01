package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

type inputT []string

func calibrationValue(line string) int {
	re := regexp.MustCompile(`\d`)

	matches := re.FindAllString(line, -1)
	numberStr := ""

	if len(matches) == 1 {
		numberStr = matches[0] + matches[0]
	} else {
		numberStr = matches[0] + matches[len(matches)-1]
	}

	number, _ := strconv.Atoi(numberStr)

	return number
}

func part1(input inputT) int {
	sum := 0

	for _, line := range input {
		sum += calibrationValue(line)
	}

	fmt.Println(sum)

	return sum
}

func main() {
	input := utils.ReadLines("d01/input.txt")

	part1(input)
}
