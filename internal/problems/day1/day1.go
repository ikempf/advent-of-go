package day1

import (
	"advent-of-go/internal/utils/io"
	"advent-of-go/internal/utils/parse"
	"fmt"
	"math"
	"sort"
	"strings"
)

func Day1() {
	fmt.Println("===Day 1 start===")

	first, second := readFiles()
	sort.Ints(first)
	sort.Ints(second)
	delta := computeDelta(first, second)

	fmt.Println("Day 1 answer: ", delta)

	fmt.Println("===Day 1 part two===")
	similarity := 0
	counts := groupAndCount(second)
	similarity = computeSimilarity(first, similarity, counts)

	fmt.Println("Day 1 answer: ", similarity)
}

func readFiles() ([]int, []int) {
	var first []int
	var second []int

	lines := io.ReadFiles("day1.txt")

	for i := range lines {
		line := lines[i]
		parts := strings.Split(line, "   ")
		first = append(first, parse.ParseInt(parts[0]))
		second = append(second, parse.ParseInt(parts[1]))
	}

	return first, second
}

func computeDelta(first []int, second []int) int {
	delta := 0
	for i := range first {
		delta += int(math.Abs(float64(first[i] - second[i])))
	}

	return delta
}

func computeSimilarity(first []int, similarity int, counts map[int]int) int {
	for i := range first {
		similarity += first[i] * counts[first[i]]
	}
	return similarity
}

func groupAndCount(elements []int) map[int]int {
	result := make(map[int]int)
	for _, element := range elements {
		result[element]++
	}
	return result
}
