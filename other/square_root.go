package other

import "math"

// 牛顿法开平方
//
// 1. 随机取一个数作为基准值val 这里使用了number * 0.5
// 2. 每次求val与number/val的平均值 之后赋值给val
// 3. 当当前的结果和上次的结果直接的值小于某个数值的时候 返回
func squareRootNewton(number int) float64 {
	// Convert to float64
	num := float64(number)

	// The val number
	val := num * 0.5
	last := 0.0

	// Loop
	for math.Abs(val-last) > 1e-6 {
		last = val // Save last value
		val = (val + num/val) * 0.5
	}

	return val
}
