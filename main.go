package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/niconiahi/gig.dance/src/routes/home"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	url := fmt.Sprintf(
		"libsql://%s.turso.io?authToken=%s",
		os.Getenv("DATABASE"),
		os.Getenv("TOKEN"),
	)

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		h := home.Handler{}
		t := template.Must(h.GetTemplate())
		t.Execute(w, h.GetData())
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
