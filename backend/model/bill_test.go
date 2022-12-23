package model

import (
	"reflect"
	"testing"
)

func TestNewPrice(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    *Price
		wantErr bool
	}{
		{
			name: "border1",
			args: args{
				n: 0,
			},
			want:    nil,
			wantErr: true,
		}, {
			name: "border2",
			args: args{
				n: 1,
			},
			want: &Price{
				Price: 1,
			},
			wantErr: false,
		}, {
			name: "top1",
			args: args{
				n: MAX_PRICE,
			},
			want: &Price{
				Price: MAX_PRICE,
			},
			wantErr: false,
		}, {
			name: "top2",
			args: args{
				n: MAX_PRICE + 1,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPrice(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
