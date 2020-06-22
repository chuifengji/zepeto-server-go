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
type ResultList struct {
	FeatureList    []models.Feature
	HairList       []models.Hair
	ExpressionList []models.Expression
	OvercoatList   []models.Overcoat
	ShirtList      []models.Shirt
	TrousersList   []models.Trousers
	ShoesList      []models.Shoes
	GlassesList    []models.Glasses
	OthersList     []models.Others
	SpecialList    []models.Special
}

func GetAppearanceList() *ResultList {
	featureList := new([]models.Feature)
	hairList := new([]models.Hair)
	expressionList := new([]models.Expression)
	overcoatList := new([]models.Overcoat)
	shirtList := new([]models.Shirt)
	trousersList := new([]models.Trousers)
	shoesList := new([]models.Shoes)
	glassesList := new([]models.Glasses)
	othersList := new([]models.Others)
	specialList := new([]models.Special)
	db.Raw(`select * FROM feature`).Scan(&featureList)
	db.Raw(`select * FROM hair`).Scan(&hairList)
	db.Raw(`select * FROM expression`).Scan(&expressionList)
	db.Raw(`select * FROM overcoat`).Scan(&overcoatList)
	db.Raw(`select * FROM shirt`).Scan(&shirtList)
	db.Raw(`select * FROM trousers`).Scan(&trousersList)
	db.Raw(`select * FROM shoes`).Scan(&shoesList)
	db.Raw(`select * FROM glasses`).Scan(&glassesList)
	db.Raw(`select * FROM others`).Scan(&othersList)
	db.Raw(`select * FROM special`).Scan(&specialList)
	resultlist := &ResultList{
		*featureList,
		*hairList,
		*expressionList,
		*overcoatList,
		*shirtList,
		*trousersList,
		*shoesList,
		*glassesList,
		*othersList,
		*specialList,
	}
	return resultlist
}
