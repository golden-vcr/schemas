package etwitch

import (
	"encoding/json"

	"github.com/golden-vcr/schemas/core"
)

type EventType string

const (
	EventTypeStreamStarted           EventType = "stream-started"
	EventTypeStreamEnded             EventType = "stream-ended"
	EventTypeStreamHypeStarted       EventType = "stream-hype-started"
	EventTypeViewerFollowed          EventType = "viewer-followed"
	EventTypeViewerRaided            EventType = "viewer-raided"
	EventTypeViewerCheered           EventType = "viewer-cheered"
	EventTypeViewerRedeemedFunPoints EventType = "viewer-redeemed-fun-points"
	EventTypeViewerSubscribed        EventType = "viewer-subscribed"
	EventTypeViewerResubscribed      EventType = "viewer-resubscribed"
	EventTypeViewerReceivedGiftSub   EventType = "viewer-received-gift-sub"
	EventTypeViewerGiftedSubs        EventType = "viewer-gifted-subs"
)

// Event is an event that has occurred on Twitch, such as a viewer interaction or a
// change in the state of the stream
type Event struct {
	Type    EventType    `json:"type"`
	Viewer  *core.Viewer `json:"viewer"`
	Payload *Payload     `json:"payload"`
}

type Payload struct {
	ViewerRaided            *PayloadViewerRaided
	ViewerCheered           *PayloadViewerCheered
	ViewerRedeemedFunPoints *PayloadViewerRedeemedFunPoints
	ViewerSubscribed        *PayloadViewerSubscribed
	ViewerResubscribed      *PayloadViewerResubscribed
	ViewerReceivedGiftSub   *PayloadViewerReceivedGiftSub
	ViewerGiftedSubs        *PayloadViewerGiftedSubs
}

func (e *Event) UnmarshalJSON(data []byte) error {
	type fields struct {
		Type    EventType       `json:"type"`
		Viewer  *core.Viewer    `json:"viewer"`
		Payload json.RawMessage `json:"payload"`
	}
	var f fields
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}

	e.Type = f.Type
	e.Viewer = f.Viewer
	switch f.Type {
	case EventTypeViewerRaided:
		e.Payload = &Payload{}
		return json.Unmarshal(f.Payload, &e.Payload.ViewerRaided)
	case EventTypeViewerCheered:
		e.Payload = &Payload{}
		return json.Unmarshal(f.Payload, &e.Payload.ViewerCheered)
	case EventTypeViewerRedeemedFunPoints:
		e.Payload = &Payload{}
		return json.Unmarshal(f.Payload, &e.Payload.ViewerRedeemedFunPoints)
	case EventTypeViewerSubscribed:
		e.Payload = &Payload{}
		return json.Unmarshal(f.Payload, &e.Payload.ViewerSubscribed)
	case EventTypeViewerResubscribed:
		e.Payload = &Payload{}
		return json.Unmarshal(f.Payload, &e.Payload.ViewerResubscribed)
	case EventTypeViewerReceivedGiftSub:
		e.Payload = &Payload{}
		return json.Unmarshal(f.Payload, &e.Payload.ViewerReceivedGiftSub)
	case EventTypeViewerGiftedSubs:
		e.Payload = &Payload{}
		return json.Unmarshal(f.Payload, &e.Payload.ViewerGiftedSubs)
	}
	return nil
}

func (p Payload) MarshalJSON() ([]byte, error) {
	if p.ViewerRaided != nil {
		return json.Marshal(p.ViewerRaided)
	}
	if p.ViewerCheered != nil {
		return json.Marshal(p.ViewerCheered)
	}
	if p.ViewerRedeemedFunPoints != nil {
		return json.Marshal(p.ViewerRedeemedFunPoints)
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

type PayloadViewerRedeemedFunPoints struct {
	NumPoints int    `json:"num_points"`
	Message   string `json:"message"`
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
