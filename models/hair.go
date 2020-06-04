package models

//Hair 发型
type Hair struct {
	ID        int    `gorm:"primary_key"`
	TYPE      string `gorm:"type:varchar(20);not null;"`
	URLA      string `gorm:"type:varchar(40);not null;"`
	URLB      string `gorm:"type:varchar(40);not null;"`
	THUMBNAIL string `gorm:"type:varchar(40);not null;"`
}
