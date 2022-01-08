package db

import (
	entity "Go_app/Models/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func open() {
	dsn := "root:nanako0303@tcp(localhost:3306)/Go_app?charset=utf8&parseTime=True&loc=Local"
	v2_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	// Usersというテーブルを作ります
	v2_db.Migrator().CreateTable(&entity.User{})
}

func GormConnect() *gorm.DB {
	dsn := "root:nanako0303@tcp(localhost:3306)/Go_app?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
