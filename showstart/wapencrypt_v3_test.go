package showstart

import (
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

func TestWapEncryptV3(t *testing.T) {
	type fields struct {
		Config *WapEncryptConfigV3
		Source *SourceV3
	}
	type args struct {
		now int64
	}

	x := time.Now().Unix()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "one",
			fields: fields{
				Config: &WapEncryptConfigV3{
					Sign:        "05cedc227d1a0b181174d92e50d871ea",
					StFlpv:      "nJmxO93zd8Qt7aWor4W3",
					Token:       "c559aab27834dfa16c939420794b038e",
					UserId:      4613349,
					AccessToken: "ad6c5049021801fa64cdd17ed38ca5e4",
					IdToken:     "1bcda8b712def05264cdd17e1d21f21a",
				},
				Source: &SourceV3{
					URL:    "/wap/activity/V2/ticket/list",
					Method: "POST",
					Data: map[string]interface{}{
						"activityId": "204504",
						"coupon":     "",
					},
					Headers: map[string]string{},
					Secret:  false,
				},
			},
			args: args{
				now: x,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s3, err := NewWapEncryptV3(tt.fields.Config, tt.fields.Source)
			if err != nil {
				t.Error(err)
				return
			}
			headers := s3.Source.Headers
			for k, v := range headers {
				t.Logf("key: %s, value: %s", k, v)
			}
			url := s3.GetRequestUrl()
			assert.Equal(t, url, "https://wap.showstart.com/v3/wap/activity/V2/ticket/list")
		})
	}
}
