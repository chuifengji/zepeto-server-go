package db

import "MindaZepeto/models"

func Createtable() {
	GetDB().AutoMigrate(
		&models.User{},
	)
}
