package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

type inputT []string

var mapToDigit = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}

func calibrationValue(line string) int {
	re := regexp.MustCompile(`\d`)

	matches := re.FindAllString(line, -1)
	numberStr := ""

	if len(matches) == 0 {
		return 0
	}

	if len(matches) == 1 {
		numberStr = matches[0] + matches[0]
	} else {
		numberStr = matches[0] + matches[len(matches)-1]
	}

	number, _ := strconv.Atoi(numberStr)

	return number
}

func translate(line string) string {
	result := ""

	for i := range line {
		for key, value := range mapToDigit {
			if i+len(key) > len(line) {
				continue
			}

			substring := line[i : i+len(key)]

			if substring == key {
				result += value
			}
		}
	}

	return result
}

func part1(input inputT) int {
	sum := 0

	for _, line := range input {
		sum += calibrationValue(line)
	}

	fmt.Println(sum)

	return sum
}

func part2(input inputT) int {
	sum := 0

	for _, line := range input {
		sum += calibrationValue(translate(line))
	}

	fmt.Println(sum)

	return sum
}

func main() {
	input := utils.ReadLines("input.txt")

	part1(input)
	part2(input)
}
