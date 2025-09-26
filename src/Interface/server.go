package main

import (
	"fmt"
	"log"
	"net/http"
)

// REFERENCE: https://blog.logrocket.com/creating-a-web-server-with-golang
// Lance un serveur web basique sur le port 8080 pour servir les fichiers statiques
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// Pour tester, aller sur http://localhost:8080 dans un navigateur
// Pour arrêter le serveur, faire Ctrl+C dans le terminal
