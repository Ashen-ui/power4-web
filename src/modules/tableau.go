package module

func InitPlateau() [6][7]string {
	var board [6][7]string
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			board[i][j] = "| - |"
		}
	}
	return board
}

func DisplayBoard(board [6][7]string) string {
	result := ""
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			result += board[i][j] + " "
		}
		result += "\n"
	}
	return result
}

func IsFull(board [6][7]string) bool {
	for j := 0; j < 7; j++ {
		if board[0][j] == "| - |" {
			return false
		}
	}
	return true
}

func IsColFull(board [6][7]string, col int) bool {
	return board[0][col] != "| - |"
}

func GetSymbol(board [6][7]string, row, col int) string {
	return board[row][col]
}

// CheckWin vérifie si un joueur a aligné 4 jetons.
// Retourne ("X", true) si X a gagné, ("O", true) si O a gagné, sinon ("", false).
func CheckWin(board [6][7]string) (string, bool) {
	for r := 0; r < 6; r++ {
		for c := 0; c < 7; c++ {
			cell := board[r][c]
			var p string
			switch cell {
			case "| X |":
				p = "X"
			case "| O |":
				p = "O"
			default:
				continue
			}

			// horizontal (to the right)
			if c+3 < 7 {
				if board[r][c+1] == cell && board[r][c+2] == cell && board[r][c+3] == cell {
					return p, true
				}
			}
			// vertical (downwards)
			if r+3 < 6 {
				if board[r+1][c] == cell && board[r+2][c] == cell && board[r+3][c] == cell {
					return p, true
				}
			}
			// diagonal down-right
			if r+3 < 6 && c+3 < 7 {
				if board[r+1][c+1] == cell && board[r+2][c+2] == cell && board[r+3][c+3] == cell {
					return p, true
				}
			}
			// diagonal down-left
			if r+3 < 6 && c-3 >= 0 {
				if board[r+1][c-1] == cell && board[r+2][c-2] == cell && board[r+3][c-3] == cell {
					return p, true
				}
			}
		}
	}
	return "", false
}
