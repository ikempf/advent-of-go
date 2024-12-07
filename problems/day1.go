package problems

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day1() {
	input, _ := os.Open("input/day1.txt")
	defer input.Close()

	fmt.Println("===Day 1 start===")
	var first []int
	var second []int

	first, second = readFiles(input, first, second)
	sort.Ints(first)
	sort.Ints(second)
	var delta = computeDelta(first, second)

	fmt.Println("Day 1 answer: ", delta)

	fmt.Println("===Day 1 part two===")
	var similarity int = 0
	counts := groupAndCount(second)
	similarity = computeSimilarity(first, similarity, counts)

	fmt.Println("Day 1 answer: ", similarity)
}

func readFiles(input *os.File, first []int, second []int) ([]int, []int) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")
		first = append(first, parse(parts[0]))
		second = append(second, parse(parts[1]))
	}
	return first, second
}

func parse(value string) int {
	val, _ := strconv.Atoi(value)
	return val
}

func computeDelta(first []int, second []int) int {
	var delta int = 0
	for i := 0; i < len(first); i++ {
		delta += int(math.Abs(float64(first[i] - second[i])))
	}
	return delta
}

func computeSimilarity(first []int, similarity int, counts map[int]int) int {
	for i := 0; i < len(first); i++ {
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
