package main

import "fmt"

// SolveSudoku takes a slice of 81 ints representing a 9x9 Sudoku board
// in row-major order. Empty cells are 0.
// The board is guaranteed to be valid and solvable.
func SolveSudoku(grid []int) []int {
	solveBacktracking(grid)
	return grid
}

func solveBacktracking(grid []int) bool {
	for i := 0; i < 81; i++ {
		if grid[i] == 0 {
			row := i / 9
			col := i % 9

			for num := 1; num <= 9; num++ { // check every num possible into that empty cell
				if canPlace(grid, row, col, num) {
					grid[i] = num // grid is modified, that position is not empty anymore

					if solveBacktracking(grid) { // start filling the next empty position, enven though i start from 0 again
						return true
					} else {
						grid[i] = 0
					}

				}
			}

			return false // if every number has been tried, then we hit the wall, return false to set ealier positition to 0 and find another path.
		}
	}

	// no empty cells left — solved
	return true
}

func canPlace(grid []int, row, col, val int) bool {
	// check row
	for c := 0; c < 9; c++ {
		if grid[row*9+c] == val {
			return false
		}
	}

	// check column
	for r := 0; r < 9; r++ {
		if grid[r*9+col] == val {
			return false
		}
	}

	// check 3x3 box
	boxRow := (row / 3) * 3
	boxCol := (col / 3) * 3

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if grid[(boxRow+r)*9+(boxCol+c)] == val {
				return false
			}
		}
	}

	return true
}

func main() {
	puzzle := []int{
		5, 3, 0, 0, 7, 0, 0, 0, 0,
		6, 0, 0, 1, 9, 5, 0, 0, 0,
		0, 9, 8, 0, 0, 0, 0, 6, 0,
		8, 0, 0, 0, 6, 0, 0, 0, 3,
		4, 0, 0, 8, 0, 3, 0, 0, 1,
		7, 0, 0, 0, 2, 0, 0, 0, 6,
		0, 6, 0, 0, 0, 0, 2, 8, 0,
		0, 0, 0, 4, 1, 9, 0, 0, 5,
		0, 0, 0, 0, 8, 0, 0, 7, 9,
	}

	solution := SolveSudoku(puzzle)

	fmt.Println(solution)
}
