package showstart

import "testing"

func Test_checkState(t *testing.T) {
	type args struct {
		state interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "one",
			args: args{
				state: float64(1),
			},
			wantErr: false,
		},
		{
			name: "two",
			args: args{
				state: "1",
			},
			wantErr: false,
		},
		{
			name: "err",
			args: args{
				state: "user not login",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkState(tt.args.state); (err != nil) != tt.wantErr {
				t.Errorf("checkState() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
