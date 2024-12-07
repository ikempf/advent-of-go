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

	fmt.Println("====Day 1 start====")

	var first []int
	var second []int
	var delta int = 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")
		first = append(first, parse(parts[0]))
		second = append(second, parse(parts[1]))
	}

	sort.Ints(first)
	sort.Ints(second)

	for i := 0; i < len(first); i++ {
		delta += int(math.Abs(float64(first[i] - second[i])))
	}

	println("Day 1 answer: ", delta)
}

func parse(value string) int {
	val, _ := strconv.Atoi(value)
	return val
}
