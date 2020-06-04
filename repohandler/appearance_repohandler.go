package repohandler

import (
	"MindaZepeto/database"
	"MindaZepeto/models"
)

var db = database.GetDB()

// func AddAppearance(typeAppearance string, url string) {
// 	appearance := models.Appearance{TYPE: typeAppearance, URL: url}
// 	db.Create(&appearance)
// }

// func DeleteAppearance(id int) {
// 	appearance := models.Appearance{ID: id}
// 	db.Delete(&appearance)
// }

// func ModifyAppearance(typeAppearance string, url string, id int) {
// 	appearance := models.Appearance{TYPE: typeAppearance, URL: url}
// 	db.Model(models.Appearance{}).Where("ID = ?", id).Update(&appearance)

// }

func GetAppearanceList() interface{} {
	type ResultList struct {
		featureList    *[]models.Feature
		hairList       *[]models.Hair
		expressionList *[]models.Expression
		overcoatList   *[]models.Overcoat
		shirtList      *[]models.Shirt
		trousersList   *[]models.Trousers
		shoesList      *[]models.Shoes
		glassesList    *[]models.Glasses
		othersList     *[]models.Others
	}
	featureList := new([]models.Feature)
	hairList := new([]models.Hair)
	expressionList := new([]models.Expression)
	overcoatList := new([]models.Overcoat)
	shirtList := new([]models.Shirt)
	trousersList := new([]models.Trousers)
	shoesList := new([]models.Shoes)
	glassesList := new([]models.Glasses)
	othersList := new([]models.Others)
	db.Raw(`select * FROM feature`).Scan(featureList)
	db.Raw(`select * FROM hair`).Scan(hairList)
	db.Raw(`select * FROM expression`).Scan(expressionList)
	db.Raw(`select * FROM overcoat`).Scan(overcoatList)
	db.Raw(`select * FROM shirt`).Scan(shirtList)
	db.Raw(`select * FROM trousers`).Scan(trousersList)
	db.Raw(`select * FROM shoes`).Scan(shoesList)
	db.Raw(`select * FROM glasses`).Scan(glassesList)
	db.Raw(`select * FROM others`).Scan(othersList)
	resultlist := &ResultList{
		featureList,
		hairList,
		expressionList,
		overcoatList,
		shirtList,
		trousersList,
		shoesList,
		glassesList,
		othersList,
	}
	return resultlist
}
