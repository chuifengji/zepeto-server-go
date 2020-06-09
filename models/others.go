package models

//Others 其他
type Others struct {
	ID        int    `gorm:"primary_key"`
	TYPE      string `gorm:"type:varchar(20);not null;"`
	URL       string `gorm:"type:varchar(70);not null;"`
	THUMBNAIL string `gorm:"type:varchar(70);not null;"`
}
