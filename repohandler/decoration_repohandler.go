package repohandler

import (
	"MindaZepeto/models"
)

//AddDecoration add decoration for the background ,param one is its name,param two is its network address.
func AddDecoration(name string, url string) {
	decoration := models.Decoration{Name: name, URL: url}
	db.Create(&decoration)
}

// DeleteDecoration delete as id
func DeleteDecoration(id int) {
	decoration := models.Decoration{ID: id}
	db.Delete(&decoration)
}

//ModifyDecoration y1s1,grom实现的ｗｈｅｒｅ方法一点也不优雅。
func ModifyDecoration(name string, url string, id int) string {
	decoration := models.Decoration{Name: name, URL: url}
	if err := db.Model(models.Decoration{}).Where("ID = ?", id).Update(&decoration).Error; err != nil {
		//错误处理
		return "error"
	} else {
		return "success"
	}
}

func GetDecorationList() *[]models.Decoration {
	list := new([]models.Decoration)
	err := db.Raw(`select * FROM decoration`).Scan(list).RowsAffected //RowsAffected:影响行数
	if err > 0 {
		return list
	} else {
		return nil
	}
}
