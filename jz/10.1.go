package jz

// 斐波那契数列
//
// 求斐波那契数列的第 n 项，n <= 39。
//         { 0                  n = 0
// 	f(n) = { 1                  n = 1
//         { f(n-1) + f(n-2)    n > 1
func Fibonacci(n uint) uint {
	if n == 0 {
		return 0
	}

	first, second := uint(0), uint(1)

	// loop
	for i := uint(2); i <= n; i++ {
		first, second = second, first+second
	}

	return second
}
