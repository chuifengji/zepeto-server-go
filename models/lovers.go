package models

//Lovers 留着备用，以后或许会加
type Lovers struct {
	MyID    string `gorm:"type:varchar(10);not null;"`
	LoverID string `gorm:"type:varchar(10);not null;"`
}
