package model

import (
	"reflect"
	"testing"
)

func TestInsertUserToDB(t *testing.T) {
	type args struct {
		UserName    string
		UserID      string
		RawPassword string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				UserName:    "miteh",
				UserID:      "aojiru",
				RawPassword: "asdf",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := NewUser(tt.args.UserName, tt.args.UserID, tt.args.RawPassword)
			if err != nil {
				t.Fatal(err)
			}
			if err := InsertUsers(*u); (err != nil) != tt.wantErr {
				t.Errorf("InsertUserToDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMain(m *testing.M) {
	err := InitGorm()
	if err != nil {
		panic(err)
	}
	m.Run()
}

func TestNewHashedPassWord(t *testing.T) {
	type args struct {
		rawPW string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				rawPW: "asdf",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewHashedPassWord(tt.args.rawPW)
		})
	}
}

func TestNewUserName(t *testing.T) {
	type args struct {
		un string
	}
	tests := []struct {
		name    string
		args    args
		want    *UserName
		wantErr bool
	}{
		{
			name: "空の文字列にエラーを返す",
			args: args{
				un: "",
			},
			wantErr: true,
		}, {
			name: "1文字の文字列を正常に処理する",
			args: args{
				un: "あ",
			},
			want: &UserName{
				UserName: "あ",
			},
			wantErr: false,
		}, {
			name: "30文字の文字列を正常に処理する",
			args: args{
				un: "123456789012345678901234567890",
			},
			want: &UserName{
				UserName: "123456789012345678901234567890",
			},
			wantErr: false,
		}, {
			name: "31文字の文字列にエラーを返す",
			args: args{
				un: "1234567890123456789012345678901",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserName(tt.args.un)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUserName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserName() = %v, want %v", got, tt.want)
			}
		})
	}
}
