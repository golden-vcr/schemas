package eonscreen

import (
	genreq "github.com/golden-vcr/schemas/generation-requests"
	etwitch "github.com/golden-vcr/schemas/twitch-events"
)

type PayloadImage struct {
	Viewer      etwitch.Viewer    `json:"viewer"`
	Style       genreq.ImageStyle `json:"style"`
	Description string            `json:"description"`
	ImageUrl    string            `json:"image_url"`
}
