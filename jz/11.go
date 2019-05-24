package jz

// 旋转数组的最小数字
//
// 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个非递减排序的数组的一个旋转，输出旋转数组的最小元素。
//  1, 1, 2, 2, 3, 4 -> 2, 2, 3, 4,  1, 1
//  1, 1, 2, 2, 3, 4 -> 3, 4,  1, 1, 2, 2
//  1, 1, 1, 1, 2, 2 -> 2, 2,  1, 1, 1, 1
//  1, 1, 1, 1, 2, 2 -> 1, 2, 2,  1, 1, 1,
//
// 数组二分后 根据数组的第一个元素和中间元素比较 确定旋转的数组在哪一部分 如果相等则无法确定 因为不是排序的数组 所以这里可以使用顺序查找法
func minNumberInRotatedArray(rotatedArray []int) int {
	if len(rotatedArray) == 0 {
		return 0
	}

	// first split array in middle
	mid := len(rotatedArray) / 2

	if rotatedArray[0] > rotatedArray[mid] {
		// we're in the rotated array
		// moving left from middle
		for i := mid; i > 0; i-- {
			if rotatedArray[i-1] > rotatedArray[i] {
				// this is mini item
				return rotatedArray[i]
			}
		}
		// not found, return first value
		return rotatedArray[0]
	} else if rotatedArray[0] < rotatedArray[mid] {
		// we're not in the rotated array
		// moving right from middle
		for i := mid; i < len(rotatedArray)-1; i++ {
			if rotatedArray[i] > rotatedArray[i+1] {
				// this is mini item
				return rotatedArray[i+1]
			}
		}
		// right sub array is increase array
		return rotatedArray[0] // the first is min
	} else {
		var min int
		// don't know the direction to search for, loop it
		for i := 0; i < len(rotatedArray); i++ {
			if rotatedArray[i] < min {
				min = rotatedArray[i]
			}
		}
		return min
	}
}
