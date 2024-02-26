package etwitch

import (
	"encoding/json"
	"testing"

	"github.com/nicklaw5/helix/v2"
	"github.com/stretchr/testify/assert"
)

func Test_FromEventSub(t *testing.T) {
	// Test EventSub payloads are copied directly from:
	// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types
	tests := []struct {
		subscriptionType string
		variant          string
		inputEvent       string
		wantErr          error
		want             string
	}{
		{
			"stream.online",
			"",
			`{
				"id": "9001",
				"broadcaster_user_id": "1337",
				"broadcaster_user_login": "cool_user",
				"broadcaster_user_name": "Cool_User",
				"type": "live",
				"started_at": "2020-10-11T10:11:12.123Z"
			}`,
			nil,
			`{
				"type": "stream-started",
				"viewer": null,
				"payload": null
			}`,
		},
		{
			"stream.offline",
			"",
			`{
				"broadcaster_user_id": "1337",
				"broadcaster_user_login": "cool_user",
				"broadcaster_user_name": "Cool_User"
			}`,
			nil,
			`{
				"type": "stream-ended",
				"viewer": null,
				"payload": null
			}`,
		},
		{
			"channel.hype_train.begin",
			"",
			`{
				"id": "1b0AsbInCHZW2SQFQkCzqN07Ib2",
				"broadcaster_user_id": "1337",
				"broadcaster_user_login": "cool_user",
				"broadcaster_user_name": "Cool_User",
				"total": 137,
				"progress": 137,
				"goal": 500,
				"top_contributions": [
					{ "user_id": "123", "user_login": "pogchamp", "user_name": "PogChamp", "type": "bits", "total": 50 },
					{ "user_id": "456", "user_login": "kappa", "user_name": "Kappa", "type": "subscription", "total": 45 }
				],
				"last_contribution": { "user_id": "123", "user_login": "pogchamp", "user_name": "PogChamp", "type": "bits", "total": 50 },
				"level": 2,
				"started_at": "2020-07-15T17:16:03.17106713Z",
				"expires_at": "2020-07-15T17:16:11.17106713Z"
			}`,
			nil,
			`{
				"type": "stream-hype-started",
				"viewer": null,
				"payload": null
			}`,
		},
		{
			"channel.follow",
			"",
			`{
				"user_id": "1234",
				"user_login": "cool_user",
				"user_name": "Cool_User",
				"broadcaster_user_id": "1337",
				"broadcaster_user_login": "cooler_user",
				"broadcaster_user_name": "Cooler_User",
				"followed_at": "2020-07-15T18:16:11.17106713Z"
			}`,
			nil,
			`{
				"type": "viewer-followed",
				"viewer": {
					"twitch_user_id": "1234",
					"twitch_display_name": "Cool_User"
				},
				"payload": null
			}`,
		},
		{
			"channel.raid",
			"",
			`{
				"from_broadcaster_user_id": "1234",
				"from_broadcaster_user_login": "cool_user",
				"from_broadcaster_user_name": "Cool_User",
				"to_broadcaster_user_id": "1337",
				"to_broadcaster_user_login": "cooler_user",
				"to_broadcaster_user_name": "Cooler_User",
				"viewers": 9001
			}`,
			nil,
			`{
				"type": "viewer-raided",
				"viewer": {
					"twitch_user_id": "1234",
					"twitch_display_name": "Cool_User"
				},
				"payload": {
					"num_raiders": 9001
				}
			}`,
		},
		{
			"channel.cheer",
			"",
			`{
				"is_anonymous": false,
				"user_id": "1234",
				"user_login": "cool_user",
				"user_name": "Cool_User",
				"broadcaster_user_id": "1337",
				"broadcaster_user_login": "cooler_user",
				"broadcaster_user_name": "Cooler_User",
				"message": "pogchamp",
				"bits": 1000
			}`,
			nil,
			`{
				"type": "viewer-cheered",
				"viewer": {
					"twitch_user_id": "1234",
					"twitch_display_name": "Cool_User"
				},
				"payload": {
					"num_bits": 1000,
					"message": "pogchamp"
				}
			}`,
		},
		{
			"channel.cheer",
			"anonymous, null viewer",
			`{
				"is_anonymous": true,
				"user_id": null,
				"user_login": null,
				"user_name": null,
				"broadcaster_user_id": "1337",
				"broadcaster_user_login": "cooler_user",
				"broadcaster_user_name": "Cooler_User",
				"message": "pogchamp",
				"bits": 1000
			}`,
			nil,
			`{
				"type": "viewer-cheered",
				"viewer": null,
				"payload": {
					"num_bits": 1000,
					"message": "pogchamp"
				}
			}`,
		},
		{
			"channel.subscribe",
			"",
			`{
				"user_id": "1234",
				"user_login": "cool_user",
				"user_name": "Cool_User",
				"broadcaster_user_id": "1337",
				"broadcaster_user_login": "cooler_user",
				"broadcaster_user_name": "Cooler_User",
				"tier": "1000",
				"is_gift": false
			}`,
			nil,
			`{
				"type": "viewer-subscribed",
				"viewer": {
					"twitch_user_id": "1234",
					"twitch_display_name": "Cool_User"
				},
				"payload": {
					"credit_multiplier": 1
				}
			}`,
		},
		{
			"channel.subscribe",
			"tier 2",
			`{
				"user_id": "1234",
				"user_login": "cool_user",
				"user_name": "Cool_User",
				"broadcaster_user_id": "1337",
				"broadcaster_user_login": "cooler_user",
				"broadcaster_user_name": "Cooler_User",
				"tier": "2000",
				"is_gift": false
			}`,
			nil,
			`{
				"type": "viewer-subscribed",
				"viewer": {
					"twitch_user_id": "1234",
					"twitch_display_name": "Cool_User"
				},
				"payload": {
					"credit_multiplier": 2
				}
			}`,
		},
		{
			"channel.subscribe",
			"gift",
			`{
				"user_id": "1234",
				"user_login": "cool_user",
				"user_name": "Cool_User",
				"broadcaster_user_id": "1337",
				"broadcaster_user_login": "cooler_user",
				"broadcaster_user_name": "Cooler_User",
				"tier": "1000",
				"is_gift": true
			}`,
			nil,
			`{
				"type": "viewer-received-gift-sub",
				"viewer": {
					"twitch_user_id": "1234",
					"twitch_display_name": "Cool_User"
				},
				"payload": {
					"credit_multiplier": 1
				}
			}`,
		},
		{
			"channel.subscribe",
			"gift, tier 3",
			`{
				"user_id": "1234",
				"user_login": "cool_user",
				"user_name": "Cool_User",
				"broadcaster_user_id": "1337",
				"broadcaster_user_login": "cooler_user",
				"broadcaster_user_name": "Cooler_User",
				"tier": "3000",
				"is_gift": true
			}`,
			nil,
			`{
				"type": "viewer-received-gift-sub",
				"viewer": {
					"twitch_user_id": "1234",
					"twitch_display_name": "Cool_User"
				},
				"payload": {
					"credit_multiplier": 5
				}
			}`,
		},
		{
			"channel.subscription.message",
			"",
			`{
				"user_id": "1234",
				"user_login": "cool_user",
				"user_name": "Cool_User",
				"broadcaster_user_id": "1337",
				"broadcaster_user_login": "cooler_user",
				"broadcaster_user_name": "Cooler_User",
				"tier": "2000",
				"message": {
					"text": "Love the stream! FevziGG",
					"emotes": [
						{
							"begin": 23,
							"end": 30,
							"id": "302976485"
						}
					]
				},
				"cumulative_months": 15,
				"streak_months": 1,
				"duration_months": 6
			}`,
			nil,
			`{
				"type": "viewer-resubscribed",
				"viewer": {
					"twitch_user_id": "1234",
					"twitch_display_name": "Cool_User"
				},
				"payload": {
					"credit_multiplier": 2,
					"num_cumulative_months": 15,
					"message": "Love the stream! FevziGG"
				}
			}`,
		},
		{
			"channel.subscription.gift",
			"",
			`{
				"user_id": "1234",
				"user_login": "cool_user",
				"user_name": "Cool_User",
				"broadcaster_user_id": "1337",
				"broadcaster_user_login": "cooler_user",
				"broadcaster_user_name": "Cooler_User",
				"total": 2,
				"tier": "3000",
				"cumulative_total": 284,
				"is_anonymous": false
			}`,
			nil,
			`{
				"type": "viewer-gifted-subs",
				"viewer": {
					"twitch_user_id": "1234",
					"twitch_display_name": "Cool_User"
				},
				"payload": {
					"credit_multiplier": 5,
					"num_subscriptions": 2
				}
			}`,
		},
		{
			"channel.subscription.gift",
			"anonymous, null viewer",
			`{
				"user_id": null,
				"user_login": null,
				"user_name": null,
				"broadcaster_user_id": "1337",
				"broadcaster_user_login": "cooler_user",
				"broadcaster_user_name": "Cooler_User",
				"total": 2,
				"tier": "1000",
				"cumulative_total": null,
				"is_anonymous": true
			}`,
			nil,
			`{
				"type": "viewer-gifted-subs",
				"viewer": null,
				"payload": {
					"credit_multiplier": 1,
					"num_subscriptions": 2
				}
			}`,
		},
	}
	for _, tt := range tests {
		name := tt.subscriptionType
		if tt.variant != "" {
			name += " " + tt.variant
		}
		t.Run(name, func(t *testing.T) {
			subscription := &helix.EventSubSubscription{Type: tt.subscriptionType}
			data := json.RawMessage(tt.inputEvent)
			gotEvent, err := FromEventSub(subscription, data)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, gotEvent)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, gotEvent)
				got, err := json.MarshalIndent(gotEvent, "\t\t\t", "\t")
				assert.NoError(t, err)
				assert.Equal(t, tt.want, string(got))
			}
		})
	}
}
