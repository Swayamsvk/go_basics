package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Setver starting...")
	http.ListenAndServe(":3000", nil)
}
