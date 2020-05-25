package database

import "MindaZepeto/models"

//Createtable 用于表格（Model）的初始化
func Createtable() {

	GetDB().AutoMigrate(
		&models.User{},
		&models.Friends{},
		&models.Decoration{},
		&models.Background{},
		&models.Appearance{},
	)
}
