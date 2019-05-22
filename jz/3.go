package jz

// 3. 数组中重复的数字
//
// 在一个长度为 n 的数组里的所有数字都在 0 到 n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字是重复的，也不知道每个数字重复几次。请找出数组中任意一个重复的数字。
//
// Input:
// {2, 3, 1, 0, 2, 5}
//
// Output:
// 2
//
// 将number的值放到相应的位置上求解
func duplicate(nums []int) int {
	for i := range nums {
		for {
			if i == nums[i] {
				break
			} else if nums[i] == nums[nums[i]] {
				// found
				return nums[i]
			} else {
				// swap by using index instead of number
				nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
			}
		}
	}
	panic("not found")
}
