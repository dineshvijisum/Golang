package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>hello kirthi<h1>")
	})
	http.ListenAndServe(":9000", nil)
}
