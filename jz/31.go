package jz

// 栈的压入、弹出序列
//
// 输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否可能为该栈的弹出顺序。假设压入栈的所有数字均不相等。
// 例如序列1,2,3,4,5是某栈的压入顺序，序列4,5,3,2,1是该压栈序列对应的一个弹出序列，但4,3,5,1,2就不可能是该压栈序列的弹出序列。（注意：这两个序列的长度是相等的）
//
// 使用一个stack来模拟入栈出栈操作
func IsPopOrder(pushOrder []int, popOrder []int) bool {
	nextPushIndex := 0
	popIndex := 0
	pushLength := len(pushOrder)
	popLength := len(popOrder)
	stack := make([]int, 0, pushLength)

	// Does not finished yet
	for nextPushIndex <= pushLength && popIndex <= popLength-1 {
		if len(stack) == 0 {
			// Push
			stack = append(stack, pushOrder[nextPushIndex])
			nextPushIndex++
		} else if nextPushIndex == pushLength || pushOrder[nextPushIndex-1] == popOrder[popIndex] {
			// No items to push
			// This is a pop order
			stackLen := len(stack)
			if stack[stackLen-1] != popOrder[popIndex] {
				return false
			}
			// Pop it
			stack = stack[:len(stack)-1]
			popIndex++
		} else {
			// We should just push it
			stack = append(stack, pushOrder[nextPushIndex])
			nextPushIndex++
		}
	}

	return len(stack) == 0
}
