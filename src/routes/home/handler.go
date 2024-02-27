package home

type Event struct {
	Name string
}

type Loader struct {
	Event Event
}

type Head struct {
	Title string
	// metas []Meta
	// links []Link
	// styles []Style
	// scripts []Script
}

type Handler struct {
	Head   Head
	Loader Loader
}

func (d *Handler) GetData() Handler {
	return Handler{
		Head: Head{
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
