package main

import (
	"fmt"
	"golang.org/x/term"
)

const (
	BLOCK      = "█"
	BLANK_CELL = " "
)

func main() {
	fmt.Println("hewwo")

	//TODO when changing the char in each box;create a new board to prevent the next block to be influenced by each change during the loop

	board := GenerateBoard()
	board = PopulateBoard(board)
	DisplayBoard(board)
	println()
	for i := 0; i < 100000; i++ {
		fmt.Print("\033[H\003")
		DisplayBoard(board)
		board = ChangeBlock(board)
	}
}

func PopulateBoard(b [][]bool) [][]bool {
	for i, rows := range b {
		for j := range rows {
			if i%2 == 0 {
				b[i][j] = true
			} else {
				b[i][j] = false
			}
		}
	}
	return b
}

func DisplayBoard(b [][]bool) {
	for _, rows := range b {
		for _, cols := range rows {
			//print(i, " ", j, " ")
			if cols == true {
				print(BLOCK)

			} else {
				print(BLANK_CELL)
			}
			//fmt.Printf("%s ", cols)
		}

		println()
	}
}

func GenerateBoard() [][]bool {

	//height := 50
	//width := 80

	width, height := getConsoleSize()

	var board [][]bool
	for i := 0; i < height; i++ {
		var row []bool
		for j := 0; j < width; j++ {
			row = append(row, false)
		}
		board = append(board, row)
	}

	return board
}

func ChangeBlock(b [][]bool) [][]bool {

	orws := len(b)
	ocls := len(b[0])
	newB := make([][]bool, orws)
	for i := range newB {
		newB[i] = make([]bool, ocls)
	}

	for i, rows := range b {
		for j := range rows {

			if b[i][j] == true { // alive
				switch CheckAroundAndFindOut(b, i, j) {
				case 2, 3:
					newB[i][j] = true
				default:
					newB[i][j] = false
				}
			} else {
				switch CheckAroundAndFindOut(b, i, j) {
				case 3:
					newB[i][j] = false
				default:
					newB[i][j] = false
				}
			}
		}
	}
	return newB
}

func getConsoleSize() (int, int) {
	defaultWidth := 80
	defaultHeight := 50
	width, height, err := term.GetSize(0)
	if err != nil {
		width = defaultWidth
		height = defaultHeight
	}
	return width, height

}

func CheckAroundAndFindOut(grid [][]bool, row, col int) int {
	Count := 0
	rows := len(grid)
	cols := len(grid[0])
	indices := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for _, i := range indices {
		r := row + i[0]
		c := col + i[1]
		if r >= 0 && r < rows && c >= 0 && c < cols {
			if grid[r][c] {
				Count++
			}
		}
	}
	return Count
}
