package serveur

import (
	module "POWER4/src/modules"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

var mu sync.Mutex

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
			if module.GameData.Cols < 20 && module.GameData.Cols < module.GameData.Condition {
				module.GameData.Cols++
			}
		case "decrement_cols":
			if module.GameData.Cols > 4 && module.GameData.Cols > module.GameData.Condition {
				module.GameData.Cols--
			}
		case "increment_condition":
			if module.GameData.Condition < 7 {
				module.GameData.Condition++
				if module.GameData.Condition > module.GameData.Cols {
					module.GameData.Cols = module.GameData.Condition
				}
				if module.GameData.Condition > module.GameData.Rows {
					module.GameData.Rows = module.GameData.Condition
				}
			}
		case "decrement_condition":
			if module.GameData.Condition > 3 {
				module.GameData.Condition--
			}
		case "partie_classique":
			module.GameData.Rows = 6
			module.GameData.Cols = 7
			module.GameData.Condition = 4
			module.InitGameCustom(module.GameData.Rows, module.GameData.Cols, module.GameData.Condition)
			http.Redirect(w, r, "/game?new=1&classic=1", http.StatusSeeOther)
			return
		case "partie_perso":
			module.InitGameCustom(module.GameData.Rows, module.GameData.Cols, module.GameData.Condition)
			http.Redirect(w, r, "/game?new=1", http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmplPath := filepath.Join("templates", "index.html") // relatif à la racine
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Println("Erreur template index :", err)
		http.Error(w, "Erreur template : "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, module.GameData)
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	if r.URL.Query().Get("new") == "1" {
		if r.URL.Query().Get("classic") == "1" {
			// Nouvelle partie classique : initialise la grille et le tour mais garde les scores
			rows, cols, cond := 6, 7, 4
			winsX, winsO := module.CurrentGame.WinsX, module.CurrentGame.WinsO
			module.InitGameCustom(rows, cols, cond)
			module.CurrentGame.WinsX = winsX
			module.CurrentGame.WinsO = winsO
			module.GameData.Rows = 6
			module.GameData.Cols = 7
			module.GameData.Condition = 4
		} else {
			// Partie personnalisée : garde les scores
			winsX, winsO := module.CurrentGame.WinsX, module.CurrentGame.WinsO
			module.InitGameCustom(module.GameData.Rows, module.GameData.Cols, module.GameData.Condition)
			module.CurrentGame.WinsX = winsX
			module.CurrentGame.WinsO = winsO
		}
	}

	tmplPath := filepath.Join("templates", "game.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Println("Erreur template game :", err)
		http.Error(w, "Erreur template : "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, module.CurrentGame)
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Sécurité si la grille n'est pas initialisée
	if len(module.CurrentGame.Grid) == 0 {
		module.InitGameCustom(module.GameData.Rows, module.GameData.Cols, module.GameData.Condition)
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	action := r.FormValue("action")

	switch action {
	case "menu":
		// Retour au menu
		module.CurrentGame = module.Game{}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	case "reset":
		// Réinitialise la partie mais garde le score
		module.InitGameCustom(module.GameData.Rows, module.GameData.Cols, module.GameData.Condition)
		module.CurrentGame.Winner = ""
		http.Redirect(w, r, "/game", http.StatusSeeOther)
		return
	case "reset_scores":
		// Réinitialise la partie et les scores
		module.CurrentGame.WinsX = 0
		module.CurrentGame.WinsO = 0
		module.InitGameCustom(module.GameData.Rows, module.GameData.Cols, module.GameData.Condition)
		module.CurrentGame.Winner = ""
		http.Redirect(w, r, "/game", http.StatusSeeOther)
		return
	default:
		var col int
		fmt.Sscanf(r.FormValue("col"), "%d", &col)
		module.PlayMove(col)

		if module.Check_Win_Con() {
			if module.CurrentGame.Turn == "| X |" {
				module.CurrentGame.Winner = "O"
				module.CurrentGame.WinsO++
			} else {
				module.CurrentGame.Winner = "X"
				module.CurrentGame.WinsX++
			}
		}
		http.Redirect(w, r, "/game", http.StatusSeeOther)
	}
}

func Serveur() {
	// Fichiers statiques
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handlers
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/game", gameHandler)
	http.HandleFunc("/play", playHandler)

	fmt.Println("Serveur démarré sur http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
