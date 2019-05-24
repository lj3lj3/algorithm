package jz

import "math"

// 变态跳台阶
//
// 一只青蛙一次可以跳上 1 级台阶，也可以跳上 2 级... 它也可以跳上 n 级。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
//
//  n = 0 f(0)
//  n = 1 f(0) + 1
//  n = 2 f(0) + f(1) + 1
//  n = 3 f(0) + f(1) + f(2) + 1
//  n = 4 f(0) + f(1) + f(2) + f(3) + 1
//  n = 5 f(0) + f(1) + f(2) + f(3) + f(4) + 1 = 2*f(4) = 2*f(n-1)
//
// 通过推导 得出结果实际上是求2的n-1次方
func jumpFloorII(n uint) uint {
	if n == 0 {
		return 0
	} else {
		return uint(math.Pow(float64(2), float64(n-1)))
	}
}
