package models

//Decoration 这golint管的还真宽
type Decoration struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"type:varchar(20);not null;"`
	Age  string `gorm:"type:varchar(10);not null;"`
	Sex  string `gorm:"type:varchar(20);not null;"`
}
