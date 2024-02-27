package main

import (
	"html/template"
	"log"
	"net/http"
)

type Event struct {
	Name string
}

type Head struct {
	Title string
	// metas []Meta
	// links []Link
	// styles []Style
	// scripts []Script
}

type Loader struct {
	Event Event
}

type Data[T any] struct {
	Head   Head
	Loader T
}

type HomeLoader struct {
	Event Event
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(
			template.ParseFiles(
				"src/root.html",
				"src/routes/home.html",
			),
		)
		t.Execute(w, Data[HomeLoader]{
			Head:   Head{Title: "Home page"},
			Loader: HomeLoader{Event: Event{Name: "Home page"}},
		})
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
