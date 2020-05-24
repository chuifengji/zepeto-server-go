package models

//User 这golint管的还真宽　多对多的关系先不管
type User struct {
	ID             uint   `gorm:"primary_key"`
	OPENDID        string `gorm:"type:varchar(50);not null;"`
	AVATAR         string `gorm:"type:varchar(50);not null;"`
	NAME           string `gorm:"type:varchar(20);not null;"`
	COLLEG         string `gorm:"type:varchar(20);not null;"`
	MAJOR          string `gorm:"type:varchar(20);not null;"`
	CLASS          string `gorm:"type:varchar(10);not null;"`
	personnalImage string `gorm:"type:varchar(10);not null;"`
}
