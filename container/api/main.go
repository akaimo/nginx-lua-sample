package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hc", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})
	http.HandleFunc("/api/json", func(w http.ResponseWriter, r *http.Request) {
		v := struct {
			Ok      bool `json:"ok"`
			Message string `json:"message"`
		}{
			Ok:      true,
			Message: "ok!",
		}
		j, _ := json.Marshal(v)
		_, _ = w.Write(j)
	})
	http.HandleFunc("/api/error", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		v := struct {
			Ok      bool `json:"ok"`
			Message string `json:"message"`
		}{
			Ok:      false,
			Message: "error!!!",
		}
		j, _ := json.Marshal(v)
		_, _ = w.Write(j)
	})
	http.ListenAndServe(":8080", nil)
}
