package module

// Game représente l'état d'une partie de Puissance 4
type Game struct {
	Grid    [6][7]string
	Turn    string
	Joueur1 Joueur
	Joueur2 Joueur
}

var CurrentGame Game
var WinsX int
var WinsO int

// InitGame initialise une nouvelle partie avec un plateau vide
func InitGame() {
	CurrentGame = Game{}
	CurrentGame.Grid = InitPlateau()
	CurrentGame.Turn = "| X |"
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

func Check_Win_Con() string {
	//Check Horizontal
	for i := 0; i < 6; i++ {
		for j := 0; j < 4; j++ {
			StartPt := CurrentGame.Grid[i][j]
			if StartPt != "| - |" {
				if StartPt == CurrentGame.Grid[i][j+1] &&
					StartPt == CurrentGame.Grid[i][j+2] &&
					StartPt == CurrentGame.Grid[i][j+3] {
					return StartPt
				}
			}
		}
	}
	//Check vertical
	for i := 0; i < 3; i++ {
		for j := 0; j < 7; j++ {
			StartPt := CurrentGame.Grid[i][j]
			if StartPt != "| - |" {
				if StartPt == CurrentGame.Grid[i+1][j] &&
					StartPt == CurrentGame.Grid[i+2][j] &&
					StartPt == CurrentGame.Grid[i+3][j] {
					return StartPt
				}
			}
		}
	}
	//Check Diagonal top left to bottom right
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			StartPt := CurrentGame.Grid[i][j]
			if StartPt != "| - |" {
				if StartPt == CurrentGame.Grid[i+1][j+1] &&
					StartPt == CurrentGame.Grid[i+2][j+2] &&
					StartPt == CurrentGame.Grid[i+3][j+3] {
					return StartPt
				}
			}
		}
	}
	//Check Diagonal bottom left to top right
	for i := 3; i < 6; i++ {
		for j := 0; j < 4; j++ {
			StartPt := CurrentGame.Grid[i][j]
			if StartPt != "| - |" {
				if StartPt == CurrentGame.Grid[i-1][j+1] &&
					StartPt == CurrentGame.Grid[i-2][j+2] &&
					StartPt == CurrentGame.Grid[i-3][j+3] {
					return StartPt
				}
			}
		}
	}
	return "| - |"
}

func Winner() string {
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

func Reset() {
	InitGame()
}

func CheckDraw() string {
	if IsFull(CurrentGame.Grid) && Check_Win_Con() == "| - |" {
		return "draw"
	}
	return "| - |"
}

func GetGame() Game {
	return CurrentGame
}
