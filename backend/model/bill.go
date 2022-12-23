package model

import "errors"

var MAX_PRICE = 1000000

type uIDPrice struct {
	userID int
	price  int
}

func NewUIDPrice(useriD int, price int) *uIDPrice {
	return &uIDPrice{
		userID: useriD,
		price:  price,
	}
}

type Bill struct {
	ID           int `gorm:"primaryKey"`
	Subject      string
	Price        Price `gorm:"embedded"`
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

func NewBill(subject string, price int, uIDPrice []uIDPrice, note string) (*Bill, error) {
	if len(subject) < 1 {
		return nil, errors.New("Subject must be filled")
	}
	pr, err := NewPrice(price)
	if err != nil {
		return nil, err
	}
	bus := []*BillUsers{}
	for _, up := range uIDPrice {
		bu, err := NewBillUsers(up.userID, up.price)
		if err != nil {
			return nil, err
		}
		bus = append(bus, bu)
	}
	bill := &Bill{
		Subject:      subject,
		Price:        *pr,
		BillingUsers: bus,
		Note:         note,
	}
	return bill, nil
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
