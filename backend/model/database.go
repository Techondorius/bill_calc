package model

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitGorm() error {
	pw := os.Getenv("MYSQL_ROOT_PASSWORD")
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("root:%s@tcp(db:3306)/gin_app?charset=utf8&parseTime=True&loc=Local", pw)))
	if err != nil {
		return err
	}
	DB = db.Debug()

	ub := BillUsers{}
	u := User{}
	b := Bill{}

	fmt.Println(ub, u, b)

	DB.AutoMigrate(&u, &b, &ub)
	return nil
}

func InsertTestData() {
	u := []User{
		{
			Name: "Miteh",
		}, {
			Name: "Techondorius",
		},
	}
	InsertUsers(u...)
}
