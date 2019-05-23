package jz

// 矩形覆盖
//
// 我们可以用 2*1 的小矩形横着或者竖着去覆盖更大的矩形。请问用 n 个 2*1 的小矩形无重叠地覆盖一个 2*n 的大矩形，总共有多少种方法？
//
// 从n为1起开始计算 当n等于1时 值为1 当n等于2时 值为2 当n等于3时 值为3 当n等于4时 值为5 本质为Fibonacci数列
func rectCover(n uint) uint {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	first, second := uint(1), uint(2)

	for i := uint(3); i <= n; i++ {
		first, second = second, first+second
	}

	return second
}
