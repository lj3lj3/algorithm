package jz

// 跳台阶
//
// 一只青蛙一次可以跳上 1 级台阶，也可以跳上 2 级。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
//
//  n = 0, 0
//  n = 1, 1
//  n = 2, 2
//  n = 3, 3
//  n = 4, 5
func JumpFloor(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	first, second := uint(1), uint(2)

	for i := uint(3); i <= n; i++ {
		first, second = second, first+second
	}

	return second
}
