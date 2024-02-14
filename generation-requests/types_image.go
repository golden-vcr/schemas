package genreq

import "encoding/json"

// PayloadImage describes a request to generate one or more images for an alert
type PayloadImage struct {
	Style  ImageStyle  `json:"style"`
	Inputs ImageInputs `json:"inputs"`
}

// ImageStyle represents the style of alert we want to generate an image for
type ImageStyle string

const (
	ImageStyleGhost   ImageStyle = "ghost"
	ImageStyleClipArt ImageStyle = "clip-art"
)

// ImageInputs contains the user-provided information that we'll use to build a prompt
// for our image generation request; these inputs vary by image style
type ImageInputs struct {
	Ghost   *ImageInputsGhost
	ClipArt *ImageInputsClipArt
}

func (e *PayloadImage) UnmarshalJSON(data []byte) error {
	type fields struct {
		Style  ImageStyle      `json:"style"`
		Inputs json.RawMessage `json:"inputs"`
	}
	var f fields
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}

	e.Style = f.Style
	switch f.Style {
	case ImageStyleGhost:
		return json.Unmarshal(f.Inputs, &e.Inputs.Ghost)
	case ImageStyleClipArt:
		return json.Unmarshal(f.Inputs, &e.Inputs.ClipArt)
	}
	return nil
}

func (i ImageInputs) MarshalJSON() ([]byte, error) {
	if i.Ghost != nil {
		return json.Marshal(i.Ghost)
	}
	if i.ClipArt != nil {
		return json.Marshal(i.ClipArt)
	}
	return json.Marshal(nil)
}

type ImageInputsGhost struct {
	Subject string `json:"subject"`
}

type ImageInputsClipArt struct {
	Color   string `json:"color"`
	Subject string `json:"subject"`
}
