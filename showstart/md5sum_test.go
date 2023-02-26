package showstart

import "testing"

func TestMd5Sum(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one",
			args: args{
				value: "hello world",
			},
			want: "5eb63bbbe01eeed093cb22bb8f5acdc3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5Sum(tt.args.value); got != tt.want {
				t.Errorf("Md5Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMd5SumByte(t *testing.T) {
	type args struct {
		value []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one",
			args: args{
				value: []byte("hello world"),
			},
			want: "5eb63bbbe01eeed093cb22bb8f5acdc3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5SumByte(tt.args.value); got != tt.want {
				t.Errorf("Md5SumByte() = %v, want %v", got, tt.want)
			}
		})
	}
}
