package jz

// 顺时针打印矩阵
//
// 输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字，
// 例如，如果输入如下4 X 4矩阵： 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16
// 则依次打印出数字1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10
//
// 可以使用单独的数组来记录是否已经有打印当前位置 这里使用了总数判断 如果数组中的总数等于二维数组内的元素数 则循环完成
func PrintMatrix(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	// Get length of matrix
	rows := len(matrix)
	cols := len(matrix[0])
	count := rows * cols
	results := make([]int, 0, count)

	// Loop
	row := 0
	col := 0
	layer := 0
	for ; ; layer++ {
		// Move right
		for ; col < cols-layer; col++ {
			results = append(results, matrix[row][col])
		}
		if len(results) == count {
			break
		}
		col--
		row++
		// Move down
		for ; row < rows-layer; row++ {
			results = append(results, matrix[row][col])
		}
		if len(results) == count {
			break
		}
		col--
		row--
		// Move left
		for ; col >= layer; col-- {
			results = append(results, matrix[row][col])
		}
		if len(results) == count {
			break
		}
		row--
		col++
		// Move up
		for ; row >= layer+1; row-- {
			results = append(results, matrix[row][col])
		}
		if len(results) == count {
			break
		}
		col++
		row++
	}

	return results
}
