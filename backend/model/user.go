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
	UserName string `gorm:"not null"`
}

func NewUserName(un string) (*UserName, error) {
	if !(len(un) >= USER_NAME_MIN_LEN && len(un) <= USER_NAME_MAX_LEN) {
		return nil, fmt.Errorf("UserNameの長さがおかしいです。UserNameは%d~%d文字です。", USER_NAME_MIN_LEN, USER_NAME_MAX_LEN)
	}
	return &UserName{un}, nil
}

type UserID struct {
	UserID string `gorm:"unique;not null"`
}

func NewUserID(uid string) (*UserID, error) {
	if !(len(uid) >= USER_ID_MIN_LEN && len(uid) <= USER_ID_MAX_LEN) {
		return nil, fmt.Errorf("UserIDの長さがおかしいです。UserIDは%d~%d文字です。", USER_ID_MIN_LEN, USER_ID_MAX_LEN)
	}
	return &UserID{uid}, nil
}

func (uid *UserID) CheckDuplication(uidCheck string) (bool, error) {
	checkingID, err := NewUserID(uidCheck)
	if err != nil {
		return false, err
	}
	u := &User{UserID: *checkingID}
	a := DB.Find(u).RowsAffected
	if a != 0 {
		return true, nil
	}
	return false, nil

}

type HashedPassWord struct {
	HashedPassWord []byte `gorm:"not null"`
}

func NewHashedPassWord(rawPW string) (*HashedPassWord, error) {
	if len(rawPW) == 0 {
		return nil, fmt.Errorf("pw is empty")
	}
	password := []byte(rawPW)
	hashed, _ := bcrypt.GenerateFromPassword(password, 10)
	fmt.Println((hashed))
	hpw := HashedPassWord{
		HashedPassWord: hashed,
	}
	return &hpw, nil
}

func FindUserByID(userID int) (*User, bool) {
	u := &User{ID: userID}
	if err := DB.Find(u).Error; err != nil {
		return nil, false
	}
	return u, true
}

func InsertUsers(u ...User) error {
	db := DB.Create(u)
	log.Println(db.RowsAffected)
	if err := db.Error; err != nil {
		return err
	} else {
		log.Println(err)
	}
	return nil
}

func NewUser(userName string, userID string, rawPassword string) (*User, error) {
	u := &User{}

	un, err := NewUserName(userName)
	if err != nil {
		return nil, err
	}
	u.UserName = *un

	uid, err := NewUserID(userID)
	if err != nil {
		return nil, err
	}
	u.UserID = *uid

	pw, err := NewHashedPassWord(rawPassword)
	if err != nil {
		return nil, err
	}
	u.HashedPassWord = *pw

	return u, nil
}
