package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/niconiahi/gig.dance/src/routes/home"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		h := home.Handler{}
		t := template.Must(template.ParseFiles(h.GetFiles()...))
		t.Execute(w, h.GetData())
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
