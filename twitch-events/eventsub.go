package etwitch

import (
	"encoding/json"
	"errors"

	"github.com/nicklaw5/helix/v2"
)

var ErrUnsupportedEventSubType = errors.New("unsupported EventSub type")

func FromEventSub(subscription *helix.EventSubSubscription, data json.RawMessage) (*Event, error) {
	switch subscription.Type {
	case helix.EventSubTypeStreamOnline:
		return fromStreamOnlineEvent(data)
	case helix.EventSubTypeStreamOffline:
		return fromStreamOfflineEvent(data)
	case helix.EventSubTypeChannelFollow:
		return fromChannelFollowEvent(data)
	case helix.EventSubTypeChannelRaid:
		return fromChannelRaidEvent(data)
	case helix.EventSubTypeChannelCheer:
		return fromChannelCheerEvent(data)
	case helix.EventSubTypeChannelSubscription:
		return fromChannelSubscriptionEvent(data)
	case helix.EventSubTypeChannelSubscriptionMessage:
		return fromChannelSubscriptionMessageEvent(data)
	case helix.EventSubTypeChannelSubscriptionGift:
		return fromChannelSubscriptionGiftEvent(data)
	default:
		return nil, ErrUnsupportedEventSubType
	}
}
