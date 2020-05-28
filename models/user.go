package models

//User 这golint管的还真宽　多对多的关系先不管
type User struct {
	ID             int    `gorm:"primary_key"`
	USERID         string `gorm:"type:varchar(50)"`
	NAME           string `gorm:"type:varchar(20)"`
	COLLEGE        string `gorm:"type:varchar(20)"`
	MAJOR          string `gorm:"type:varchar(20)"`
	CLASS          string `gorm:"type:varchar(10)"`
	personnalImage string `gorm:"type:varchar(20)"`
}
