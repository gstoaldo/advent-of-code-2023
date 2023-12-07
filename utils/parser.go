package utils

import (
	"os"
	"strings"
)

func Filepath() string {
	filepath := "input.txt"

	if len(os.Args) == 2 {
		filepath = os.Args[1]
	}

	return filepath
}

func ReadFile(filepath string) string {
	file, err := os.ReadFile(filepath)
	if err != nil {
		panic("error reading file")
	}

	return string(file)
}

func ReadLines(filepath string) []string {
	file := ReadFile(filepath)

	return strings.Split(string(file), "\n")
}
