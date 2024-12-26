package parse

import "strconv"

func Int(value string) int {
	val, _ := strconv.Atoi(value)
	return val
}
