package other

import "math"

// 快速求平方根倒数
//
// https://en.wikipedia.org/wiki/Fast_inverse_square_root
func reverseSquareRootFast(number int) float64 {
	num := float64(number)
	val := num

	// Bits hacking
	numBits := math.Float64bits(val)
	// 0x 5FE6 EB50 C7B5 37A9
	numBits = 0x5FE6EB50C7B537A9 - (numBits >> 1)
	// Set value
	val = math.Float64frombits(numBits)

	// Newton way
	val = val * (1.5 - (0.5 * num * val * val))
	val = val * (1.5 - (0.5 * num * val * val))

	return val
}
