package jz

// 数值的整数次方
//
// 给定一个 double 类型的浮点数 base 和 int 类型的整数 exponent，求 base 的 exponent 次方。
//
// 17次方等于16次方乘以1次方 16次方等于8次方的平方 8次方等于4次方的平方 4次方等于2次方的平方 2次方是1次方的平方
//
// 负次方等于次方的倒数
func Power(base float64, exponent int) float64 {
	if exponent == 0 {
		return 1
	} else if exponent == 1 {
		return base
	}

	isNegative := false
	if exponent < 0 {
		exponent = -exponent
		isNegative = true
	}

	result := Power(base, exponent/2)
	result = result * result

	if exponent%2 == 1 {
		// need times base
		result *= base
	}

	if isNegative {
		return 1 / result
	} else {
		return result
	}
}
