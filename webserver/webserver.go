package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Printf("Starting server at port 8081\n")
	mux := http.NewServeMux()
	mux.HandleFunc("/", redirectToDocumentation)
	http.ListenAndServe(":8081", mux)

}

func redirectToDocumentation(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://go.dev/doc/", http.StatusSeeOther)
	return

}
