package module

// Game représente l'état d'une partie de Puissance 4
type Game struct {
	Grid      [][]string
	Turn      string
	Condition int
	Joueur1   Joueur
	Joueur2   Joueur
}

var GameData = struct {
	Rows      int
	Cols      int
	Condition int
}{
	Rows:      6,
	Cols:      7,
	Condition: 4,
}

var CurrentGame Game
var WinsX int
var WinsO int

func InitGameCustom(rows, cols, condition int) {
	var grid = make([][]string, rows)
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
	}

	// Copie manuelle de la grille dynamique dans la grille fixe 6x7 si besoin
	for i := 0; i < rows && i < len(CurrentGame.Grid); i++ {
		for j := 0; j < cols && j < len(CurrentGame.Grid[i]); j++ {
			CurrentGame.Grid[i][j] = grid[i][j]
		}
	}
}

// PlayMove place un jeton dans la colonne spécifiée
func PlayMove(col int) {
	// Parcourt la grille de bas en haut (dernière ligne vers première ligne)
	for row := len(CurrentGame.Grid) - 1; row >= 0; row-- {
		// Vérifie si la case est vide
		if CurrentGame.Grid[row][col] == "| - |" {
			// Place le jeton du joueur actuel dans la case
			CurrentGame.Grid[row][col] = CurrentGame.Turn
			// Change le tour : si c'était X, passe à O, sinon passe à X
			if CurrentGame.Turn == "| X |" {
				CurrentGame.Turn = "| O |"
			} else {
				CurrentGame.Turn = "| X |"
			}
			// Sort de la fonction une fois le jeton placé
			return
		}
	}
}

func Check_Win_Con() bool {
	//Check Horizontal
	for i := 0; i < GameData.Rows; i++ {
		for j := 0; j < GameData.Cols-3; j++ {
			StartPt := CurrentGame.Grid[i][j]
			if StartPt != "| - |" {
				if StartPt == CurrentGame.Grid[i][j+1] &&
					StartPt == CurrentGame.Grid[i][j+2] &&
					StartPt == CurrentGame.Grid[i][j+3] {
					return true
				}
			}
		}
	}
	//Check vertical
	for i := 0; i < GameData.Rows-3; i++ {
		for j := 0; j < GameData.Cols; j++ {
			StartPt := CurrentGame.Grid[i][j]
			if StartPt != "| - |" {
				if StartPt == CurrentGame.Grid[i+1][j] &&
					StartPt == CurrentGame.Grid[i+2][j] &&
					StartPt == CurrentGame.Grid[i+3][j] {
					return true
				}
			}
		}
	}
	//Check Diagonal top left to bottom right
	for i := 0; i < GameData.Rows-3; i++ {
		for j := 0; j < GameData.Cols-3; j++ {
			StartPt := CurrentGame.Grid[i][j]
			if StartPt != "| - |" {
				if StartPt == CurrentGame.Grid[i+1][j+1] &&
					StartPt == CurrentGame.Grid[i+2][j+2] &&
					StartPt == CurrentGame.Grid[i+3][j+3] {
					return true
				}
			}
		}
	}
	//Check Diagonal bottom left to top right
	for i := 3; i < GameData.Rows; i++ {
		for j := 0; j < GameData.Cols-3; j++ {
			StartPt := CurrentGame.Grid[i][j]
			if StartPt != "| - |" {
				if StartPt == CurrentGame.Grid[i-1][j+1] &&
					StartPt == CurrentGame.Grid[i-2][j+2] &&
					StartPt == CurrentGame.Grid[i-3][j+3] {
					return true
				}
			}
		}
	}
	return false
}

func Winner() bool {
	if Check_Win_Con() {
		if CurrentGame.Turn == "| X |" {
			IncrementWin("O")
		} else if CurrentGame.Turn == "| O |" {
			IncrementWin("X")
		}
	}
	return Check_Win_Con()
}

// IncrementWin increments the win counter for the given player "X" or "O".
func IncrementWin(player string) {
	switch player {
	case "X":
		WinsX++
	case "O":
		WinsO++
	}
}

// GetWinCounts retourne les compteurs de victoires (winsX, winsO)
func GetWinCounts() (int, int) {
	return WinsX, WinsO
}

func CheckDraw() bool {
	if IsFull(CurrentGame.Grid) && !Check_Win_Con() {
		return true
	}
	return false
}

func GetGame() Game {
	return CurrentGame
}
