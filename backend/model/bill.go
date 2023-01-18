package model

import (
	"errors"
	"fmt"
	"time"
)

const MAX_PRICE = 1000000

type UIDPrice struct {
	UserID int `json:"userID"`
	Price  int `json:"price"`
}

func NewUIDPrice(useriD int, price int) *UIDPrice {
	return &UIDPrice{
		UserID: useriD,
		Price:  price,
	}
}

func (uidp UIDPrice) ToBillUsers() BillUsers {
	return BillUsers{
		UserID: uidp.UserID,
		Price: Price{
			Price: uidp.Price,
		},
	}
}

type Bill struct {
	ID           int `gorm:"primaryKey"`
	Subject      string
	Date         time.Time
	BillingUsers []*BillUsers
	Note         string
}

func FindBillByID(billID int) (*Bill, bool) {
	bill := &Bill{ID: billID}
	if err := DB.Find(bill).Error; err != nil {
		return nil, false
	}
	return bill, true
}

func NewBill(subject string, date time.Time, uIDPrice []UIDPrice, note string) (*Bill, error) {
	if len(subject) < 1 {
		return nil, errors.New("Subject must be filled")
	}

	bus := []*BillUsers{}
	for _, up := range uIDPrice {
		bu, err := NewBillUsers(up.UserID, up.Price)
		if err != nil {
			return nil, err
		}
		bus = append(bus, bu)
	}
	bill := &Bill{
		Subject:      subject,
		Date:         date,
		BillingUsers: bus,
		Note:         note,
	}
	return bill, nil
}

func (bill *Bill) AddUserToBillingUsers(uid int, price int) error {
	bu, err := NewBillUsers(uid, price)
	if err != nil {
		return err
	}
	bill.BillingUsers = append(bill.BillingUsers, bu)
	return nil
}

func (bill *Bill) SaveToDB() error {
	if DB.Create(bill).Error != nil {
		return fmt.Errorf("no rows affected...")
	}
	return nil
}

type Price struct {
	Price int
}

func NewPrice(priceInt int) (*Price, error) {
	pr := &Price{}
	if priceInt < 1 || priceInt > MAX_PRICE {
		return nil, errors.New("Price value error")
	}
	pr.Price = priceInt
	return pr, nil
}
