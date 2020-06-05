package models

//Friends 这golint管的还真宽
type Friends struct {
	MyID     string `gorm:"type:varchar(10);not null;index"`
	FriendID string `gorm:"type:varchar(10);not null;"`
}

//好友关系表的定义，不知道有没有更好的实现方式。
