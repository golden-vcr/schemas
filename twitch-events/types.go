package etwitch

import (
	"encoding/json"
)

type EventType string

const (
	EventTypeStreamStarted         EventType = "stream-started"
	EventTypeStreamEnded           EventType = "stream-ended"
	EventTypeViewerFollowed        EventType = "viewer-followed"
	EventTypeViewerRaided          EventType = "viewer-raided"
	EventTypeViewerCheered         EventType = "viewer-cheered"
	EventTypeViewerSubscribed      EventType = "viewer-subscribed"
	EventTypeViewerResubscribed    EventType = "viewer-resubscribed"
	EventTypeViewerReceivedGiftSub EventType = "viewer-received-gift-sub"
	EventTypeViewerGiftedSubs      EventType = "viewer-gifted-subs"
)

// Event is an event that has occurred on Twitch, such as a viewer interaction or a
// change in the state of the stream
type Event struct {
	Type    EventType `json:"type"`
	Viewer  *Viewer   `json:"viewer"`
	Payload *Payload  `json:"payload"`
}

type Viewer struct {
	TwitchUserId      string `json:"twitch_user_id"`
	TwitchDisplayName string `json:"twitch_display_name"`
}

type Payload struct {
	ViewerRaided          *PayloadViewerRaided
	ViewerCheered         *PayloadViewerCheered
	ViewerSubscribed      *PayloadViewerSubscribed
	ViewerResubscribed    *PayloadViewerResubscribed
	ViewerReceivedGiftSub *PayloadViewerReceivedGiftSub
	ViewerGiftedSubs      *PayloadViewerGiftedSubs
}

func (p Payload) MarshalJSON() ([]byte, error) {
	if p.ViewerRaided != nil {
		return json.Marshal(p.ViewerRaided)
	}
	if p.ViewerCheered != nil {
		return json.Marshal(p.ViewerCheered)
	}
	if p.ViewerSubscribed != nil {
		return json.Marshal(p.ViewerSubscribed)
	}
	if p.ViewerResubscribed != nil {
		return json.Marshal(p.ViewerResubscribed)
	}
	if p.ViewerReceivedGiftSub != nil {
		return json.Marshal(p.ViewerReceivedGiftSub)
	}
	if p.ViewerGiftedSubs != nil {
		return json.Marshal(p.ViewerGiftedSubs)
	}
	return json.Marshal(nil)
}

type PayloadViewerRaided struct {
	NumRaiders int `json:"num_raiders"`
}

type PayloadViewerCheered struct {
	NumBits int    `json:"num_bits"`
	Message string `json:"message"`
}

type PayloadViewerSubscribed struct {
	CreditMultiplier int `json:"credit_multiplier"`
}

type PayloadViewerResubscribed struct {
	CreditMultiplier    int    `json:"credit_multiplier"`
	NumCumulativeMonths int    `json:"num_cumulative_months"`
	Message             string `json:"message"`
}

type PayloadViewerReceivedGiftSub struct {
	CreditMultiplier int `json:"credit_multiplier"`
}

type PayloadViewerGiftedSubs struct {
	CreditMultiplier int `json:"credit_multiplier"`
	NumSubscriptions int `json:"num_subscriptions"`
}
