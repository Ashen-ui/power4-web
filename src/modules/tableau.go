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
