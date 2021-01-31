package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hc", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})
	http.ListenAndServe(":8080", nil)
}
