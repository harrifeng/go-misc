package main

import (
	"fmt"
	"os"
	"net/http"
)

func main() {

	http.HandleFunc("/", demoHandler)
	http.ListenAndServe(":12345", nil)
	fmt.Println()
	os.Exit(0)
}

func demoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The path is %s", r.URL.Path)
}
