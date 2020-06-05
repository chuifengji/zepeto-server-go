package models

//Grouphotoacb 合照
type Grouphotoacb struct {
	ID        int    `gorm:"primary_key"`
	USERID    string `gorm:"type:varchar(70);not null;index"`
	LOCATION  string `gorm:"type:varchar(20);not null;"`
	URL       string `gorm:"type:varchar(40);not null;"`
	THUMBNAIL string `gorm:"type:varchar(40);not null;"`
}
