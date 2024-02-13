package core

import "github.com/google/uuid"

// Viewer represents a user who is interacting with the platform in some way, either
// directly via Twitch or via the website, authenticated via Twitch
type Viewer struct {
	TwitchUserId      string `json:"twitch_user_id"`
	TwitchDisplayName string `json:"twitch_display_name"`
}

// State describes the current broadcast state, derived from the latest series of events
// that have been produced to broadcast-events
type State struct {
	BroadcastId int       `json:"broadcast_id"`
	ScreeningId uuid.UUID `json:"screening_id"`
	TapeId      int       `json:"tape_id"`
}
