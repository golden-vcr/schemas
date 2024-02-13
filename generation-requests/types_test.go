package genreq

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/golden-vcr/schemas/core"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_Request(t *testing.T) {
	tests := []struct {
		name    string
		req     Request
		jsonReq string
	}{
		{
			"request for a ghost image",
			Request{
				Type: RequestTypeImage,
				Viewer: core.Viewer{
					TwitchUserId:      "90790024",
					TwitchDisplayName: "wasabimilkshake",
				},
				State: core.State{
					BroadcastId: 13,
					ScreeningId: uuid.MustParse("96d1ca5c-7658-48c9-8193-9d1739854467"),
					TapeId:      124,
				},
				Payload: Payload{
					Image: &PayloadImage{
						Style: ImageStyleGhost,
						Inputs: ImageInputs{
							Ghost: &ImageInputsGhost{
								Subject: "a seal",
							},
						},
					},
				},
			},
			`{"type":"image","viewer":{"twitch_user_id":"90790024","twitch_display_name":"wasabimilkshake"},"state":{"broadcast_id":13,"screening_id":"96d1ca5c-7658-48c9-8193-9d1739854467","tape_id":124},"payload":{"style":"ghost","inputs":{"subject":"a seal"}}}`,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("marshal %s to JSON", tt.name), func(t *testing.T) {
			want := tt.jsonReq
			got, err := json.Marshal(tt.req)
			assert.NoError(t, err)
			assert.Equal(t, want, string(got))
		})
		t.Run(fmt.Sprintf("unmarshal %s from JSON", tt.name), func(t *testing.T) {
			want := tt.req
			var got Request
			err := json.Unmarshal([]byte(tt.jsonReq), &got)
			assert.NoError(t, err)
			assert.Equal(t, want, got)
		})
	}
}
