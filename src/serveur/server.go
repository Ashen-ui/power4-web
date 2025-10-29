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

// permet de créer des variables globales pour les paramètres personnalisés pour faire des parties persionalisé
var (
	value     int
	rows      = 6
	cols      = 7
	condition = 4
	mu        sync.Mutex
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	if r.Method == http.MethodPost {
		action := r.FormValue("action")
		switch action {
		case "increment":
			value++
		case "decrement":
			value--
		case "increment_rows":
			if rows < 20 {
				rows++
			}
		case "decrement_rows":
			if rows > 4 {
				rows--
			}
		case "increment_cols":
			if cols < 20 {
				cols++
			}
		case "decrement_cols":
			if cols > 4 {
				cols--
			}
		case "increment_condition":
			if condition < 7 {
				condition++
			}
		case "decrement_condition":
			if condition > 4 {
				condition--
			}
		}
	}

	execDir, err := os.Getwd()
	if err != nil {
		http.Error(w, "Erreur répertoire : "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmplPath := filepath.Join(execDir, "templates", "index.html")

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Erreur template : "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Value     int
		Rows      int
		Cols      int
		Condition int
	}{
		Value:     value,
		Rows:      rows,
		Cols:      cols,
		Condition: condition,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Erreur exécution : "+err.Error(), http.StatusInternalServerError)
	}
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	execDir, err := os.Getwd()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération du répertoire : "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Initialize a new game only if requested (?new=1)
	if r.URL.Query().Get("new") == "1" {
		module.InitGame()
	}

	tmplPath := filepath.Join(execDir, "templates", "game.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Erreur de chargement du template : "+err.Error(), http.StatusInternalServerError)
		return
	}

	plateau := module.GetGame()

	// initialize view using package-level GameView/Cell
	view := GameView{Colonnes: make([][]Cell, 7)}

	for col := 0; col < 7; col++ {
		colCells := make([]Cell, 0, 6)
		for row := 0; row < 6; row++ {
			val := ""
			switch plateau.Grid[row][col] {
			case "| X |":
				val = "R"
			case "| O |":
				val = "B"
			default:
				val = ""
			}
			colCells = append(colCells, Cell{Valeur: val})
		}
		view.Colonnes[col] = colCells
	}

	// Current player mapping: internal "| X |" -> "X" and "| O |" -> "O"
	view.Current = ""
	if module.GetGame().Turn == "| X |" {
		view.Current = "X"
	} else if module.GetGame().Turn == "| O |" {
		view.Current = "O"
	}

	// Check for a winner
	if winner, ok := module.CheckWin(module.GetGame().Grid); ok {
		view.Winner = winner
	} else {
		view.Winner = ""
	}

	// populate win counters
	wx, wo := module.GetWinCounts()
	view.WinsX = wx
	view.WinsO = wo

	if err := tmpl.Execute(w, view); err != nil {
		http.Error(w, "Erreur lors de l'exécution du template : "+err.Error(), http.StatusInternalServerError)
	}
}

// playHandler accepts a POST with form value "col" (0-6) and plays the move, then redirects back to /game
func playHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}
	colStr := r.FormValue("col")
	if colStr == "" {
		http.Error(w, "Colonne manquante", http.StatusBadRequest)
		return
	}
	// parse integer
	var col int
	_, err := fmt.Sscanf(colStr, "%d", &col)
	if err != nil || col < 0 || col > 6 {
		http.Error(w, "Colonne invalide", http.StatusBadRequest)
		return
	}

	// Play move
	module.PlayMove(col)

	// After playing, check for a winner
	if winner, ok := module.CheckWin(module.GetGame().Grid); ok {
		// increment counters
		switch winner {
		case "X":
			module.IncrementWin("X")
		case "O":
			module.IncrementWin("O")
		}
		// Redirect to /game to display the winner
		http.Redirect(w, r, "/game", http.StatusSeeOther)
		return
	}

	// Redirect back to game view
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

func Serveur() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/game", gameHandler)
	http.HandleFunc("/play", playHandler)

	fmt.Println("Serveur démarré sur http://localhost:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
