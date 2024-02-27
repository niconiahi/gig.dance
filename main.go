package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/niconiahi/go-htmx-demo/src/routes/home"
)

type Head struct {
	Title string
	// metas []Meta
	// links []Link
	// styles []Style
	// scripts []Script
}

type Data[T any] struct {
	Head   Head
	Loader T
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		h := home.Handler{}
		t := template.Must(template.ParseFiles(h.GetFiles()...))
		t.Execute(w, h.GetData())
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
