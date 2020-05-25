package repohandler

import (
	"MindaZepeto/database"
	"MindaZepeto/models"
)

var db = database.GetDB()

func AddAppearance(typeAppearance string, url string) {
	appearance := models.Appearance{TYPE: typeAppearance, URL: url}
	db.Create(&appearance)
}

func DeleteAppearance(id int) {
	appearance := models.Appearance{ID: id}
	db.Delete(&appearance)
}

func ModifyAppearance(typeAppearance string, url string, id int) string {
	appearance := models.Appearance{TYPE: typeAppearance, URL: url}
	if err := db.Model(models.Appearance{}).Where("ID = ?", id).Update(&appearance).Error; err != nil {
		//错误处理
		return "error"

	} else {
		return "success"
	}

}

func GetAppearanceList() *[]models.Appearance {

	list := new([]models.Appearance)
	err := db.Raw(`select * FROM appearance`).Scan(list).RowsAffected //RowsAffected:影响行数
	if err > 0 {
		return list
	} else {
		return nil
	}

}
