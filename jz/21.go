package jz

// 调整数组顺序使奇数位于偶数前面
//
// 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有的奇数位于数组的前半部分，所有的偶数位于数组的后半部分，并保证奇数和奇数，偶数和偶数之间的相对位置不变。
//
// 使用一个单独的数组来保存临时副本 然后计算出even的数量 循环数组中的每个值 分别按类型放置到相应的位置
func ReOrderArray(arr []int) {
	// clone a new array, by counting the odd count
	src := make([]int, len(arr))
	copy(src, arr)

	// ignore data in the arr
	// calc the count of odd
	oddCount := 0
	for _, value := range arr {
		oddCount += value & 1
	}

	// split the arr into two parts, but store in one array
	oddPos := 0
	evenPos := oddCount
	for _, value := range src {
		if value&1 == 1 {
			// odd
			arr[oddPos] = value
			oddPos++
		} else {
			// even
			arr[evenPos] = value
			evenPos++
		}
	}
}

// 使用冒泡来进行排序
func ReOrderArrayII(arr []int) {
	moveCount := 1
	arrLen := len(arr)
	for moveCount > 0 {
		// reset move count
		moveCount = 0
		for key, value := range arr {
			// this value is even and next value is odd
			if value&1 == 0 && key+1 < arrLen && arr[key+1]&1 == 1 {
				// swap
				arr[key], arr[key+1] = arr[key+1], value
				// increase count of movement
				moveCount++
			}
		}
	}
}
