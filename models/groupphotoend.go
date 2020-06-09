package models

//Grouphotoend 合照
type Grouphotoend struct {
	ID       int    `gorm:"primary_key"`
	USERID   string `gorm:"type:varchar(70);not null;index"`
	LOCATION string `gorm:"type:varchar(20);not null;"`
	TIME     string `gorm:"type:varchar(20);"`
	URL      string `gorm:"type:varchar(70);not null;"`
}
