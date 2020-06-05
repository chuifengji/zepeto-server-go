package models

//Grouphoto 合照
type Grouphoto struct {
	ID        int    `gorm:"primary_key"`
	USERID    string `gorm:"type:varchar(40);not null;"`
	LOCATION  string `gorm:"type:varchar(20);not null;"`
	URL       string `gorm:"type:varchar(40);not null;"`
	THUMBNAIL string `gorm:"type:varchar(40);not null;"`
}
