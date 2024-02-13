package eonscreen

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/golden-vcr/schemas/core"
	genreq "github.com/golden-vcr/schemas/generation-requests"
	"github.com/stretchr/testify/assert"
)

func Test_Event(t *testing.T) {
	tests := []struct {
		name   string
		ev     Event
		jsonEv string
	}{
		{
			"status changed: tape 50 is now being screened",
			Event{
				Type: EventTypeStatus,
				Payload: Payload{
					Status: &PayloadStatus{
						CurrentTapeId: 50,
					},
				},
			},
			`{"type":"status","payload":{"current_tape_id":50}}`,
		},
		{
			"onscreen toast for a user that just followed",
			Event{
				Type: EventTypeToast,
				Payload: Payload{
					Toast: &PayloadToast{
						Type: ToastTypeFollowed,
						Viewer: &core.Viewer{
							TwitchUserId:      "90790024",
							TwitchDisplayName: "wasabimilkshake",
						},
					},
				},
			},
			`{"type":"toast","payload":{"type":"followed","viewer":{"twitch_user_id":"90790024","twitch_display_name":"wasabimilkshake"}}}`,
		},
		{
			"onscreen toast for a user that just raided",
			Event{
				Type: EventTypeToast,
				Payload: Payload{
					Toast: &PayloadToast{
						Type: ToastTypeRaided,
						Viewer: &core.Viewer{
							TwitchUserId:      "90790024",
							TwitchDisplayName: "wasabimilkshake",
						},
						Data: &ToastData{
							Raided: &ToastDataRaided{
								NumViewers: 41,
							},
						},
					},
				},
			},
			`{"type":"toast","payload":{"type":"raided","viewer":{"twitch_user_id":"90790024","twitch_display_name":"wasabimilkshake"},"data":{"num_viewers":41}}}`,
		},
		{
			"onscreen toast for a user that just cheered",
			Event{
				Type: EventTypeToast,
				Payload: Payload{
					Toast: &PayloadToast{
						Type: ToastTypeCheered,
						Viewer: &core.Viewer{
							TwitchUserId:      "90790024",
							TwitchDisplayName: "wasabimilkshake",
						},
						Data: &ToastData{
							Cheered: &ToastDataCheered{
								NumBits: 200,
								Message: "hello world",
							},
						},
					},
				},
			},
			`{"type":"toast","payload":{"type":"cheered","viewer":{"twitch_user_id":"90790024","twitch_display_name":"wasabimilkshake"},"data":{"num_bits":200,"message":"hello world"}}}`,
		},
		{
			"onscreen toast for a user that just cheered anonymously",
			Event{
				Type: EventTypeToast,
				Payload: Payload{
					Toast: &PayloadToast{
						Type: ToastTypeCheered,
						Data: &ToastData{
							Cheered: &ToastDataCheered{
								NumBits: 200,
								Message: "hello world",
							},
						},
					},
				},
			},
			`{"type":"toast","payload":{"type":"cheered","viewer":null,"data":{"num_bits":200,"message":"hello world"}}}`,
		},
		{
			"onscreen toast for a user that just subscribed",
			Event{
				Type: EventTypeToast,
				Payload: Payload{
					Toast: &PayloadToast{
						Type: ToastTypeSubscribed,
						Viewer: &core.Viewer{
							TwitchUserId:      "90790024",
							TwitchDisplayName: "wasabimilkshake",
						},
					},
				},
			},
			`{"type":"toast","payload":{"type":"subscribed","viewer":{"twitch_user_id":"90790024","twitch_display_name":"wasabimilkshake"}}}`,
		},
		{
			"onscreen toast for a user that just resubscribed",
			Event{
				Type: EventTypeToast,
				Payload: Payload{
					Toast: &PayloadToast{
						Type: ToastTypeResubscribed,
						Viewer: &core.Viewer{
							TwitchUserId:      "90790024",
							TwitchDisplayName: "wasabimilkshake",
						},
						Data: &ToastData{
							Resubscribed: &ToastDataResubscribed{
								NumCumulativeMonths: 3,
								Message:             "good job",
							},
						},
					},
				},
			},
			`{"type":"toast","payload":{"type":"resubscribed","viewer":{"twitch_user_id":"90790024","twitch_display_name":"wasabimilkshake"},"data":{"num_cumulative_months":3,"message":"good job"}}}`,
		},
		{
			"onscreen toast for a user that just gifted subs",
			Event{
				Type: EventTypeToast,
				Payload: Payload{
					Toast: &PayloadToast{
						Type: ToastTypeGiftedSubs,
						Viewer: &core.Viewer{
							TwitchUserId:      "90790024",
							TwitchDisplayName: "wasabimilkshake",
						},
						Data: &ToastData{
							GiftedSubs: &ToastDataGiftedSubs{
								NumSubscriptions: 5,
							},
						},
					},
				},
			},
			`{"type":"toast","payload":{"type":"gifted-subs","viewer":{"twitch_user_id":"90790024","twitch_display_name":"wasabimilkshake"},"data":{"num_subscriptions":5}}}`,
		},
		{
			"onscreen toast for a user that just gifted subs anonymously",
			Event{
				Type: EventTypeToast,
				Payload: Payload{
					Toast: &PayloadToast{
						Type: ToastTypeGiftedSubs,
						Data: &ToastData{
							GiftedSubs: &ToastDataGiftedSubs{
								NumSubscriptions: 5,
							},
						},
					},
				},
			},
			`{"type":"toast","payload":{"type":"gifted-subs","viewer":null,"data":{"num_subscriptions":5}}}`,
		},
		{
			"playback of a ghost image alert",
			Event{
				Type: EventTypeImage,
				Payload: Payload{
					Image: &PayloadImage{
						Viewer: core.Viewer{
							TwitchUserId:      "90790024",
							TwitchDisplayName: "wasabimilkshake",
						},
						Style:       genreq.ImageStyleGhost,
						Description: "a seal",
						ImageUrl:    "https://my-cool-images.biz/seal.jpg",
					},
				},
			},
			`{"type":"image","payload":{"viewer":{"twitch_user_id":"90790024","twitch_display_name":"wasabimilkshake"},"style":"ghost","description":"a seal","image_url":"https://my-cool-images.biz/seal.jpg"}}`,
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
