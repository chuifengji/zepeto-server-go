package models

//Feature 脸型
type Feature struct {
	ID        int    `gorm:"primary_key"`
	TYPE      string `gorm:"type:varchar(20);not null;"`
	URLA      string `gorm:"type:varchar(70);not null;"`
	URLB      string `gorm:"type:varchar(70);not null;"`
	URLC      string `gorm:"type:varchar(70);not null;"`
	THUMBNAIL string `gorm:"type:varchar(70);not null;"`
}
