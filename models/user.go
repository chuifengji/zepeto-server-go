package models

//User USERID 也应该是primary_key
type User struct {
	ID          int    `gorm:"primary_key"`
	USERID      string `gorm:"type:varchar(70);index"`
	NAME        string `gorm:"type:varchar(20)"`
	COLLEGE     string `gorm:"type:varchar(20)"`
	MAJOR       string `gorm:"type:varchar(20)"`
	CLASS       string `gorm:"type:varchar(10)"`
	MYIMG       string `gorm:"type:varchar(35)"`
	CanSearchMe string `gorm:"type:varchar(5)"`
}

//Users 当返回的是user的集合时
type Users struct {
	ID      int
	NAME    string
	COLLEGE string
	MAJOR   string
	CLASS   string
	MYIMG   string
}
