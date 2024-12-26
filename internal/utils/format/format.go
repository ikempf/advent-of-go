package format

import "fmt"

func FormatDay(day int, result1 string, result2 string) {
	fmt.Printf("===Day %d start===\n", day)
	fmt.Printf("Day %d part 1: %s\n", day, result1)
	fmt.Printf("Day %d part 2: %s\n", day, result2)
	fmt.Printf("===Day %d end===\n", day)
}
