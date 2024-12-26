package day1

import (
	"advent-of-go/internal/utils/format"
	"advent-of-go/internal/utils/io"
	"advent-of-go/internal/utils/parse"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Day1() {
	first, second := readLocations()
	sort.Ints(first)
	sort.Ints(second)
	delta := computeDelta(first, second)

	similarity := 0
	counts := groupAndCount(second)
	similarity = computeSimilarity(first, similarity, counts)

	format.FormatDay(1, strconv.Itoa(delta), strconv.Itoa(similarity))
}

func readLocations() ([]int, []int) {
	lines := io.ReadFile("day1.txt")
	first := make([]int, len(lines))
	second := make([]int, len(lines))

	for i := range lines {
		line := lines[i]
		parts := strings.Split(line, "   ")
		first[i] = parse.Int(parts[0])
		second[i] = parse.Int(parts[1])
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
