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

type Handler struct{}

type Data struct {
	Head   html.Head
	Loader Loader
}

func (h *Handler) GetData() Data {
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
		"src/root.html",
		"src/routes/home/route.html",
	}
}
