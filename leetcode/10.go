package leetcode

import "algorithm/jz"

// Regular Expression Matching
func isMatch(s string, p string) bool {
	return jz.RegMatch([]rune(s), []rune(p))
}
