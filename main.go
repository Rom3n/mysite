package main

import (
	"fmt"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Initial commit")
}

func main() {
	http.HandleFunc("/", home_page)
	http.ListenAndServe("localhost:8080", nil)
}
