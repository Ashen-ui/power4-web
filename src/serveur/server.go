package serveur

import (
	module "POWER4/src/modules"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

// Cell is a view-model cell used by templates
type Cell struct{ Valeur string }

// GameView is the view model passed to game.html
type GameView struct {
	Colonnes [][]Cell
	Current  string
	Winner   string
	WinsX    int
	WinsO    int
}

// variables globales pour les paramètres personnalisés
var (
	mu sync.Mutex

	gameData = struct {
		Rows      int
		Cols      int
		Condition int
	}{
		Rows:      6,
		Cols:      7,
		Condition: 4,
	}
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	if r.Method == http.MethodPost {
		action := r.FormValue("action")
		switch action {
		case "increment_rows":
			if gameData.Rows < 20 {
				gameData.Rows++
			}
		case "decrement_rows":
			if gameData.Rows > 4 {
				gameData.Rows--
			}
		case "increment_cols":
			if gameData.Cols < 20 {
				gameData.Cols++
			}
		case "decrement_cols":
			if gameData.Cols > 4 {
				gameData.Cols--
			}
		case "increment_condition":
			if gameData.Condition < 7 {
				gameData.Condition++
			}
		case "decrement_condition":
			if gameData.Condition > 4 {
				gameData.Condition--
			}
		case "set_classic":
			gameData.Rows = 6
			gameData.Cols = 7
			gameData.Condition = 4
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	execDir, _ := os.Getwd()
	tmplPath := filepath.Join(execDir, "templates", "index.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Erreur template : "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, gameData)
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	execDir, _ := os.Getwd()

	// Nouvelle partie demandée
	if r.URL.Query().Get("new") == "1" {
		if r.URL.Query().Get("classic") == "1" {
			// Valeurs classiques fixes
			module.InitGameCustom(6, 7, 4)
			gameData.Rows = 6
			gameData.Cols = 7
			gameData.Condition = 4
		} else {
			// Partie personnalisée
			module.InitGameCustom(gameData.Rows, gameData.Cols, gameData.Condition)
		}
	}

	tmplPath := filepath.Join(execDir, "templates", "game.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Erreur template : "+err.Error(), http.StatusInternalServerError)
		return
	}

	plateau := module.GetGame()
	view := GameView{Colonnes: make([][]Cell, gameData.Cols)}

	for col := 0; col < gameData.Cols; col++ {
		colCells := make([]Cell, gameData.Rows)
		for row := 0; row < gameData.Rows; row++ {
			val := ""
			if row < len(plateau.Grid) && col < len(plateau.Grid[row]) {
				switch plateau.Grid[row][col] {
				case "| X |":
					val = "R"
				case "| O |":
					val = "B"
				}
			}
			colCells[row] = Cell{Valeur: val}
		}
		view.Colonnes[col] = colCells
	}

	if plateau.Turn == "| X |" {
		view.Current = "X"
	} else if plateau.Turn == "| O |" {
		view.Current = "O"
	}

	if module.Check_Win_Con() == true {
		winner := ""
		if plateau.Turn == "| X |" {
			winner = "O"
		} else if plateau.Turn == "| O |" {
			winner = "X"
		}
		view.Winner = winner
	}

	wx, wo := module.GetWinCounts()
	view.WinsX = wx
	view.WinsO = wo

	tmpl.Execute(w, view)
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	colStr := r.FormValue("col")
	var col int
	fmt.Sscanf(colStr, "%d", &col)
	module.PlayMove(col)

	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

func Serveur() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/game", gameHandler)
	http.HandleFunc("/play", playHandler)

	fmt.Println("Serveur démarré sur http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
