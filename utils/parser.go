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

func ReadFile(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic("error reading file")
	}

	return string(file)
}

func ReadLines(path string) []string {
	file := ReadFile(path)

	return strings.Split(string(file), "\n")
}
