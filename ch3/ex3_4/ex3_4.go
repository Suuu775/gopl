package ex34

import (
	"fmt"
	"log"
	"net/http"

	ex33 "github.com/Suuu775/gopl/ch3/ex3_3"
)

func Ex34() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintln(w, ex33.Surface())
}
