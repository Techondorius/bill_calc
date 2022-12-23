package model

type BillUsers struct {
	BillID int   `gorm:"primaryKey;autoIncrement:false"`
	UserID int   `gorm:"primaryKey;autoIncrement:false"`
	Price  Price `gorm:"embedded"`
}

func NewBillUsers(userID int, price int) (*BillUsers, error) {
	pr, err := NewPrice(price)
	if err != nil {
		return nil, err
	}
	bu := &BillUsers{
		UserID: userID,
		Price:  *pr,
	}

	return bu, nil
}

func InsertBillUsers(billID int, bus ...BillUsers) error {
	if err := DB.Create(bus).Error; err != nil {
		return err
	}
	return nil
}
