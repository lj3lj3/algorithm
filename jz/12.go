package jz

// 矩阵中的路径
//
// 判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一个格子开始，每一步可以在矩阵中向上下左右移动一个格子。
// 如果一条路径经过了矩阵中的某一个格子，则该路径不能再进入该格子。
//
// 使用回溯法（backtracking）进行求解 其实就是搜索每个方向
// 搜索各个方向的时候 使用或运算进行递归调用 如果当前方向返回false 则自动向下一个方向进行搜索 不会出现重复搜索的路径
// 每次搜索过的路径都标注 退回的时候需要将当前的路径标注清除
func HasPath(matrix []rune, rows int, cols int, str []rune) bool {
	// construct real matrix and step history
	matrix2D := make([][]rune, rows)
	var rowIndex, colIndex int
	for _, value := range matrix {
		if colIndex == cols {
			// new row
			rowIndex++
			colIndex = 0
		}
		if colIndex == 0 {
			matrix2D[rowIndex] = make([]rune, cols)
		}

		matrix2D[rowIndex][colIndex] = value
		colIndex++
	}

	var waitCount int
	ch := make(chan bool)
	defer close(ch)

	// matrix2D is ready
	// search start point in the matrix
	for rowIndex, rowData := range matrix2D {
		for colIndex, colData := range rowData {
			if colData == str[0] {
				waitCount++
				// check for it
				go func(rowIndex, colIndex int) {
					// create temp slice to store passed path
					passed := make([][]bool, rows)
					for key := range passed {
						passed[key] = make([]bool, cols)
					}

					// get result
					ch <- hasPathGo(matrix2D, passed, rows, cols, rowIndex, colIndex, 0, str)
				}(rowIndex, colIndex)
			}
		}
	}

	result := false
	// wait for it
	for i := 0; i < waitCount; i++ {
		if <-ch {
			result = true
		}
	}

	return result
}

func hasPathGo(matrix [][]rune, passed [][]bool, rows, cols, startRow, startCol, strIndex int, str []rune) bool {
	// check if start row and start col is legal
	if startRow < 0 || startRow >= rows || startCol < 0 || startCol >= cols {
		return false
	}

	// check if value is same AND not passed
	if str[strIndex] != matrix[startRow][startCol] || passed[startRow][startCol] {
		return false
	}

	// passed
	passed[startRow][startCol] = true

	// check if search is over
	if strIndex == len(str)-1 {
		return true
	}

	// search around
	if hasPathGo(matrix, passed, rows, cols, startRow-1, startCol, strIndex+1, str) ||
		hasPathGo(matrix, passed, rows, cols, startRow, startCol+1, strIndex+1, str) ||
		hasPathGo(matrix, passed, rows, cols, startRow+1, startCol, strIndex+1, str) ||
		hasPathGo(matrix, passed, rows, cols, startRow, startCol-1, strIndex+1, str) {
		return true
	} else {
		// fall back
		passed[startRow][startCol] = false
		return false
	}
}
