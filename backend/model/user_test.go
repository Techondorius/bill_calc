package model

import (
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
			if err := InsertUsers(*tt.args.u); (err != nil) != tt.wantErr {
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

