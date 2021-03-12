package main

import (
	"fmt"
	"net/http"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "localhost/bar?name=\"message!\""
	}
	fmt.Fprint(w, "%s", name)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hellow World!")
	})
	http.ListenAndServe(":3000", nil)
}
