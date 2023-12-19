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
	exps         []func(item partT) bool
	destinations []string
	fallback     string
}

func condition(category, operator, valueStr string) func(item partT) bool {
	return func(item partT) bool {
		if operator == "<" {
			return item[categoryToIndex[category]] < utils.ToInt(valueStr)
		}

		return item[categoryToIndex[category]] > utils.ToInt(valueStr)
	}
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

		exps := []func(part partT) bool{}
		destinations := []string{}
		for _, expMatch := range reExp.FindAllStringSubmatch(line, -1) {
			exps = append(exps, condition(expMatch[2], expMatch[3], expMatch[4]))
			destinations = append(destinations, expMatch[5])
		}

		workflows[origin] = workflowT{
			exps:         exps,
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

	for i, exp := range wf.exps {
		if exp(part) && wf.destinations[i] == "A" {
			return true
		}

		if exp(part) && wf.destinations[i] == "R" {
			return false
		}

		if exp(part) {
			return acceptPart(wf.destinations[i], workflows, part)
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

func main() {
	workflows, parts := parse(utils.Filepath())
	fmt.Println(part1(workflows, parts))
}
