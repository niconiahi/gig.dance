package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/niconiahi/gig.dance/packages/db"
	"github.com/niconiahi/gig.dance/routes/home"
)

func main() {
	db := db.GetDb()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		h := home.Handler{}
		t := template.Must(h.GetTemplate())
		t.Execute(w, h.GetData())
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
