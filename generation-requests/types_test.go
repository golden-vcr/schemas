package genreq

import (
	"encoding/json"
	"fmt"
	"testing"

	etwitch "github.com/golden-vcr/schemas/twitch-events"
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
				Viewer: etwitch.Viewer{
					TwitchUserId:      "90790024",
					TwitchDisplayName: "wasabimilkshake",
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
			`{"type":"image","viewer":{"twitch_user_id":"90790024","twitch_display_name":"wasabimilkshake"},"payload":{"style":"ghost","inputs":{"subject":"a seal"}}}`,
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
