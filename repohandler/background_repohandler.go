package repohandler

import "MindaZepeto/models"

func AddBackground(name string, url string) {
	background := models.Background{Name: name, URL: url}
	db.Create(&background)
}

func DeleteBackground(id int) {
	background := models.Background{ID: id}
	db.Delete(&background)
}

func ModifyBackground(name string, url string, id int) {
	background := models.Background{Name: name, URL: url}
	db.Model(models.Background{}).Where("ID = ?", id).Update(&background)
}

func GetBackgroundList() *[]models.Background {
	list := new([]models.Background)
	err := db.Raw(`select * FROM background`).Scan(list).RowsAffected //RowsAffected:影响行数
	if err > 0 {
		return list
	} else {
		return nil
	}
}
