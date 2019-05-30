package leetcode

import "algorithm/jz"

// Word Search
func Exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])

	// combine into one dimension
	matrix := make([]rune, 0, rows*cols)
	for _, boardRow := range board {
		for _, value := range boardRow {
			matrix = append(matrix, rune(value))
		}
	}

	return jz.HasPath(matrix, rows, cols, []rune(word))
}
