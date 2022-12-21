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
	DB = db

	u := User{}
	db.AutoMigrate(u)
	return nil
}
