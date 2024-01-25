package eonscreen

import "encoding/json"

// Event represents an event that should be displayed onscreen during the stream
type Event struct {
	Type    EventType `json:"type"`
	Payload Payload   `json:"payload"`
}

// EventType indicates the type of event (e.g. change in stream status, toast
// recognizing a viewer interaction, display of a generated image), all of which are
// displayed differently in the onscreen graphics
type EventType string

const (
	EventTypeStatus EventType = "status"
	EventTypeToast  EventType = "toast"
	EventTypeImage  EventType = "image"
)

// Payload carries event-type-specific data describing what needs to happen onscreen
type Payload struct {
	Status *PayloadStatus
	Toast  *PayloadToast
	Image  *PayloadImage
}

func (e *Event) UnmarshalJSON(data []byte) error {
	type fields struct {
		Type    EventType       `json:"type"`
		Payload json.RawMessage `json:"payload"`
	}
	var f fields
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}

	e.Type = f.Type
	switch f.Type {
	case EventTypeStatus:
		return json.Unmarshal(f.Payload, &e.Payload.Status)
	case EventTypeToast:
		return json.Unmarshal(f.Payload, &e.Payload.Toast)
	case EventTypeImage:
		return json.Unmarshal(f.Payload, &e.Payload.Image)
	}
	return nil
}

func (p Payload) MarshalJSON() ([]byte, error) {
	if p.Status != nil {
		return json.Marshal(p.Status)
	}
	if p.Toast != nil {
		return json.Marshal(p.Toast)
	}
	if p.Image != nil {
		return json.Marshal(p.Image)
	}
	return json.Marshal(nil)
}
