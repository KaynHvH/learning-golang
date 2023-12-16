package main

import "fmt"

var board [3][3]uint

func drawBoard() {
	for i := 0; i < len(board[0]); i++ {
		for j := 0; j < len(board[1]); j++ {
			switch board[i][j] {
			case 0:
				fmt.Print("   ")
			case 1:
				fmt.Print(" X ")
			case 2:
				fmt.Print(" O ")
			}

			if j < 2 {
				fmt.Print("|")
			}
		}

		fmt.Println()
		if i < 2 {
			fmt.Println("-----------")
		}
	}
}

func displayBoard() {
	fmt.Println("Tic Tac Toe Board:")
	drawBoard()
}

func makeMove(row, col, player uint) bool {
	if row < 1 || row > 3 || col < 1 || col > 3 || board[row-1][col-1] != 0 {
		fmt.Println("Wrong move. Try again!")
		return false
	}

	board[row-1][col-1] = player
	return true
}

func checkWin(player uint) bool {
	for i := 0; i < 3; i++ {
		if (board[i][0] == player && board[i][1] == player && board[i][2] == player) ||
			(board[0][i] == player && board[1][i] == player && board[2][i] == player) {
			return true
		}
	}

	if (board[0][0] == player && board[1][1] == player && board[2][2] == player) ||
		(board[0][2] == player && board[1][1] == player && board[2][0] == player) {
		return true
	}

	return false
}

func checkDraw() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	displayBoard()

	for {
		var row, col uint
		fmt.Print("Player X, enter a row (1-3): ")
		fmt.Scanln(&row)
		fmt.Print("Player X, enter a col (1-3): ")
		fmt.Scanln(&col)

		if makeMove(row, col, 1) {
			displayBoard()
		} else {
			continue
		}

		if checkWin(1) {
			fmt.Println("Player X won!")
			break
		} else if checkDraw() {
			fmt.Println("Draw!")
			break
		}

		fmt.Print("Player O, enter a row (1-3): ")
		fmt.Scanln(&row)
		fmt.Print("Player O, enter a col (1-3): ")
		fmt.Scanln(&col)

		if makeMove(row, col, 2) {
			displayBoard()
		} else {
			continue
		}

		if checkWin(2) {
			fmt.Println("Player O won!")
			break
		} else if checkDraw() {
			fmt.Println("Draw!")
			break
		}
	}
}
