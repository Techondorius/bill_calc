package model

import "errors"

type User struct {
	ID   int `gorm:"primary_key"`
	Name string
}

func NewUser(name string) (*User, error) {
	if len(name) < 2 {
		return nil, errors.New("name must be more than a charactor")
	}
	return &User{Name: name}, nil
}

func InsertUserToDB(u *User) error {
	q := DB.Create(u)
	if q.Error != nil {
		return q.Error
	}
	return nil
}

func SelectAllUserFromDB() []*User {
	us := &[]*User{}
	q := DB.Find(us)
	if q.Error != nil {
		return nil
	}
	return *us
}

func SelectUserByIDFromDB(id uint) *User {
	us := &User{ID: 1}
	q := DB.Find(us)
	if q.Error != nil {
		return nil
	}
	return us
}
