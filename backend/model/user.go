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

type UserID struct {
	UserID string `gorm:"unique;not null"`
}

type HashedPassWord struct {
	HashedPassWord []byte `gorm:"not null"`
}

func NewUserName(un string) (*UserName, error) {
	if !(len(un) >= USER_NAME_MIN_LEN && len(un) <= USER_NAME_MAX_LEN) {
		return nil, fmt.Errorf("UserNameの長さがおかしいです。UserNameは%d~%d文字です。", USER_NAME_MIN_LEN, USER_NAME_MAX_LEN)
	}
	return &UserName{un}, nil
}

func NewUserID(uid string) (*UserID, error) {
	if !(len(uid) >= USER_ID_MIN_LEN && len(uid) <= USER_ID_MAX_LEN) {
		return nil, fmt.Errorf("UserIDの長さがおかしいです。UserIDは%d~%d文字です。", USER_ID_MIN_LEN, USER_ID_MAX_LEN)
	}
	return &UserID{uid}, nil
}

func NewHashedPassWord(rawPW string) (*HashedPassWord, error) {
	if len(rawPW) == 0 {
		return nil, fmt.Errorf("pw is empty")
	}
	password := []byte(rawPW)
	hashed, _ := bcrypt.GenerateFromPassword(password, 10)
	hpw := HashedPassWord{
		HashedPassWord: hashed,
	}
	return &hpw, nil
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

func (u *User) CompairHashPWandRow(rawPW string) error {
	if err := bcrypt.CompareHashAndPassword(u.HashedPassWord.HashedPassWord, []byte(rawPW)); err != nil {
		return err
	}
	return nil
}

func (uid *UserID) CheckIDDuplication(uidCheck string) error {
	checkingID, err := NewUserID(uidCheck)
	if err != nil {
		return err
	}
	u := &User{UserID: *checkingID}
	a := DB.Find(u).RowsAffected
	if a != 0 {
		return fmt.Errorf("ID already Taken")
	}
	return nil
}

func (uid *UserID) ReturnString() string {
	return uid.UserID
}

func FindUserByID(userID string) (*User, error) {
	u := &User{}
	if err := DB.First(u, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return u, nil
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
