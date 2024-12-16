package parse

import "strconv"

func ParseInt(value string) int {
	val, _ := strconv.Atoi(value)
	return val
}
