package showstart

import (
	"log"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

func getTestShowStart() *ShowStart {
	// 读取配置文件 默认在 ../.showstart.yaml
	viper.SetConfigFile("../.showstart.yaml")
	if err := viper.ReadInConfig(); err == nil {
		config := &WapEncryptConfig{
			Sign:   viper.GetString("sign"),
			StFlpv: viper.GetString("st_flpv"),
			Token:  viper.GetString("token"),
			UserId: viper.GetUint32("userId"),
			AesKey: viper.GetString("aesKey"),
		}
		log.Printf("config is %+v\n", config)
		return &ShowStart{
			Cookies: config,
			Client:  resty.New(),
		}
	} else {
		log.Fatalln("请确保配置文件路径正确")
		return nil
	}
}

func TestShowStart_GetAddressList(t *testing.T) {
	type args struct {
		pageNo int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "one",
			args:    args{pageNo: 1},
			wantErr: false,
		},
		{
			name:    "two",
			args:    args{pageNo: 2},
			wantErr: false,
		},
	}
	s := getTestShowStart()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetAddressList(tt.args.pageNo)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAddressList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got is %+v\n", got)
		})
	}
}

func TestShowStart_GetAddress(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "one",
			wantErr: false,
		},
	}
	s := getTestShowStart()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.GetAddress()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestShowStart_GetCpList1(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "one",
			wantErr: false,
		},
	}
	s := getTestShowStart()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetCpList(1)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCpList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got is %+v\n", got)
		})
	}
}

func TestShowStart_GetTicketList(t *testing.T) {
	type args struct {
		activityId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "one",
			args: args{
				activityId: "168949",
			},
			wantErr: false,
		},
	}
	s := getTestShowStart()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetTicketList(tt.args.activityId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTicketList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, item := range got {
				t.Logf("sessionName: %+v\n", item.SessionName)
				for _, ticket := range item.TicketList {
					t.Logf(
						"ticketId: %+v, ticketType: %+v, sellingPrice: %+v\n",
						ticket.TicketId,
						ticket.TicketType,
						ticket.SellingPrice,
					)
				}
			}
		})
	}
}
