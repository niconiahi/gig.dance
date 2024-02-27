package handler

import (
	"github.com/niconiahi/gig.dance/src/utils/html"
)

type Data[T any] struct {
	Head   html.Head
	Loader T
}
