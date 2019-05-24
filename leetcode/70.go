package leetcode

import "algorithm/jz"

func ClimbStairs(N int) int {
	return int(jz.JumpFloor(uint(N)))
}
