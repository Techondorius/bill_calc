package model

import (
	"reflect"
	"testing"
)

func TestInsertUserToDB(t *testing.T) {
	type args struct {
		u *User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				u: &User{
					ID:   0,
					Name: "1",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InsertUserToDB(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("InsertUserToDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMain(m *testing.M) {
	InitGorm()
	m.Run()

}

func TestSelectAllUserFromDB(t *testing.T) {
	tests := []struct {
		name string
		want []*User
	}{
		{
			name: "no",
			want: []*User{
				{
					ID:   0,
					Name: "1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = SelectAllUserFromDB()
		})
	}
}

func TestSelectUserByIDFromDB(t *testing.T) {
	type args struct {
		id uint
	}
	tests := []struct {
		name string
		args args
		want *User
	}{
		{
			name: "",
			args: args{
				id: 1,
			},
			want: &User{
				ID:   1,
				Name: "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectUserByIDFromDB(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
