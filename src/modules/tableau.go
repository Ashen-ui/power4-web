package module

func InitPlateau() [][]string {
	var board [][]string
	board = make([][]string, GameData.Rows)
	for i := 0; i < GameData.Rows; i++ {
		board[i] = make([]string, GameData.Cols)
	}
	for i := 0; i < GameData.Rows; i++ {
		for j := 0; j < GameData.Cols; j++ {
			board[i][j] = "| - |"
		}
	}
	return board
}

func IsFull(board [][]string) bool {
	for j := 0; j < GameData.Cols; j++ {
		if board[0][j] == "| - |" {
			return false
		}
	}
	return true
}

func IsColFull(board [][]string, col int) bool {
	return board[0][col] != "| - |"
}

func GetSymbol(board [][]string, row, col int) string {
	return board[row][col]
}
