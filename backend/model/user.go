package model

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

const (
	USER_NAME_MIN_LEN = 1
	USER_NAME_MAX_LEN = 30
	USER_ID_MIN_LEN   = 3
	USER_ID_MAX_LEN   = 15
)

type User struct {
	ID             int            `gorm:"primaryKey"`
	UserName       UserName       `gorm:"embedded"`
	UserID         UserID         `gorm:"embedded"`
	HashedPassWord HashedPassWord `gorm:"embedded"`
	Bills          []BillUsers
}

type UserName struct {
	UserName string
}

func NewUserName(un string) (*UserName, error) {
	log.Println(len(un))
	log.Println(len(un) >= USER_NAME_MIN_LEN)
	log.Println(len(un) <= USER_NAME_MAX_LEN)
	if !(len(un) >= USER_NAME_MIN_LEN && len(un) <= USER_NAME_MAX_LEN) {
		return nil, fmt.Errorf("UserNameが短すぎます。UserNameは%d~%d文字です。", USER_NAME_MIN_LEN, USER_NAME_MAX_LEN)
	}
	return &UserName{un}, nil
}

type UserID struct {
	UserID string
}

type HashedPassWord struct {
	HashedPassWord []byte
}

func NewHashedPassWord(rawPW string) HashedPassWord {
	password := []byte("password")
	hashed, _ := bcrypt.GenerateFromPassword(password, 10)
	fmt.Println((hashed))
	hpw := HashedPassWord{
		HashedPassWord: hashed,
	}
	return hpw
}

func FindUserByID(userID int) (*User, bool) {
	u := &User{ID: userID}
	if err := DB.Find(u).Error; err != nil {
		return nil, false
	}
	return u, true
}

func InsertUsers(u ...User) error {
	if err := DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}
