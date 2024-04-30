package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	board := CreateBoard()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		ShowBoard(board)

		fmt.Print("Enter row and column (e.g., '3D' or 'D3') or 'q' to quit: ")
		scanner.Scan()
		input := strings.ToUpper(scanner.Text())
		if input == "Q" {
			break
		}

		if len(input) != 2 {
			fmt.Println("Invalid input. Please enter two values.")
			continue
		}

		row, col := ParseInput(input)
		if row == -1 || col == -1 {
			fmt.Println("Invalid row or column.")
			continue
		}

		if !board[row][col].IsVisible {
			if board[row][col].IsMine {
				fmt.Println("Game Over! You hit a mine.")
				RevealBoard(board)
				ShowBoard(board)
				break
			}
			RevealCell(board, row, col)
		} else {
			fmt.Println("Cell already revealed. Pick another one.")
		}
	}
}
