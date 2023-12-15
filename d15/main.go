package main

import (
	"fmt"
	"regexp"
	"strconv"
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

type lensT struct {
	label       string
	focalLength int
}

func remove(hashmap map[int][]lensT, boxID int, lens lensT) {
	boxContent := hashmap[boxID]

	newContent := []lensT{}

	for _, l := range boxContent {
		if l.label != lens.label {
			newContent = append(newContent, l)
		}
	}

	hashmap[boxID] = newContent
}

func addOrReplace(hashmap map[int][]lensT, boxID int, lens lensT) {
	boxContent := hashmap[boxID]

	newContent := []lensT{}
	found := false

	for _, l := range boxContent {
		if l.label == lens.label {
			newContent = append(newContent, lens)
			found = true
		} else {
			newContent = append(newContent, l)
		}
	}

	if !found {
		newContent = append(newContent, lens)
	}

	hashmap[boxID] = newContent
}

func getHashmap(instructions []string) map[int][]lensT {
	hashmap := map[int][]lensT{}

	re := regexp.MustCompile(`(\w+)([=|-])(\d)?`)

	for _, instruction := range instructions {
		groups := re.FindStringSubmatch(instruction)

		boxID := hash(groups[1])
		lens := lensT{label: groups[1]}
		operation := groups[2]

		if operation == "-" {
			remove(hashmap, boxID, lens)
		}

		if operation == "=" {
			lens.focalLength, _ = strconv.Atoi(groups[3])
			addOrReplace(hashmap, boxID, lens)
		}
	}

	return hashmap
}

func focusingPower(hashmap map[int][]lensT) int {
	result := 0

	for boxID, content := range hashmap {
		for i, l := range content {
			result += (boxID + 1) * (i + 1) * l.focalLength
		}
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

func part2(instructions []string) int {
	return focusingPower(getHashmap(instructions))
}

func main() {
	instructions := parse(utils.Filepath())
	fmt.Println(part1(instructions))
	fmt.Println(part2(instructions))
}
