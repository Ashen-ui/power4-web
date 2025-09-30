package module

type Game struct {
	Grid [6][7]string // 6 lignes, 7 colonnes
	Turn string       // "X" ou "O"
}

var CurrentGame Game

func InitGame() {
	CurrentGame = Game{}
	CurrentGame.Turn = "X"
}

func PlayMove(col int) {
	for row := len(CurrentGame.Grid) - 1; row >= 0; row-- {
		if CurrentGame.Grid[row][col] == "" {
			CurrentGame.Grid[row][col] = CurrentGame.Turn
			if CurrentGame.Turn == "X" {
				CurrentGame.Turn = "O"
			} else {
				CurrentGame.Turn = "X"
			}
			return
		}
	}
}

func GetGame() Game {
	return CurrentGame
}
