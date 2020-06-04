package models

//Glasses 眼镜
type Glasses struct {
	ID        int    `gorm:"primary_key"`
	TYPE      string `gorm:"type:varchar(20);not null;"`
	URL       string `gorm:"type:varchar(40);not null;"`
	THUMBNAIL string `gorm:"type:varchar(40);not null;"`
}
