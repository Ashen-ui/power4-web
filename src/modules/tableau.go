package module

import "fmt"

func InitPlateau() [6][7]string {
	var plateau [6][7]string = [6][7]string{
		[7]string{"| - |", "| - |", "| - |", "| - |", "| - |", "| - |", "| - |"},
		[7]string{"| - |", "| - |", "| - |", "| - |", "| - |", "| - |", "| - |"},
		[7]string{"| - |", "| - |", "| - |", "| - |", "| - |", "| - |", "| - |"},
		[7]string{"| - |", "| - |", "| - |", "| - |", "| - |", "| - |", "| - |"},
		[7]string{"| - |", "| - |", "| - |", "| - |", "| - |", "| - |", "| - |"},
		[7]string{"| - |", "| - |", "| - |", "| - |", "| - |", "| - |", "| - |"},
	}
	return plateau
}

// color = "| X |" ou "| O |"
func AddToPlateau(X int, Y int, color string, plateau *[6][7]string) {
	plateau[X][Y] = color
}

func PrintPlateau(plateau *[6][7]string) {
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(plateau[i][j])
		}
	}
}
