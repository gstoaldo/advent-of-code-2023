package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

type partT []int // x, m, a, s
var categoryToIndex = map[string]int{
	"x": 0,
	"m": 1,
	"a": 2,
	"s": 3,
}

type workflowT struct {
	rules        [][]string // category, operator, value
	destinations []string
	fallback     string
}

func checkRule(category, operator, valueStr string, part partT) bool {
	if operator == "<" {
		return part[categoryToIndex[category]] < utils.ToInt(valueStr)
	}

	return part[categoryToIndex[category]] > utils.ToInt(valueStr)
}

func parse(filepath string) (map[string]workflowT, []partT) {
	workflows := map[string]workflowT{}
	parts := []partT{}

	chunks := strings.Split(utils.ReadFile(filepath), "\n\n")

	reOrigin := regexp.MustCompile(`\w+`)
	reExp := regexp.MustCompile(`((\w+)([<|>])(\d+)):(\w+)`)
	reFallback := regexp.MustCompile(`(\w+)}`)

	for _, line := range strings.Split(chunks[0], "\n") {
		origin := reOrigin.FindString(line)
		fallback := reFallback.FindStringSubmatch(line)[1]

		rules := [][]string{}
		destinations := []string{}
		for _, ruleValues := range reExp.FindAllStringSubmatch(line, -1) {
			rules = append(rules, []string{ruleValues[2], ruleValues[3], ruleValues[4]})
			destinations = append(destinations, ruleValues[5])
		}

		workflows[origin] = workflowT{
			rules:        rules,
			destinations: destinations,
			fallback:     fallback,
		}
	}

	reParts := regexp.MustCompile(`\d+`)
	for _, line := range strings.Split(chunks[1], "\n") {
		part := partT{}
		for _, s := range reParts.FindAllString(line, -1) {
			part = append(part, utils.ToInt(s))
		}
		parts = append(parts, part)
	}

	return workflows, parts
}

func acceptPart(origin string, workflows map[string]workflowT, part partT) bool {
	wf := workflows[origin]

	for i, rule := range wf.rules {
		passRule, destination := checkRule(rule[0], rule[1], rule[2], part), wf.destinations[i]

		if passRule && destination == "A" {
			return true
		}

		if passRule && destination == "R" {
			return false
		}

		if passRule {
			return acceptPart(destination, workflows, part)
		}
	}

	if wf.fallback == "A" {
		return true
	}

	if wf.fallback == "R" {
		return false
	}

	return acceptPart(wf.fallback, workflows, part)
}

func sortParts(workflows map[string]workflowT, parts []partT) []partT {
	acceptedParts := []partT{}
	for _, part := range parts {
		if acceptPart("in", workflows, part) {
			acceptedParts = append(acceptedParts, part)
		}
	}

	return acceptedParts
}

func part1(workflows map[string]workflowT, parts []partT) int {
	sum := 0

	for _, parts := range sortParts(workflows, parts) {
		for _, category := range parts {
			sum += category
		}
	}

	return sum
}

func splitRanges(ranges [][]int, rule []string) ([][]int, [][]int) {
	ruleCheckSplit, otherSplit := [][]int{}, [][]int{}

	index := categoryToIndex[rule[0]]

	for i, r := range ranges {
		if i != index {
			ruleCheckSplit = append(ruleCheckSplit, r)
			otherSplit = append(otherSplit, r)
		} else {
			limit := utils.ToInt(rule[2])

			if limit < r[0] || limit > r[1] {
				panic("edge case, limit out of bounds")
			}

			if rule[1] == "<" {
				ruleCheckSplit = append(ruleCheckSplit, []int{r[0], limit - 1})
				otherSplit = append(otherSplit, []int{limit, r[1]})
			}

			if rule[1] == ">" {
				ruleCheckSplit = append(ruleCheckSplit, []int{limit + 1, r[1]})
				otherSplit = append(otherSplit, []int{r[0], limit})
			}
		}
	}

	return ruleCheckSplit, otherSplit
}

func acceptRanges(origin string, workflows map[string]workflowT, ranges [][]int) int {
	sum := 0
	wf := workflows[origin]
	for i, rule := range wf.rules {
		splitA, splitB := splitRanges(ranges, rule)
		ranges = splitB

		if wf.destinations[i] == "A" {
			sum += countCombinations(splitA)
			continue
		}

		if wf.destinations[i] == "R" {
			continue
		}

		sum += acceptRanges(wf.destinations[i], workflows, splitA)
	}

	if wf.fallback == "A" {
		return sum + countCombinations(ranges)
	}

	if wf.fallback == "R" {
		return sum
	}

	return sum + acceptRanges(wf.fallback, workflows, ranges)
}

func countCombinations(ranges [][]int) int {
	result := 1

	for _, r := range ranges {
		result *= utils.Abs(r[0]-r[1]) + 1
	}

	return result
}

func part2(workflows map[string]workflowT, parts []partT) int {
	ranges := [][]int{{1, 4000}, {1, 4000}, {1, 4000}, {1, 4000}}
	return acceptRanges("in", workflows, ranges)
}

func main() {
	workflows, parts := parse(utils.Filepath())
	fmt.Println(part1(workflows, parts))
	fmt.Println(part2(workflows, parts))
}
