package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	width    = 10
	height   = 10
	numMines = 15
)

type Cell struct {
	IsMine    bool
	IsVisible bool
	NumNearby int
}

type Board [][]Cell

func CreateBoard() Board {
	time.Now().UnixNano()
	board := make(Board, height)
	for i := range board {
		board[i] = make([]Cell, width)
	}

	for i := 0; i < numMines; i++ {
		for {
			row := rand.Intn(height)
			col := rand.Intn(width)
			if !board[row][col].IsMine {
				board[row][col].IsMine = true
				break
			}
		}
	}

	for row := range board {
		for col := range board[row] {
			if !board[row][col].IsMine {
				count := 0
				for _, dr := range []int{-1, 0, 1} {
					for _, dc := range []int{-1, 0, 1} {
						if !(dr == 0 && dc == 0) && IsValidCell(board, row+dr, col+dc) && board[row+dr][col+dc].IsMine {
							count++
						}
					}
				}
				board[row][col].NumNearby = count
			}
		}
	}

	return board
}

func IsValidCell(board Board, row, col int) bool {
	return row >= 0 && row < len(board) && col >= 0 && col < len(board[0])
}

func RevealCell(board Board, row, col int) {
	if !IsValidCell(board, row, col) || board[row][col].IsVisible {
		return
	}

	board[row][col].IsVisible = true
	if board[row][col].NumNearby == 0 {
		for _, dr := range []int{-1, 0, 1} {
			for _, dc := range []int{-1, 0, 1} {
				RevealCell(board, row+dr, col+dc)
			}
		}
	}
}

func RevealBoard(board Board) {
	for row := range board {
		for col := range board[row] {
			board[row][col].IsVisible = true
		}
	}
}

func ShowBoard(board Board) {
	fmt.Print("  ")
	for c := 0; c < width; c++ {
		fmt.Printf("%c ", 'A'+c)
	}
	fmt.Println()

	for r, row := range board {
		fmt.Printf("%d ", r)
		for _, cell := range row {
			if cell.IsVisible {
				if cell.IsMine {
					fmt.Print("X ")
				} else {
					fmt.Print(cell.NumNearby, " ")
				}
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func ParseInput(input string) (int, int) {
	row, _ := strconv.Atoi(string(input[1]))
	col := int(input[0] - 'A')
	return row, col
}
