package module

// Game représente l'état d'une partie de Puissance 4
type Game struct {
	Grid      [][]string
	Turn      string
	Condition int
	Winner    string
	WinsX     int
	WinsO     int
}

type GameParams struct {
	Rows      int
	Cols      int
	Condition int
}

// Variables globales pour l'état du jeu
var CurrentGame Game
var GameData = GameParams{
	Rows:      6,
	Cols:      7,
	Condition: 4,
}

// InitGameCustom initialise une nouvelle partie
func InitGameCustom(rows, cols, condition int) {
	grid := make([][]string, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]string, cols)
		for j := 0; j < cols; j++ {
			grid[i][j] = "| - |"
		}
	}

	CurrentGame = Game{
		Turn:      "| X |",
		Condition: condition,
		Grid:      grid,
		WinsX:     CurrentGame.WinsX,
		WinsO:     CurrentGame.WinsO,
		Winner:    "",
	}
}

// PlayMove place un jeton dans la colonne donnée
func PlayMove(col int) {
	for row := len(CurrentGame.Grid) - 1; row >= 0; row-- {
		if CurrentGame.Grid[row][col] == "| - |" {
			CurrentGame.Grid[row][col] = CurrentGame.Turn
			if CurrentGame.Turn == "| X |" {
				CurrentGame.Turn = "| O |"
			} else {
				CurrentGame.Turn = "| X |"
			}
			return
		}
	}
}

// Check_Win_Con vérifie si un joueur a gagné
func Check_Win_Con() bool {
	rows := len(CurrentGame.Grid)
	cols := len(CurrentGame.Grid[0])
	c := CurrentGame.Condition

	// Horizontal
	for i := 0; i < rows; i++ {
		for j := 0; j <= cols-c; j++ {
			start := CurrentGame.Grid[i][j]
			if start != "| - |" {
				win := true
				for k := 1; k < c; k++ {
					if CurrentGame.Grid[i][j+k] != start {
						win = false
						break
					}
				}
				if win {
					return true
				}
			}
		}
	}

	// Vertical
	for i := 0; i <= rows-c; i++ {
		for j := 0; j < cols; j++ {
			start := CurrentGame.Grid[i][j]
			if start != "| - |" {
				win := true
				for k := 1; k < c; k++ {
					if CurrentGame.Grid[i+k][j] != start {
						win = false
						break
					}
				}
				if win {
					return true
				}
			}
		}
	}

	// Diagonal top-left -> bottom-right
	for i := 0; i <= rows-c; i++ {
		for j := 0; j <= cols-c; j++ {
			start := CurrentGame.Grid[i][j]
			if start != "| - |" {
				win := true
				for k := 1; k < c; k++ {
					if CurrentGame.Grid[i+k][j+k] != start {
						win = false
						break
					}
				}
				if win {
					return true
				}
			}
		}
	}

	// Diagonal bottom-left -> top-right
	for i := c - 1; i < rows; i++ {
		for j := 0; j <= cols-c; j++ {
			start := CurrentGame.Grid[i][j]
			if start != "| - |" {
				win := true
				for k := 1; k < c; k++ {
					if CurrentGame.Grid[i-k][j+k] != start {
						win = false
						break
					}
				}
				if win {
					return true
				}
			}
		}
	}

	return false
}

// CheckDraw vérifie si le plateau est plein et pas de gagnant
func CheckDraw() bool {
	for _, row := range CurrentGame.Grid {
		for _, cell := range row {
			if cell == "| - |" {
				return false
			}
		}
	}
	return !Check_Win_Con()
}

// IncrementWin incrémente le score d'un joueur
func IncrementWin(player string) {
	switch player {
	case "X":
		CurrentGame.WinsX++
	case "O":
		CurrentGame.WinsO++
	}
}

// GetWinCounts retourne les scores
func GetWinCounts() (int, int) {
	return CurrentGame.WinsX, CurrentGame.WinsO
}

// ResetScores remet les scores à zéro
func ResetScores() {
	CurrentGame.WinsX = 0
	CurrentGame.WinsO = 0
}
