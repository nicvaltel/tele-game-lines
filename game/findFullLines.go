package game

import "fmt"

func (game *GameState) ProcessFullLines() {
	consecutiveLines := findConsecutiveLines(game.field, FULL_ROW)
	for r, cols := range consecutiveLines {
		for c, val := range cols {
			if val != 0 {
				game.field[r][c] = 0
			}
		}
	}
}

func findConsecutiveLines(matrix [][]uint8, consecutiveLen int) [][]int {

	result := make([][]int, FIELD_SIZE)
	for i := 0; i < FIELD_SIZE; i++ {
		result[i] = make([]int, FIELD_SIZE)
	}

	// Find horizontal consecutive rows
	for row := 0; row < FIELD_SIZE; row++ {
		consecutiveCount := 1
		colStart := 0
		for col := 1; col < FIELD_SIZE; col++ {
			if matrix[row][col] != 0 && matrix[row][col] == matrix[row][col-1] {
				consecutiveCount++
			} else {
				if consecutiveCount >= consecutiveLen {
					for cc := colStart; cc < col; cc++ {
						result[row][cc] = 1
					}
				}
				consecutiveCount = 1
				colStart = col
			}
		}
		if consecutiveCount >= consecutiveLen {
			for cc := colStart; cc < FIELD_SIZE; cc++ {
				result[row][cc] = 1
			}
		}
	}

	// Find vertical consecutive rows
	for col := 0; col < FIELD_SIZE; col++ {
		consecutiveCount := 1
		rowStart := 0
		for row := 1; row < FIELD_SIZE; row++ {
			if matrix[row][col] != 0 && matrix[row][col] == matrix[row-1][col] {
				consecutiveCount++
			} else {
				if consecutiveCount >= consecutiveLen {
					for rr := rowStart; rr < row; rr++ {
						result[rr][col] = 1
					}
				}
				consecutiveCount = 1
				rowStart = row
			}
		}
		if consecutiveCount >= consecutiveLen {
			for rr := rowStart; rr < FIELD_SIZE; rr++ {
				result[rr][col] = 1
			}
		}
	}

	// Find diagonal (top-left to bottom-right) consecutive cells
	for startRow := 0; startRow < FIELD_SIZE; startRow++ {
		for startCol := 0; startCol < FIELD_SIZE; startCol++ {
			consecutiveCount := 1
			row, col := startRow+1, startCol+1
			for row < FIELD_SIZE && col < FIELD_SIZE {
				if matrix[row][col] != 0 && matrix[row][col] == matrix[row-1][col-1] {
					consecutiveCount++
				} else {
					break
				}
				if consecutiveCount >= consecutiveLen {
					for i := 0; i < consecutiveLen; i++ {
						result[startRow+i][startCol+i] = 1
					}
				}
				row++
				col++
			}
		}
	}

	// Find diagonal (bottom-left to top-right) consecutive cells
	for startRow := FIELD_SIZE - 1; startRow >= 0; startRow-- {
		for startCol := 0; startCol < FIELD_SIZE; startCol++ {
			consecutiveCount := 1
			row, col := startRow-1, startCol+1
			for row >= 0 && col < FIELD_SIZE {
				if matrix[row][col] != 0 && matrix[row][col] == matrix[row+1][col-1] {
					consecutiveCount++
				} else {
					break
				}
				if consecutiveCount >= consecutiveLen {
					for i := 0; i < consecutiveLen; i++ {
						result[startRow-i][startCol+i] = 1
					}
				}
				row--
				col++
			}
		}
	}

	return result
}

func TestFullLines() {
	matrix := [][]uint8{
		{0, 0, 1, 0, 0, 4, 0, 1, 0},
		{5, 0, 1, 0, 4, 8, 0, 1, 0},
		{0, 5, 1, 4, 0, 8, 0, 1, 0},
		{0, 1, 4, 0, 1, 8, 0, 1, 0},
		{0, 4, 0, 5, 1, 0, 0, 1, 0},
		{1, 1, 1, 0, 5, 8, 0, 1, 0},
		{0, 0, 0, 1, 0, 5, 0, 1, 0},
		{0, 0, 0, 0, 0, 8, 5, 1, 0},
		{0, 8, 8, 8, 8, 8, 0, 5, 0},
	}

	consecutiveLen := 5
	consecutiveLines := findConsecutiveLines(matrix, consecutiveLen)

	for _, cols := range consecutiveLines {
		for _, val := range cols {
			fmt.Printf("%4d", val)
		}
		fmt.Println()
	}
}
