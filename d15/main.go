package main

import (
	"fmt"
	"strings"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

func parse(filepath string) []string {
	return strings.Split(utils.ReadFile(filepath), ",")
}

func hash(instruction string) int {
	result := 0

	for _, char := range instruction {
		result += int(char)
		result *= 17
		result = result % 256
	}

	return result
}

func part1(instructions []string) int {
	result := 0

	for _, instruction := range instructions {
		result += hash(instruction)
	}

	return result
}

func main() {
	instructions := parse(utils.Filepath())
	fmt.Println(part1(instructions))
}
