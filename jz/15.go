package jz

// 输入一个整数，输出该数二进制表示中 1 的个数。
func NumberOf1(n int) int {
	count := 0
	for n > 1 {
		if n&1 == 1 {
			// the lowest bit is 1
			count++
		}

		n >>= 1 // move all bits to right
	}

	if n == 1 {
		count++
	}

	return count
}
