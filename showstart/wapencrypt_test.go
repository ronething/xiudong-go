package showstart

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestWapEncrypt_GetRequestUrl(t *testing.T) {
	type fields struct {
		Config *WapEncryptConfig
		Source *Source
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "one",
			fields: fields{
				Config: &WapEncryptConfig{
					UserId: 4612349,
				},
				Source: &Source{
					APISource: "",
					APIId:     0,
					EventID:   0,
				},
			},
			want: "https://wap.showstart.com/api/hw/00000000jlSJ",
		},
		{
			name: "two",
			fields: fields{
				Config: &WapEncryptConfig{
					UserId: 4612349,
				},
				Source: &Source{
					APISource: "order",
					APIId:     0,
					EventID:   0,
				},
			},
			want: "https://wap.showstart.com/api/order/00000000jlSJ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WapEncrypt{
				Config: tt.fields.Config,
				Source: tt.fields.Source,
			}
			if got := w.GetRequestUrl(); got != tt.want {
				t.Errorf("GetRequestUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRandStr(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name    string
		args    args
		wantNum int
	}{
		{
			name:    "one",
			args:    args{num: 8},
			wantNum: 8,
		},
		{
			name:    "two",
			args:    args{num: 0},
			wantNum: 20,
		},
	}
	count := 0
	for {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := getRandStr(tt.args.num); len(got) != tt.wantNum {
					t.Errorf("getRandStr() = %v, want %v", got, tt.wantNum)
				}
			})
		}
		count++
		if count == 1000 { // 每个测试样例 测试 1000 次
			return
		}
	}
}

func TestWapEncrypt_getInjectData(t *testing.T) {
	type fields struct {
		Config *WapEncryptConfig
		Source *Source
	}
	type args struct {
		qtime   int64
		randStr string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "one",
			fields: fields{
				Config: &WapEncryptConfig{
					Sign:   "7f6cfc0b3be20d820879188376b11beb",
					StFlpv: "vkx9gecd6JcZH2F9k665",
					Token:  "",
					UserId: 4612349,
					AesKey: "xxxx",
				},
				Source: &Source{
					URL:    "/wap/address/list.json",
					Method: "POST",
					Data: map[string]interface{}{
						"pageNo": 1,
					},
					APISource:  "hw",
					APIId:      0,
					EventID:    0,
					ParamsType: "query",
					Headers:    map[string]string{},
				},
			},
			args: args{
				qtime:   1641026639364,
				randStr: "uB5WivU9",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WapEncrypt{
				Config: tt.fields.Config,
				Source: tt.fields.Source,
			}
			got, err := w.getInjectData(tt.args.qtime, tt.args.randStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("getInjectData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("getInjectData() got = %v", got)
		})
	}
}

func TestWapEncrypt_getInjectHeaders(t *testing.T) {
	type fields struct {
		Config *WapEncryptConfig
		Source *Source
	}
	type args struct {
		now int64
	}
	x := time.Now().Unix()
	key := "xVgXtOUSos6jzR3mqb4aLHYybqqPFFGfx12r"
	token := getRandStr(32)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]string
	}{
		{
			name: "one",
			fields: fields{
				Config: &WapEncryptConfig{
					Sign:   "7f6cfc0b3be20d820879188376b11beb",
					StFlpv: "vkx9gecd6JcZH2F9k665",
					Token:  token,
					UserId: 4612349,
					AesKey: "xxxx",
				},
				Source: &Source{
					URL:    "/wap/address/list.json",
					Method: "POST",
					Data: map[string]interface{}{
						"pageNo": 1,
					},
					APISource:  "hw",
					APIId:      0,
					EventID:    0,
					ParamsType: "query",
					Headers:    map[string]string{},
				},
			},
			args: args{now: x},
			want: map[string]string{
				"cookie":          fmt.Sprintf("Hm_lvt_da038bae565bb601b53cc9cb25cdca74=1638605081; Hm_lpvt_da038bae565bb601b53cc9cb25cdca74=%d", x),
				"content-type":    "application/json",
				"st_flpv":         "vkx9gecd6JcZH2F9k665",
				"sign":            "7f6cfc0b3be20d820879188376b11beb",
				"terminal":        "wap",
				"user-agent":      "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1",
				"sec-fetch-dest":  "empty",
				"sec-fetch-mode":  "cors",
				"sec-fetch-site":  "same-origin",
				"accept-language": "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7",
				"accept-encoding": "gzip, deflate, br",
				"origin":          "https://wap.showstart.com",
				"r":               fmt.Sprintf("%d", x*1000),
				"cusystime":       fmt.Sprintf("%d", x*1000),
				"s":               strings.ToLower(Md5Sum(fmt.Sprintf("%d%s%s%s", x*1000, "/wap/address/list.json", "vkx9gecd6JcZH2F9k665", key))),
				"cusut":           "7f6cfc0b3be20d820879188376b11beb",
				"cuuserref":       token,
				"ctrackpath":      "",
				"csourcepath":     "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WapEncrypt{
				Config: tt.fields.Config,
				Source: tt.fields.Source,
			}
			if got := w.getInjectHeaders(tt.args.now); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getInjectHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}
