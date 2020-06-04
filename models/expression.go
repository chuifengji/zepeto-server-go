package models

//Expression 表情
type Expression struct {
	ID        int    `gorm:"primary_key"`
	TYPE      string `gorm:"type:varchar(20);not null;"`
	URL       string `gorm:"type:varchar(40);not null;"`
	THUMBNAIL string `gorm:"type:varchar(40);not null;"`
}
