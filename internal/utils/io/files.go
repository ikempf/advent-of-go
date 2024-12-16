package io

import (
	"bufio"
	"os"
)

func ReadFiles(name string) []string {
	input, _ := os.Open("internal/input/" + name)
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
