package leetcode

import "math"

// Integer Break
// Given a positive integer n, break it into the sum of at least two positive integers and maximize the product of those integers. Return the maximum product you can get.
// Example 1:
//
//  Input: 2
//  Output: 1
//  Explanation: 2 = 1 + 1, 1 × 1 = 1.
//
// Example 2:
//  Input: 10
//  Output: 36
//  Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.
// Note: You may assume that n is not less than 2 and not larger than 58.
//
// 使用归纳法发现 将整数尽可能的拆分为3会得到最大的乘积
func IntegerBreak(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	result := 1
	for n > 4 {
		n -= 3
		result *= 3
	}

	result *= n

	return result
}

// 2 ->
//  i = 2, j = 1, result[2] = 1 * (2-1), 1 * result[2-1]
// 3 ->
//  i = 2, j = 1, result[2] = 1 * (2-1), 1 * result[2-1]
//  i = 3, j = 1, result[3] = result[3], 1 * (3-1), 1 * result[3-1]
//  i = 3, j = 2, result[3] = result[3], 2 * (3-2), 2 * result[3-2]
// 4 ->
//  i = 2, j = 1, result[2] = 1 * (2-1), 1 * result[2-1]
//  i = 3, j = 1, result[3] = result[3], 1 * (3-1), 1 * result[3-1]
//  i = 3, j = 2, result[3] = result[3], 2 * (3-2), 2 * result[3-2]
//  i = 4, j = 1, result[4] = result[4], 1 * (4-1), 1 * result[4-1]
//  i = 4, j = 2, result[4] = result[4], 2 * (4-2), 2 * result[4-2]
//  i = 4, j = 3, result[4] = result[4], 3 * (4-3), 3 * result[4-3]
//
// 将数分割成两部分 从1开始 依次去和另一部分去做乘积 依次去和另一部分得最大值做乘积 找到所有可能的最大值
func IntegerBreakII(n int) int {
	// calc the value from lowest
	results := make([]int, n+1)
	results[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i-1; j++ {
			// 每一次将i分割为j和i-j两部分，对于这两个部分计算我们可以分割的最大乘积
			// 将i-j继续分割的最大乘积就是memo[i-j],i-j一定小于i，所以memo[i-j]一定被计算出来了
			results[i] = max3(results[i], j*(i-j), j*results[i-j])
		}
	}

	return results[n]
}

func max3(a, b, c int) int {
	return int(math.Max(math.Max(float64(a), float64(b)), float64(c)))
}
