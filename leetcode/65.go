package leetcode

import "algorithm/jz"

// Valid Number
func isNumber(s string) bool {
	return jz.IsNumeric([]rune(s))
}
