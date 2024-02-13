package ebroadcast

import (
	"time"

	"github.com/golden-vcr/schemas/core"
	"github.com/google/uuid"
)

// Event represents a change in the overall broadcast state
type Event struct {
	Type      EventType      `json:"type"`
	Broadcast BroadcastData  `json:"broadcast"`
	Screening *ScreeningData `json:"screening,omitempty"`
}

// EventType indicates what state change has taken place
type EventType string

const (
	EventTypeBroadcastStarted  EventType = "broadcast-started"
	EventTypeBroadcastFinished EventType = "broadcast-finished"
	EventTypeScreeningStarted  EventType = "screening-started"
	EventTypeScreeningFinished EventType = "screening-finished"
)

// BroadcastData describes the broadcast in which this event is occurring
type BroadcastData struct {
	Id        int       `json:"id"`
	StartedAt time.Time `json:"started_at"`
}

// ScreeningData describes the screening in which this event is occurring, for screening
// events only
type ScreeningData struct {
	Id        uuid.UUID `json:"id"`
	StartedAt time.Time `json:"started_at"`
	TapeId    int       `json:"tape_id"`
}

func (ev *Event) ToState(prev core.State) core.State {
	switch ev.Type {
	case EventTypeBroadcastStarted:
		return core.State{
			BroadcastId: ev.Broadcast.Id,
		}
	case EventTypeBroadcastFinished:
		return core.State{}
	case EventTypeScreeningStarted:
		return core.State{
			BroadcastId: ev.Broadcast.Id,
			ScreeningId: ev.Screening.Id,
			TapeId:      ev.Screening.TapeId,
		}
	case EventTypeScreeningFinished:
		return core.State{
			BroadcastId: ev.Broadcast.Id,
		}
	}
	return prev
}
