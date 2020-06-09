package models

//Decoration 这golint管的还真宽
type Decoration struct {
	ID        int    `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(20);not null;"`
	URL       string `gorm:"type:varchar(70);not null;"` //背景图片连接
	THUMBNAIL string `gorm:"type:varchar(70);not null;"`
}
