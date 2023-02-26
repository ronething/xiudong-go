package showstart

import "testing"

var source = `{"action": "/wap/order/confirm", "method": "POST", "query": {"ticketId": "c87a06c071c25ffa4db03de1fe7753fa", "sequence": "163072", "ticketNum": "1", "st_flpv": "vkx9gecd6JcZH2F9k665", "sign": "7f6cfc0b3be20d820879188376b11beb", "trackPath": "", "terminal": "wap"}, "body": null, "qtime": 1638726094657, "ranstr": "qaWIAFra"}`
var dst = `vwAp12H2QLt/KTjBiio0324NlZ4nfxyE4LfSfMJwdiYAf6GducAaCR1+Lp2Pft2BKJFgcSxneNpzbs/QzldCPIJPT3OpzY089xtB5pMrz/FlkusKAct4jNDLn9l1j6f561EaC49PLzthDS2aP6nyPgjZjefa4ITWXqUI5Rt420YATZpIM8jgtji5Px67ItuVck+kYF91c6h0KHTpHOwq7HjgqVvRUFLbxblrmKDu3qkb9kHoj3BNqyNOcoHcSIUwX0Cc9QKS4Gvujztvso21z9oSJRPORyOTHGzcgdQEonEBltYJiv+8J+GFTzexKJLP70X7U6tp4Of5u1cJ8oJ1Dbp0vYxYh7MhESoT9GOffcJShmNTHlvvxXzf1gO5YWEMS3Y2oQv487zceAWxaX7XOpNWHg9qfu+zW5WOwY2any7ET3Vw2H2zVfmC/psfzIt8`

func TestAESCrypto_Decrypt(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{

		{
			name: "one",
			args: args{
				src: dst,
			},
			want:    source,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AESCrypto{
				Key: []byte("xxxx"),
			}
			got, err := a.Decrypt(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAESCrypto_Decrypt2(t *testing.T) {
	s := "" // 需要解密的字符串
	a := &AESCrypto{
		Key: []byte("xxxx"),
	}
	res, _ := a.Decrypt(s)
	t.Log(res)
}

// {"action": "/wap/order/confirm", "method": "POST", "query": {"ticketId": "c87a06c071c25ffa4db03de1fe7753fa", "sequence": "163072", "ticketNum": "1", "st_flpv": "vkx9gecd6JcZH2F9k665", "sign": "7f6cfc0b3be20d820879188376b11beb", "trackPath": "", "terminal": "wap"}, "body": null, "qtime": 1638726094657, "ranstr": "qaWIAFra"}

func TestAESCrypto_Encrypt(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "one",
			args: args{
				src: []byte(source),
			},
			want:    dst,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AESCrypto{
				Key: []byte("xxxx"),
			}
			got, err := a.Encrypt(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Encrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAESCrypto_DecryptAndMd5Sign(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name    string
		args    args
		sign    string
		wantErr bool
	}{

		{
			name: "one",
			args: args{
				src: "RdKnoWKXGxZP0zYw1ohcYmUf7kKwJJ5BrqAzR80J3Y2ScUZwWfFmi95X0Ca9/RQdX9RIgqdnicbI1kuMqLPmwazQiEzG94hg5cbsw7k2Ty84DdXyosvDt0OIQvLdqgNJn+otSArJiw3SA0i4/f+hE1TYzdm+JMXA8MsW2Am0w+mKsOIc4X7ckLF562cAhXU7/mRtIR9CMkfXXIHaVO5+Bgduc8IeiXGa+5YGpEx8uz2jlU1s0vrnMPTzHQPvhjDTFceU6037cKhEHV7G/yRxjAItY2PKi++fHrw32OWmgsQG2pmmh29nL0rQsX4Al80z",
			},
			sign:    "6fea539701a37a2406b8e08a8c941f33",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AESCrypto{
				Key: []byte("xxxx"),
			}
			got, err := a.Decrypt(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if Md5Sum(got) != tt.sign {
				t.Errorf("Decrypt() got sign = %v, want sign %v", got, tt.sign)
			}
		})
	}
}
