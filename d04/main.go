package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

type numbersSet map[int]bool
type card struct{ winningSet, mySet numbersSet }

func parse(lines []string) []card {
	cards := []card{}

	re := regexp.MustCompile(`\d+`)

	for _, line := range lines {
		winningSet, mySet := numbersSet{}, numbersSet{}
		parts := strings.Split(line, "|")

		winningNumbersStr := re.FindAllString(parts[0], -1)[1:]
		for _, nStr := range winningNumbersStr {
			nInt, _ := strconv.Atoi(nStr)
			winningSet[nInt] = true
		}

		myNumbersStr := re.FindAllString(parts[1], -1)
		for _, nStr := range myNumbersStr {
			nInt, _ := strconv.Atoi(nStr)
			mySet[nInt] = true
		}

		cards = append(cards, card{winningSet, mySet})

	}

	return cards
}

func part1(cards []card) int {
	sum := 0

	for _, card := range cards {
		point := 0
		for n := range card.mySet {
			if card.winningSet[n] {
				if point == 0 {
					point = 1
				} else {
					point *= 2
				}
			}
		}

		sum += point
	}

	return sum
}

func part2(cards []card) int {
	copies := map[int]int{}

	for i, card := range cards {
		cardID := i + 1
		matches := 0
		for n := range card.mySet {
			if card.winningSet[n] {
				matches++
			}
		}

		for j := 1; j <= matches && matches > 0; j++ {
			copies[cardID+j] += 1 + copies[cardID]
		}
	}

	totalCards := len(cards) // including the original set of cards
	for _, ncopies := range copies {
		totalCards += ncopies
	}

	return totalCards
}

func main() {
	games := parse(utils.ReadLines(utils.Filepath()))
	p1 := part1(games)
	p2 := part2(games)

	utils.PrintSolution(p1, p2)
}
