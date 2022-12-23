package model

type User struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Bills []BillUsers
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
