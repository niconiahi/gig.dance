package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/niconiahi/gig.dance/routes/home"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		h := home.Handler{}
		t := template.Must(h.GetTemplate())
		t.Execute(w, h.GetData())
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
