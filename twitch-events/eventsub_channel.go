package etwitch

import (
	"encoding/json"
	"fmt"

	"github.com/golden-vcr/schemas/core"
	"github.com/nicklaw5/helix/v2"
)

func fromChannelFollowEvent(data json.RawMessage) (*Event, error) {
	var ev helix.EventSubChannelFollowEvent
	if err := json.Unmarshal(data, &ev); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ChannelFollowEvent: %w", err)
	}
	return &Event{
		Type: EventTypeViewerFollowed,
		Viewer: &core.Viewer{
			TwitchUserId:      ev.UserID,
			TwitchDisplayName: ev.UserName,
		},
	}, nil
}

func fromChannelRaidEvent(data json.RawMessage) (*Event, error) {
	var ev helix.EventSubChannelRaidEvent
	if err := json.Unmarshal(data, &ev); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ChannelRaidEvent: %w", err)
	}
	return &Event{
		Type: EventTypeViewerRaided,
		Viewer: &core.Viewer{
			TwitchUserId:      ev.FromBroadcasterUserID,
			TwitchDisplayName: ev.FromBroadcasterUserName,
		},
		Payload: &Payload{
			ViewerRaided: &PayloadViewerRaided{
				NumRaiders: ev.Viewers,
			},
		},
	}, nil
}

func fromChannelCheerEvent(data json.RawMessage) (*Event, error) {
	var ev helix.EventSubChannelCheerEvent
	if err := json.Unmarshal(data, &ev); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ChannelCheerEvent: %w", err)
	}
	var viewer *core.Viewer
	if !ev.IsAnonymous {
		viewer = &core.Viewer{
			TwitchUserId:      ev.UserID,
			TwitchDisplayName: ev.UserName,
		}
	}
	return &Event{
		Type:   EventTypeViewerCheered,
		Viewer: viewer,
		Payload: &Payload{
			ViewerCheered: &PayloadViewerCheered{
				NumBits: ev.Bits,
				Message: ev.Message,
			},
		},
	}, nil
}

func fromChannelSubscriptionEvent(data json.RawMessage) (*Event, error) {
	var ev helix.EventSubChannelSubscribeEvent
	if err := json.Unmarshal(data, &ev); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ChannelSubscribeEvent: %w", err)
	}
	creditMultiplier, err := getCreditMultiplierFromTier(ev.Tier)
	if err != nil {
		return nil, err
	}
	if ev.IsGift {
		return &Event{
			Type: EventTypeViewerReceivedGiftSub,
			Viewer: &core.Viewer{
				TwitchUserId:      ev.UserID,
				TwitchDisplayName: ev.UserName,
			},
			Payload: &Payload{
				ViewerReceivedGiftSub: &PayloadViewerReceivedGiftSub{
					CreditMultiplier: creditMultiplier,
				},
			},
		}, nil
	}
	return &Event{
		Type: EventTypeViewerSubscribed,
		Viewer: &core.Viewer{
			TwitchUserId:      ev.UserID,
			TwitchDisplayName: ev.UserName,
		},
		Payload: &Payload{
			ViewerSubscribed: &PayloadViewerSubscribed{
				CreditMultiplier: creditMultiplier,
			},
		},
	}, nil
}

func fromChannelSubscriptionMessageEvent(data json.RawMessage) (*Event, error) {
	var ev helix.EventSubChannelSubscriptionMessageEvent
	if err := json.Unmarshal(data, &ev); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ChannelSubscriptionMessageEvent: %w", err)
	}
	creditMultiplier, err := getCreditMultiplierFromTier(ev.Tier)
	if err != nil {
		return nil, err
	}
	return &Event{
		Type: EventTypeViewerResubscribed,
		Viewer: &core.Viewer{
			TwitchUserId:      ev.UserID,
			TwitchDisplayName: ev.UserName,
		},
		Payload: &Payload{
			ViewerResubscribed: &PayloadViewerResubscribed{
				CreditMultiplier:    creditMultiplier,
				NumCumulativeMonths: ev.CumulativeMonths,
				Message:             ev.Message.Text,
			},
		},
	}, nil
}

func fromChannelSubscriptionGiftEvent(data json.RawMessage) (*Event, error) {
	var ev helix.EventSubChannelSubscriptionGiftEvent
	if err := json.Unmarshal(data, &ev); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ChannelSubscriptionGiftEvent: %w", err)
	}
	creditMultiplier, err := getCreditMultiplierFromTier(ev.Tier)
	if err != nil {
		return nil, err
	}
	var viewer *core.Viewer
	if !ev.IsAnonymous {
		viewer = &core.Viewer{
			TwitchUserId:      ev.UserID,
			TwitchDisplayName: ev.UserName,
		}
	}
	return &Event{
		Type:   EventTypeViewerGiftedSubs,
		Viewer: viewer,
		Payload: &Payload{
			ViewerGiftedSubs: &PayloadViewerGiftedSubs{
				CreditMultiplier: creditMultiplier,
				NumSubscriptions: ev.Total,
			},
		},
	}, nil
}
