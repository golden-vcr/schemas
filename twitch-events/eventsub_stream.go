package etwitch

import "encoding/json"

func fromStreamOnlineEvent(data json.RawMessage) (*Event, error) {
	return &Event{
		Type: EventTypeStreamStarted,
	}, nil
}

func fromStreamOfflineEvent(data json.RawMessage) (*Event, error) {
	return &Event{
		Type: EventTypeStreamEnded,
	}, nil
}
