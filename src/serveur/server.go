package serveur

import (
	"fmt"
	"log"
	"net/http"
)

// REFERENCE: https://blog.logrocket.com/creating-a-web-server-with-golang
func Serveur() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
