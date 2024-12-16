package day2

import (
	"advent-of-go/internal/utils/io"
	"advent-of-go/internal/utils/parse"
	"fmt"
	"math"
	"strings"
)

const (
	Unknown    int = 0
	Increasing     = 1
	Decreasing     = 2
)

func Day2() {
	fmt.Println("===Day 2 start===")

	reports := readReports()
	safeCount := countSafeReports(reports, reportIsSafe)

	fmt.Printf("Day 2 answer, number of safe reports: %d\n", safeCount)

	fmt.Println("===Day 2 start part 2===")

	safeCountWithCorrection := countSafeReports(reports, reportIsSafeWithErrorCorrection)

	fmt.Printf("Day 2 answer part 2, number of safe reports: %d\n", safeCountWithCorrection)

}

func countSafeReports(reports [][]int, reportIsSafe func([]int) bool) int {
	safeCount := 0
	for i := range reports {
		if reportIsSafe(reports[i]) {
			safeCount += 1
		}
	}
	return safeCount
}

func readReports() [][]int {
	lines := io.ReadFile("day2.txt")
	reports := make([][]int, len(lines))

	for i := range lines {
		line := lines[i]
		levels := strings.Split(line, " ")
		for j := range levels {
			reports[i] = append(reports[i], parse.ParseInt(levels[j]))
		}
	}

	return reports
}

func reportIsSafe(report []int) bool {
	direction := Unknown

	for i := range report[:len(report)-1] {
		if !isSafe(report[i], report[i+1], &direction) {
			return false
		}
	}

	return true
}

func reportIsSafeWithErrorCorrection(report []int) bool {
	direction := Unknown

	for i := range report[:len(report)-2] {
		abIsSafe := isSafe(report[i], report[i+1], &direction)
		bcIsSafe := isSafe(report[i+1], report[i+2], &direction)
		acIsSafe := isSafe(report[i], report[i+2], &direction)
		if !abIsSafe && !bcIsSafe && !acIsSafe {
			return false
		}
	}

	return true
}

func isSafe(n int, m int, direction *int) bool {
	if math.Abs(float64(n-m)) > 3 {
		return false
	}
	if n < m {
		if *direction == Unknown {
			*direction = Increasing
		}
		if *direction == Decreasing {
			return false
		}
	} else if n > m {
		if *direction == Unknown {
			*direction = Decreasing
		}
		if *direction == Increasing {
			return false
		}
	} else {
		return false
	}
	return true
}
