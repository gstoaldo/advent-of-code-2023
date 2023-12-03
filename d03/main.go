package main

import (
	"regexp"
	"strconv"
	"unicode"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

type inputT []string
type number struct {
	i     int
	j0    int
	j1    int
	value int
}
type coord struct{ i, j int }

func parse(input inputT) ([]number, []coord) {
	numbers := []number{}
	symbols := []coord{}
	re := regexp.MustCompile(`\d+`)

	for i, line := range input {
		matches := re.FindAllStringIndex(line, -1)

		for _, match := range matches {
			valueInt, _ := strconv.Atoi(line[match[0]:match[1]])

			numbers = append(numbers, number{
				i:     i,
				j0:    match[0],
				j1:    match[1] - 1,
				value: valueInt,
			})
		}
	}

	for i, line := range input {
		for j, char := range line {
			if !unicode.IsDigit(char) && char != '.' {
				symbols = append(symbols, coord{i, j})
			}
		}
	}

	return numbers, symbols
}

func rect(n number) (coord, coord) {
	return coord{n.i - 1, n.j0 - 1}, coord{n.i + 1, n.j1 + 1}
}

func inside(c, topLeft, bottomRight coord) bool {
	return c.i >= topLeft.i && c.i <= bottomRight.i && c.j >= topLeft.j && c.j <= bottomRight.j
}

func part1(input inputT) int {
	numbers, symbols := parse(input)

	sum := 0

	for _, number := range numbers {
		for _, symbol := range symbols {
			topLeft, bottomRight := rect(number)
			if inside(symbol, topLeft, bottomRight) {
				sum += number.value
			}
		}
	}

	return sum
}

func part2(input inputT) int {
	numbers, symbols := parse(input)

	sum := 0

	for _, symbol := range symbols {
		if input[symbol.i][symbol.j] != '*' {
			continue
		}

		gearNumbers := []int{}

		for _, number := range numbers {
			topLeft, bottomRight := rect(number)
			if inside(symbol, topLeft, bottomRight) {
				gearNumbers = append(gearNumbers, number.value)
			}
		}

		if len(gearNumbers) == 2 {
			sum += gearNumbers[0] * gearNumbers[1]
		}
	}

	return sum
}

func main() {
	input := utils.ReadLines(utils.Filepath())
	p1 := part1(input)
	p2 := part2(input)

	utils.PrintSolution(p1, p2)
}
