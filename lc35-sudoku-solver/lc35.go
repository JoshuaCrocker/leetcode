package lc35sudokusolver

import (
	"log"
	"math"
)

const Empty byte = '.'

// Write a program to solve a Sudoku puzzle by filling the empty cells.
//
// A sudoku solution must satisfy all of the following rules:
//
//     Each of the digits 1-9 must occur exactly once in each row.
//     Each of the digits 1-9 must occur exactly once in each column.
//     Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
//
// The '.' character indicates empty cells.
func SolveSudoku(board [][]byte) {
	iterations := 10
	for NumberOfEmptyCells(board) > 0 && iterations > 0 {
		log.Printf("empty=%d", NumberOfEmptyCells(board))
		row := 0
		for row < len(board) {
			col := 0
			for col < len(board[row]) {
				if board[row][col] == Empty {
					value := solveCell(board, row, col)
					if value != Empty {
						board[row][col] = value
					}
				}
				col++
			}
			row++
		}
		iterations--
	}
}

func NumberOfEmptyCells(board [][]byte) int {
	count := 0

	row := 0
	for row < len(board) {
		col := 0
		for col < len(board[row]) {
			if board[row][col] == Empty {
				count++
			}
			col++
		}
		row++
	}

	return count
}

func NumbersInRow(board [][]byte, row int) []byte {
	output := []byte{}
	for _, value := range board[row] {
		if value != Empty {
			output = append(output, value)
		}
	}

	return output
}

func NumbersInCol(board [][]byte, col int) []byte {
	output := []byte{}
	row := 0
	for row < len(board) {
		value := board[row][col]
		if value != Empty {
			output = append(output, value)
		}
		row++
	}

	return output
}

func NumbersInCell(board [][]byte, row, col int) []byte {
	output := []byte{}
	startRow := int(math.Floor(float64(row)/3)) * 3
	startCol := int(math.Floor(float64(col)/3)) * 3

	ptrRow := startRow

	endRow := ptrRow + 3

	for ptrRow < endRow {
		ptrCol := startCol
		endCol := ptrCol + 3
		for ptrCol < endCol {
			value := board[ptrRow][ptrCol]
			if value != Empty {
				output = append(output, value)
			}
			ptrCol++
		}
		ptrRow++
	}

	return output
}

func contains(needle byte, haystack []byte) bool {
	for _, value := range haystack {
		if needle == value {
			return true
		}
	}

	return false
}

func solveCell(board [][]byte, row, col int) byte {
	valuesInRow := NumbersInRow(board, row)
	valuesInCol := NumbersInRow(board, col)
	valuesInCell := NumbersInCell(board, row, col)

	allOptions := []byte("123456789")
	options := []byte{}

	for _, option := range allOptions {
		if !contains(option, valuesInRow) && !contains(option, valuesInCol) && !contains(option, valuesInCell) {
			options = append(options, option)
		}
	}

	if len(options) == 1 {
		return options[0]
	}

	return Empty
}
