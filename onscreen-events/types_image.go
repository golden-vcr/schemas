package eonscreen

import (
	"encoding/json"

	"github.com/golden-vcr/schemas/core"
)

type ImageType string

const (
	ImageTypeStatic ImageType = "static"
	ImageTypeGhost  ImageType = "ghost"
	ImageTypeFriend ImageType = "friend"
)

type PayloadImage struct {
	Type    ImageType    `json:"type"`
	Viewer  core.Viewer  `json:"viewer"`
	Details ImageDetails `json:"details"`
}

type ImageDetails struct {
	Static *ImageDetailsStatic
	Ghost  *ImageDetailsGhost
	Friend *ImageDetailsFriend
}

func (p *PayloadImage) UnmarshalJSON(data []byte) error {
	type fields struct {
		Type    ImageType       `json:"type"`
		Viewer  core.Viewer     `json:"viewer"`
		Details json.RawMessage `json:"details"`
	}
	var f fields
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}

	p.Type = f.Type
	p.Viewer = f.Viewer
	switch f.Type {
	case ImageTypeStatic:
		return json.Unmarshal(f.Details, &p.Details.Static)
	case ImageTypeGhost:
		return json.Unmarshal(f.Details, &p.Details.Ghost)
	case ImageTypeFriend:
		return json.Unmarshal(f.Details, &p.Details.Friend)
	}
	return nil
}

func (d ImageDetails) MarshalJSON() ([]byte, error) {
	if d.Static != nil {
		return json.Marshal(d.Static)
	}
	if d.Ghost != nil {
		return json.Marshal(d.Ghost)
	}
	if d.Friend != nil {
		return json.Marshal(d.Friend)
	}
	return json.Marshal(nil)
}

type ImageDetailsStatic struct {
	ImageId string `json:"image_id"`
	Message string `json:"message"`
}

type ImageDetailsGhost struct {
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
}

type ImageDetailsFriend struct {
	ImageUrl        string `json:"image_url"`
	Description     string `json:"description"`
	Name            string `json:"name"`
	BackgroundColor string `json:"background_color"`
}
