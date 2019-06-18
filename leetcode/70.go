package leetcode

import "algorithm/jz"

// Climbing Stairs
func ClimbStairs(N int) int {
	return int(jz.JumpFloor(uint(N)))
}
