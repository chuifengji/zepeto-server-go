package models

//Grouphotobab 合照
type Grouphotobab struct {
	ID       int    `gorm:"primary_key"`
	USERID   string `gorm:"type:varchar(70);not null;index"`
	LOCATION string `gorm:"type:varchar(20);not null;"`
	URL      string `gorm:"type:varchar(70);not null;"`
	TIME     string `gorm:"type:varchar(20);"`
}
