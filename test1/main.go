package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/greet", greet)
	http.ListenAndServe(":3000", nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
