package jz

// 4. 二维数组中的查找
//
// 给定一个二维数组，其每一行从左到右递增排序，从上到下也是递增排序。给定一个数，判断这个数是否在该二维数组中。
//
// Consider the following matrix:
//[
//  [1,   4,  7, 11, 15],
//  [2,   5,  8, 12, 19],
//  [3,   6,  9, 16, 22],
//  [10, 13, 14, 17, 24],
//  [18, 21, 23, 26, 30]
//]
//
//Given target = 5, return true.
//Given target = 20, return false.
//
// 从右上角开始 左边的比它小 下边的比它大
func find(nums [][]int, target int) bool {
	// find target in the numbers
	x, y := len(nums[0])-1, 0
	// check if it is less than first item
	if target < nums[0][0] {
		return false
	}
	// check if it is bigger than last item
	if target > nums[len(nums)-1][x] {
		return false
	}
	// in the square
	lastRow := len(nums) - 1
	for {
		// check this point
		if x < 0 || y > lastRow {
			// not found
			return false
		} else if target == nums[x][y] {
			return true // found
		} else if target > nums[x][y] {
			// move down
			y++
		} else if target < nums[x][y] {
			// move left
			x--
		}
	}
}
