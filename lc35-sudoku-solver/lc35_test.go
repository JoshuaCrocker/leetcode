package lc35sudokusolver_test

import (
	"reflect"
	"testing"

	lc35 "github.com/JoshuaCrocker/leetcode/lc35-sudoku-solver"
)

const Reset = "\033[0m"
const Red = "\033[31m"

func drawGrids(left, right [][]byte) string {
	output := "\n"

	x := 0
	for x < len(left) {
		y := 0
		for y < len(left[x]) {
			output += "|"
			output += string(left[x][y])
			y++
		}

		output += "|     "

		y = 0
		for y < len(right[x]) {
			output += "|"

			// if left[x][y] != right[x][y] {
			// 	output += Red
			// }

			output += string(right[x][y])

			// if left[x][y] != right[x][y] {
			// 	output += Reset
			// }
			y++
		}

		output += "|\n"

		x++
	}

	return output
}

func TestSolveSudokuGivenExample(t *testing.T) {
	input := [][]byte{[]byte("53..7...."), []byte("6..195..."), []byte(".98....6."), []byte("8...6...3"), []byte("4..8.3..1"), []byte("7...2...6"), []byte(".6....28."), []byte("...419..5"), []byte("....8..79")}
	expected := [][]byte{[]byte("534678912"), []byte("672195348"), []byte("198342567"), []byte("859761423"), []byte("426853791"), []byte("713924856"), []byte("961537284"), []byte("287419635"), []byte("345286179")}
	lc35.SolveSudoku(input)

	if !reflect.DeepEqual(input, expected) {
		t.Log(drawGrids(input, expected))

		t.Errorf("Incorrect")
	}
}

func TestNumbersInRow(t *testing.T) {
	board := [][]byte{[]byte("53..7...."), []byte("6..195..."), []byte(".98....6."), []byte("8...6...3"), []byte("4..8.3..1"), []byte("7...2...6"), []byte(".6....28."), []byte("...419..5"), []byte("....8..79")}
	row := 0
	expected := []byte("537")

	result := lc35.NumbersInRow(board, row)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestNumbersInCol(t *testing.T) {
	board := [][]byte{[]byte("53..7...."), []byte("6..195..."), []byte(".98....6."), []byte("8...6...3"), []byte("4..8.3..1"), []byte("7...2...6"), []byte(".6....28."), []byte("...419..5"), []byte("....8..79")}
	row := 3
	expected := []byte("184")

	result := lc35.NumbersInCol(board, row)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestNumbersInCell(t *testing.T) {
	board := [][]byte{[]byte("53..7...."), []byte("6..195..."), []byte(".98....6."), []byte("8...6...3"), []byte("4..8.3..1"), []byte("7...2...6"), []byte(".6....28."), []byte("...419..5"), []byte("....8..79")}
	row := 8
	col := 3
	expected := []byte("4198")

	result := lc35.NumbersInCell(board, row, col)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestNumberOfEmptyCells(t *testing.T) {
	board := [][]byte{[]byte("53..7...."), []byte("6..195..."), []byte(".98....6."), []byte("8...6...3"), []byte("4..8.3..1"), []byte("7...2...6"), []byte(".6....28."), []byte("...419..5"), []byte("....8..79")}
	expected := 51

	result := lc35.NumberOfEmptyCells(board)

	if expected != result {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
