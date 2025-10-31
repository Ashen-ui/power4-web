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
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	if r.Method == http.MethodPost {
		action := r.FormValue("action")
		switch action {
		case "increment_rows":
			if module.GameData.Rows < 20 {
				module.GameData.Rows++
			}
		case "decrement_rows":
			if module.GameData.Rows > 4 {
				module.GameData.Rows--
			}
		case "increment_cols":
			if module.GameData.Cols < 20 {
				module.GameData.Cols++
			}
		case "decrement_cols":
			if module.GameData.Cols > 4 {
				module.GameData.Cols--
			}
		case "increment_condition":
			if module.GameData.Condition < 7 {
				module.GameData.Condition++
			}
		case "decrement_condition":
			if module.GameData.Condition > 4 {
				module.GameData.Condition--
			}
		case "set_classic":
			module.GameData.Rows = 6
			module.GameData.Cols = 7
			module.GameData.Condition = 4
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
	tmpl.Execute(w, module.GameData)
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
			module.GameData.Rows = 6
			module.GameData.Cols = 7
			module.GameData.Condition = 4
		} else {
			// Partie personnalisée
			module.InitGameCustom(module.GameData.Rows, module.GameData.Cols, module.GameData.Condition)
		}
	}

	tmplPath := filepath.Join(execDir, "templates", "game.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Erreur template : "+err.Error(), http.StatusInternalServerError)
		return
	}

	plateau := module.GetGame()
	view := GameView{Colonnes: make([][]Cell, module.GameData.Cols)}

	for col := 0; col < module.GameData.Cols; col++ {
		colCells := make([]Cell, module.GameData.Rows)
		for row := 0; row < module.GameData.Rows; row++ {
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
