package main

import (
	"html/template"
	"log"
	"net/http"
)

type Event struct {
	Name string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(
			template.ParseFiles(
				"src/root.html",
				"src/routes/home.html",
			),
		)
		t.Execute(w, Event{Name: "Jodota 22"})
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
