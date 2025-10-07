package serveur

import (
	module "POWER4/src/modules"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	execDir, err := os.Getwd()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération du répertoire : "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmplPath := filepath.Join(execDir, "templates", "index.html")

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Erreur de chargement du template : "+err.Error(), http.StatusInternalServerError)
		return
	}

	plateau := module.Plateau{
		Colonnes: [][]module.Cell{
			{{Valeur: "B"}, {Valeur: "R"}, {}, {}, {}, {}},
			{{Valeur: "R"}, {}, {}, {}, {}, {}},
			{{}, {}, {}, {}, {}, {}},
			{{}, {}, {}, {}, {}, {}},
			{{}, {}, {}, {}, {}, {}},
			{{}, {}, {}, {}, {}, {}},
			{{}, {}, {}, {}, {}, {}},
		},
	}

	if err := tmpl.Execute(w, plateau); err != nil {
		http.Error(w, "Erreur lors de l'exécution du template : "+err.Error(), http.StatusInternalServerError)
	}
}

func Serveur() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", indexHandler)

	fmt.Println("Serveur démarré sur http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
