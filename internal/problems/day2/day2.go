package day2

import (
	"advent-of-go/internal/utils/format"
	"advent-of-go/internal/utils/io"
	"advent-of-go/internal/utils/parse"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	Unknown    int = 0
	Increasing int = 1
	Decreasing int = 2
)

func Day2() {
	reports := readReports()
	safeCount := countSafeReports(reports, reportIsSafe)
	safeCountWithCorrection := countSafeReports(reports, reportIsSafeWithErrorCorrection)
	format.FormatDay(2, strconv.Itoa(safeCount), strconv.Itoa(safeCountWithCorrection))
}

func countSafeReports(reports [][]int, reportIsSafe func([]int) bool) int {
	safeCount := 0
	for i := range reports {
		if reportIsSafe(reports[i]) {
			safeCount++
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
			reports[i] = append(reports[i], parse.Int(levels[j]))
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
	fmt.Println("before", report, reportIsSafe(report))
	for i := range report {
		reportCopy := append([]int{}, report...)
		reportSlice := append(reportCopy[:i], reportCopy[i+1:]...)
		if reportIsSafe(reportSlice) {
			return true
		}
	}

	return false
}

func reportIsSafeWithErrorCorrectionOptimized(report []int) bool {
	direction := Unknown
	droppedAnElement := false

	for i := range report[:len(report)-2] {
		abIsSafe := isSafe(report[i], report[i+1], &direction)
		if !abIsSafe {
			if droppedAnElement {
				return false
			}
			droppedAnElement = true
			acIsSafe := isSafe(report[i], report[i+2], &direction)
			if !acIsSafe {
				return acIsSafe
			}
		}
	}

	return true
}

func isSafe(n int, m int, direction *int) bool {
	if math.Abs(float64(n-m)) > 3 {
		return false
	}
	switch {
	case n < m:
		if *direction == Unknown {
			*direction = Increasing
		}
		if *direction == Decreasing {
			return false
		}
	case n > m:
		if *direction == Unknown {
			*direction = Decreasing
		}
		if *direction == Increasing {
			return false
		}
	default:
		return false
	}
	return true
}
