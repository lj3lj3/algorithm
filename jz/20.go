package jz

const (
	SIGNAL = iota
	DIGIT
	DOT
	E
	SPACE
	INVALID
	END
)

// 表示数值的字符串
//
// true
//  "+100"
//  "5e2"
//  "-123"
//  "3.1416"
//  "-1E-16"
//
// false
//  "12e"
//  "1a3.14"
//  "1.2.3"
//  "+-5"
//  "12e+4.3"
//
// 使用正则匹配则可以使用确定有限自动机
//
// 这里使用二维数组来表示状态之间的转换 二维数组的一维表示各个状态 从初始到最后 所有可能的状态 状态机启动默认直接进入初始状态 有多少个可能的状态 一维数组就有多少个
//
// 二维数组的二维数组表示这个状态的各种输入值所对应的下一个状态 有多少种输入值 二维数组就有多少个
//
// 状态机的过程不一定只有一个 像其中的dot之前有数字和没有数字就是分叉过程
func IsNumeric(str []rune) bool {
	// init transform table
	table := [][]int{
		// 竖向表示当前的输入值
		//+/- n	.	e 	SPACE INVALID END
		{1, 2, 9, -1, 0, -1, -1},    // 初始	横向表示状态及每一个可能的位置 初始状态直接跳转到相应的状态即可	如遇到空格则直接循环到当前状态
		{-1, 2, 9, -1, -1, -1, -1},  // 1 +/-
		{-1, 2, 3, 5, 8, -1, 10},    // 2 n
		{-1, 4, -1, 5, 8, -1, 10},   // 3 . 之前有数字
		{-1, 4, -1, 5, 8, -1, 10},   // 4 n
		{6, 7, -1, -1, -1, -1, -1},  // 5 e
		{-1, 7, -1, -1, -1, -1, -1}, // 6 +/-
		{-1, 7, -1, -1, 8, -1, 10},  // 7 n
		{-1, -1, -1, -1, 8, -1, 10}, // 8 space	最后的空格
		{-1, 4, -1, -1, -1, -1, -1}, // 9 . 之前没有数字
	}

	str = append(str, '\n') // append the end

	state := 0
	input := -1
	// run through the table
	for _, char := range str {
		switch char {
		case ' ':
			input = SPACE
		case '+':
			fallthrough
		case '-':
			input = SIGNAL
		case '.':
			input = DOT
		case 'e':
			fallthrough
		case 'E':
			input = E
		case '\n':
			input = END
		default:
			if char >= '0' && char <= '9' {
				input = DIGIT
			} else {
				input = INVALID
			}
		}

		// run
		state = table[state][input]
		if state == -1 {
			return false
		}
	}

	return true
}
