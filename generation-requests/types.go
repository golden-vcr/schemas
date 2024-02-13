package genreq

import (
	"encoding/json"

	"github.com/golden-vcr/schemas/core"
)

// Request represents a payload produced to the 'generation-requests' queue in order to
// kick off the processing required for a cheer that requests some kind of asynchronous
// asset generation
type Request struct {
	Type    RequestType `json:"type"`
	Viewer  core.Viewer `json:"viewer"`
	Payload Payload     `json:"payload"`
}

// RequestType describes the kind of asset(s) we want to generate from this request
type RequestType string

const (
	RequestTypeImage RequestType = "image"
)

// Payload carries request-type-specific data describing the details of the request
type Payload struct {
	Image *PayloadImage
}

func (e *Request) UnmarshalJSON(data []byte) error {
	type fields struct {
		Type    RequestType     `json:"type"`
		Viewer  core.Viewer     `json:"viewer"`
		Payload json.RawMessage `json:"payload"`
	}
	var f fields
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}

	e.Type = f.Type
	e.Viewer = f.Viewer
	switch f.Type {
	case RequestTypeImage:
		return json.Unmarshal(f.Payload, &e.Payload.Image)
	}
	return nil
}

func (p Payload) MarshalJSON() ([]byte, error) {
	if p.Image != nil {
		return json.Marshal(p.Image)
	}
	return json.Marshal(nil)
}
