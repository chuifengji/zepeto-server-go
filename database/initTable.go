package database

import "MindaZepeto/models"

//Createtable 用于表格（Model）的初始化
func Createtable() {

	GetDB().AutoMigrate(
		&models.User{},
		&models.Friends{},
		&models.Decoration{},
		&models.Background{},
		&models.Expression{},
		&models.Feature{},
		&models.Glasses{},
		&models.Hair{},
		&models.Shirt{},
		&models.Shoes{},
		&models.Trousers{},
		&models.Overcoat{},
		&models.Others{},
		&models.Special{},
		&models.Grouphotoaaa{},
		&models.Grouphotoaab{},
		&models.Grouphotoaac{},
		&models.Grouphotoaba{},
		&models.Grouphotoabb{},
		&models.Grouphotoabc{},
		&models.Grouphotoaca{},
		&models.Grouphotoacb{},
		&models.Grouphotoacc{},
		&models.Grouphotobaa{},
		&models.Grouphotobab{},
		&models.Grouphotobac{},
		&models.Grouphotobba{},
		&models.Grouphotoend{},
	)
}
