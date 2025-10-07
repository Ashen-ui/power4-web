package module

// Game représente l'état d'une partie de Puissance 4
type Game struct {
	Grid    [6][7]string
	Turn    string
	Joueur1 Joueur
	Joueur2 Joueur
}

var CurrentGame Game

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

// GetGame retourne l'état actuel de la partie
func GetGame() Game {
	return CurrentGame
}
