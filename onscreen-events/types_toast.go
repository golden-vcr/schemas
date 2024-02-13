package eonscreen

import (
	"encoding/json"

	"github.com/golden-vcr/schemas/core"
)

// PayloadToast describes an onscreen toast notification that shouts out a viewer (who
// may be anonymous, in which case Viewer is nil) in response to some interaction that
// the viewer has performed
type PayloadToast struct {
	Type   ToastType    `json:"type"`
	Viewer *core.Viewer `json:"viewer"`
	Data   *ToastData   `json:"data,omitempty"`
}

// ToastType represents the different types of toast notifications we can display
type ToastType string

const (
	ToastTypeFollowed     ToastType = "followed"
	ToastTypeCheered      ToastType = "cheered"
	ToastTypeSubscribed   ToastType = "subscribed"
	ToastTypeResubscribed ToastType = "resubscribed"
	ToastTypeGiftedSubs   ToastType = "gifted-subs"
)

// ToastData contains toast-type-specific details describing the notification we want
// to display
type ToastData struct {
	Cheered      *ToastDataCheered
	Resubscribed *ToastDataResubscribed
	GiftedSubs   *ToastDataGiftedSubs
}

func (p *PayloadToast) UnmarshalJSON(data []byte) error {
	type fields struct {
		Type   ToastType       `json:"type"`
		Viewer *core.Viewer    `json:"viewer"`
		Data   json.RawMessage `json:"data"`
	}
	var f fields
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}

	p.Type = f.Type
	p.Viewer = f.Viewer
	switch f.Type {
	case ToastTypeCheered:
		p.Data = &ToastData{}
		return json.Unmarshal(f.Data, &p.Data.Cheered)
	case ToastTypeResubscribed:
		p.Data = &ToastData{}
		return json.Unmarshal(f.Data, &p.Data.Resubscribed)
	case ToastTypeGiftedSubs:
		p.Data = &ToastData{}
		return json.Unmarshal(f.Data, &p.Data.GiftedSubs)
	}
	return nil
}

func (d ToastData) MarshalJSON() ([]byte, error) {
	if d.Cheered != nil {
		return json.Marshal(d.Cheered)
	}
	if d.Resubscribed != nil {
		return json.Marshal(d.Resubscribed)
	}
	if d.GiftedSubs != nil {
		return json.Marshal(d.GiftedSubs)
	}
	return json.Marshal(nil)
}

type ToastDataCheered struct {
	NumBits int    `json:"num_bits"`
	Message string `json:"message"`
}

type ToastDataResubscribed struct {
	NumCumulativeMonths int    `json:"num_cumulative_months"`
	Message             string `json:"message"`
}

type ToastDataGiftedSubs struct {
	NumSubscriptions int `json:"num_subscriptions"`
}
