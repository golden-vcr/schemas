package ebroadcast

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_Event(t *testing.T) {
	tests := []struct {
		name   string
		ev     Event
		jsonEv string
	}{
		{
			"broadcast started",
			Event{
				Type: EventTypeBroadcastStarted,
				Broadcast: BroadcastData{
					Id:        55,
					StartedAt: time.Date(1997, 9, 1, 12, 0, 0, 0, time.UTC),
				},
			},
			`{"type":"broadcast-started","broadcast":{"id":55,"started_at":"1997-09-01T12:00:00Z"}}`,
		},
		{
			"broadcast finished",
			Event{
				Type: EventTypeBroadcastFinished,
				Broadcast: BroadcastData{
					Id:        55,
					StartedAt: time.Date(1997, 9, 1, 12, 0, 0, 0, time.UTC),
				},
			},
			`{"type":"broadcast-finished","broadcast":{"id":55,"started_at":"1997-09-01T12:00:00Z"}}`,
		},
		{
			"screening started",
			Event{
				Type: EventTypeScreeningStarted,
				Broadcast: BroadcastData{
					Id:        55,
					StartedAt: time.Date(1997, 9, 1, 12, 0, 0, 0, time.UTC),
				},
				Screening: &ScreeningData{
					Id:        uuid.MustParse("f29a4ffe-cb9f-43ba-9f91-a3b1fa350472"),
					StartedAt: time.Date(1997, 9, 1, 12, 15, 0, 0, time.UTC),
					TapeId:    109,
				},
			},
			`{"type":"screening-started","broadcast":{"id":55,"started_at":"1997-09-01T12:00:00Z"},"screening":{"id":"f29a4ffe-cb9f-43ba-9f91-a3b1fa350472","started_at":"1997-09-01T12:15:00Z","tape_id":109}}`,
		},
		{
			"screening started",
			Event{
				Type: EventTypeScreeningFinished,
				Broadcast: BroadcastData{
					Id:        55,
					StartedAt: time.Date(1997, 9, 1, 12, 0, 0, 0, time.UTC),
				},
				Screening: &ScreeningData{
					Id:        uuid.MustParse("f29a4ffe-cb9f-43ba-9f91-a3b1fa350472"),
					StartedAt: time.Date(1997, 9, 1, 12, 15, 0, 0, time.UTC),
					TapeId:    109,
				},
			},
			`{"type":"screening-finished","broadcast":{"id":55,"started_at":"1997-09-01T12:00:00Z"},"screening":{"id":"f29a4ffe-cb9f-43ba-9f91-a3b1fa350472","started_at":"1997-09-01T12:15:00Z","tape_id":109}}`,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("marshal %s to JSON", tt.name), func(t *testing.T) {
			want := tt.jsonEv
			got, err := json.Marshal(tt.ev)
			assert.NoError(t, err)
			assert.Equal(t, want, string(got))
		})
		t.Run(fmt.Sprintf("unmarshal %s from JSON", tt.name), func(t *testing.T) {
			want := tt.ev
			var got Event
			err := json.Unmarshal([]byte(tt.jsonEv), &got)
			assert.NoError(t, err)
			assert.Equal(t, want, got)
		})
	}
}
