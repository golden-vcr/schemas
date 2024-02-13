package etwitch

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/golden-vcr/schemas/core"
	"github.com/stretchr/testify/assert"
)

func Test_Event(t *testing.T) {
	tests := []struct {
		name   string
		ev     Event
		jsonEv string
	}{
		{
			"stream started event",
			Event{
				Type: EventTypeStreamStarted,
			},
			`{"type":"stream-started","viewer":null,"payload":null}`,
		},
		{
			"stream ended event",
			Event{
				Type: EventTypeStreamEnded,
			},
			`{"type":"stream-ended","viewer":null,"payload":null}`,
		},
		{
			"viewer followed event",
			Event{
				Type: EventTypeViewerFollowed,
				Viewer: &core.Viewer{
					TwitchUserId:      "90790024",
					TwitchDisplayName: "wasabimilkshake",
				},
			},
			`{"type":"viewer-followed","viewer":{"twitch_user_id":"90790024","twitch_display_name":"wasabimilkshake"},"payload":null}`,
		},
		{
			"viewer raided event",
			Event{
				Type: EventTypeViewerRaided,
				Viewer: &core.Viewer{
					TwitchUserId:      "90790024",
					TwitchDisplayName: "wasabimilkshake",
				},
				Payload: &Payload{
					ViewerRaided: &PayloadViewerRaided{
						NumRaiders: 42,
					},
				},
			},
			`{"type":"viewer-raided","viewer":{"twitch_user_id":"90790024","twitch_display_name":"wasabimilkshake"},"payload":{"num_raiders":42}}`,
		},
		{
			"viewer cheered event",
			Event{
				Type: EventTypeViewerCheered,
				Viewer: &core.Viewer{
					TwitchUserId:      "90790024",
					TwitchDisplayName: "wasabimilkshake",
				},
				Payload: &Payload{
					ViewerCheered: &PayloadViewerCheered{
						NumBits: 200,
						Message: "ghost of a seal",
					},
				},
			},
			`{"type":"viewer-cheered","viewer":{"twitch_user_id":"90790024","twitch_display_name":"wasabimilkshake"},"payload":{"num_bits":200,"message":"ghost of a seal"}}`,
		},
		{
			"viewer cheered event (anonymous)",
			Event{
				Type: EventTypeViewerCheered,
				Payload: &Payload{
					ViewerCheered: &PayloadViewerCheered{
						NumBits: 200,
						Message: "ghost of a seal",
					},
				},
			},
			`{"type":"viewer-cheered","viewer":null,"payload":{"num_bits":200,"message":"ghost of a seal"}}`,
		},
		{
			"viewer redeemed fun points event",
			Event{
				Type: EventTypeViewerRedeemedFunPoints,
				Viewer: &core.Viewer{
					TwitchUserId:      "90790024",
					TwitchDisplayName: "wasabimilkshake",
				},
				Payload: &Payload{
					ViewerRedeemedFunPoints: &PayloadViewerRedeemedFunPoints{
						NumPoints: 200,
						Message:   "ghost of a seal",
					},
				},
			},
			`{"type":"viewer-redeemed-fun-points","viewer":{"twitch_user_id":"90790024","twitch_display_name":"wasabimilkshake"},"payload":{"num_points":200,"message":"ghost of a seal"}}`,
		},
		{
			"viewer subscribed event",
			Event{
				Type: EventTypeViewerSubscribed,
				Viewer: &core.Viewer{
					TwitchUserId:      "90790024",
					TwitchDisplayName: "wasabimilkshake",
				},
				Payload: &Payload{
					ViewerSubscribed: &PayloadViewerSubscribed{
						CreditMultiplier: 1,
					},
				},
			},
			`{"type":"viewer-subscribed","viewer":{"twitch_user_id":"90790024","twitch_display_name":"wasabimilkshake"},"payload":{"credit_multiplier":1}}`,
		},
		{
			"viewer resubscribed event",
			Event{
				Type: EventTypeViewerResubscribed,
				Viewer: &core.Viewer{
					TwitchUserId:      "90790024",
					TwitchDisplayName: "wasabimilkshake",
				},
				Payload: &Payload{
					ViewerResubscribed: &PayloadViewerResubscribed{
						CreditMultiplier:    1,
						NumCumulativeMonths: 3,
						Message:             "good job",
					},
				},
			},
			`{"type":"viewer-resubscribed","viewer":{"twitch_user_id":"90790024","twitch_display_name":"wasabimilkshake"},"payload":{"credit_multiplier":1,"num_cumulative_months":3,"message":"good job"}}`,
		},
		{
			"viewer received gift sub event",
			Event{
				Type: EventTypeViewerReceivedGiftSub,
				Viewer: &core.Viewer{
					TwitchUserId:      "90790024",
					TwitchDisplayName: "wasabimilkshake",
				},
				Payload: &Payload{
					ViewerReceivedGiftSub: &PayloadViewerReceivedGiftSub{
						CreditMultiplier: 1,
					},
				},
			},
			`{"type":"viewer-received-gift-sub","viewer":{"twitch_user_id":"90790024","twitch_display_name":"wasabimilkshake"},"payload":{"credit_multiplier":1}}`,
		},
		{
			"viewer gifted subs event",
			Event{
				Type: EventTypeViewerGiftedSubs,
				Viewer: &core.Viewer{
					TwitchUserId:      "90790024",
					TwitchDisplayName: "wasabimilkshake",
				},
				Payload: &Payload{
					ViewerGiftedSubs: &PayloadViewerGiftedSubs{
						CreditMultiplier: 1,
						NumSubscriptions: 5,
					},
				},
			},
			`{"type":"viewer-gifted-subs","viewer":{"twitch_user_id":"90790024","twitch_display_name":"wasabimilkshake"},"payload":{"credit_multiplier":1,"num_subscriptions":5}}`,
		},
		{
			"viewer gifted subs event (anonymous)",
			Event{
				Type: EventTypeViewerGiftedSubs,
				Payload: &Payload{
					ViewerGiftedSubs: &PayloadViewerGiftedSubs{
						CreditMultiplier: 1,
						NumSubscriptions: 5,
					},
				},
			},
			`{"type":"viewer-gifted-subs","viewer":null,"payload":{"credit_multiplier":1,"num_subscriptions":5}}`,
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
