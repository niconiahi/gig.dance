package home

import (
	"embed"
	"html/template"
	"log"

	"github.com/niconiahi/gig.dance/packages/db"
	"github.com/niconiahi/gig.dance/packages/html"
)

type Event struct {
	Name string
}

type Loader struct {
	Event Event
}

type Handler struct{}

type Data struct {
	Head   html.Head
	Loader Loader
}

func (h *Handler) GetData() Data {
	db := db.GetDb()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	return Data{
		Head: html.Head{
			Title: "Home page",
		},
		Loader: Loader{
			Event: Event{
				Name: "Ryan Elliot - Live at Panorama Bar",
			},
		},
	}
}

func (h *Handler) GetFiles() []string {
	return []string{
		"root.html",
		"route.html",
	}
}

//go:embed route.html root.html
var files embed.FS

func (h *Handler) GetTemplate() (*template.Template, error) {
	return template.ParseFS(files, h.GetFiles()...)
}