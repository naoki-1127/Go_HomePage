package db

import (
	dbconnect "Go_app/Config"
	entity "Go_app/Models/entity"
)

func Open() {
	v2_db := dbconnect.GormConnect()

	// Usersというテーブルを作ります
	v2_db.Migrator().CreateTable(&entity.User{})
	// Contactsというテーブルを作ります
	v2_db.Migrator().CreateTable(&entity.Contact{})
}
