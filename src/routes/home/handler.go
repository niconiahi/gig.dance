package home

import (
	"github.com/niconiahi/gig.dance/src/utils/html"
)

type Event struct {
	Name string
}

type Loader struct {
	Event Event
}

type Handler struct {
	Head   html.Head
	Loader Loader
}

func (d *Handler) GetData() Handler {
	return Handler{
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

func (d *Handler) GetFiles() []string {
	return []string{
		"src/root.html",
		"src/routes/home/route.html",
	}
}
