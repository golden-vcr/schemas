package eonscreen

import (
	"github.com/golden-vcr/schemas/core"
	genreq "github.com/golden-vcr/schemas/generation-requests"
)

type PayloadImage struct {
	Viewer      core.Viewer       `json:"viewer"`
	Style       genreq.ImageStyle `json:"style"`
	Description string            `json:"description"`
	Extra       string            `json:"extra"`
	ImageUrl    string            `json:"image_url"`
}
