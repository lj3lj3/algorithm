package jz

// 机器人的运动范围
//
// 地上有一个 m 行和 n 列的方格。一个机器人从坐标 (0, 0) 的格子开始移动，每一次只能向左右上下四个方向移动一格，但是不能进入行坐标和列坐标的数位之和大于 k 的格子。
// 例如，当 k 为 18 时，机器人能够进入方格 (35,37)，因为 3+5+3+7=18。但是，它不能进入方格 (35,38)，因为 3+5+3+8=19。请问该机器人能够达到多少个格子？
//
// 同样使用回溯法 递归搜索整个二维数组
var count = 0

func MovingCount(threshold, rows, cols int) int {
	// reset count
	count = 0

	type thresholdMsg struct {
		row     int
		results []bool
	}

	// calc the threshold is valid, ans store in the 2d array
	ch := make(chan *thresholdMsg)
	waitCount := 0

	thresholdArray := make([][]bool, rows, rows)
	passed := make([][]bool, rows, rows)
	for row := range thresholdArray {
		passed[row] = make([]bool, cols, cols) // init as false

		waitCount++
		go func(row int) {
			thresholdArray := make([]bool, cols, cols)
			for col := range thresholdArray {
				// calc the threshold
				thresholdArray[col] = calcThresholdResult(threshold, row, col)
			}

			ch <- &thresholdMsg{
				row:     row,
				results: thresholdArray,
			}
		}(row)
	}

	for i := 0; i < waitCount; i++ {
		msg := <-ch
		thresholdArray[msg.row] = msg.results
	}

	// let's go
	going(passed, thresholdArray, 0, 0)

	return count
}

func going(passed, thresholdArray [][]bool, row, col int) {
	// check if row and col are valid
	if row < 0 || col < 0 || row >= len(thresholdArray) || col >= len(thresholdArray[0]) ||
		thresholdArray[row][col] || passed[row][col] { // threshold or passed
		// invalid
		return
	}

	passed[row][col] = true
	count++

	// go to next direction
	going(passed, thresholdArray, row, col-1)
	going(passed, thresholdArray, row+1, col)
	going(passed, thresholdArray, row, col+1)
	going(passed, thresholdArray, row-1, col)
}

func calcThresholdResult(threshold, row, col int) bool {
	count := 0
	// calc the count of row
	for row > 0 {
		count += row % 10 // the value of the lowest number
		row /= 10         // ignore the lowest number
	}
	// calc the count of col
	for col > 0 {
		count += col % 10
		col /= 10
	}

	return count > threshold
}
