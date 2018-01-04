package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pwd(w, r)

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func pwd(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CWD: %v", html.EscapeString(r.URL.Path))
}
