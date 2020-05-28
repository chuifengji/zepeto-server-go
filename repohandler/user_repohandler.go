package repohandler

import (
	"MindaZepeto/models"
)

// verifyHasThisUserI 验证该ＵＳＥＲＩＤ是否已经注册过。
func VerifyHasThisUserID(USERID string) models.User {
	user := models.User{}
	//	db.Where(models.User{USERID: USERID}).FirstOrInit(&user)
	db.Where("user_id = ?", USERID).Find(&user)
	return user
}

func CreateUser(USER_ID string) models.User {
	user := models.User{USERID: USER_ID}
	db.Create(&user)
	return user
}

func UpdateInfo(user_id string, name string, college string, major string, class string) *models.User {
	user := new(models.User)
	result := new(models.User) //更新之后并没有返回最新的值，而是返回的修改对象本身，所以还得自己重新查询。
	err := db.Model(&user).Where("user_id = ?", user_id).Updates(map[string]interface{}{"NAME": name, "MAJOR": major, "COLLEGE": college, "CLASS": class}).RowsAffected
	if err > 0 {
		db.Model(&user).Where("user_id = ?", user_id).First(&result)
		return result
	} else {
		return nil
	}
}

func UpdateSelfImage(user_id string, url string) *models.User {
	user := new(models.User)
	result := new(models.User) //更新之后并没有返回最新的值，而是返回的修改对象本身，所以还得自己重新查询。
	err := db.Model(&user).Where("user_id = ?", user_id).Update("url", url).RowsAffected
	if err > 0 {
		db.Model(&user).Where("user_id = ?", user_id).First(&result)
		return result
	} else {
		return nil
	}
}

func MakeFriend(userid string, friendid string) *models.Friends {
	list := new(models.Friends)
	list.UserID = userid
	list.FriendID = friendid
	result := new(models.Friends)
	err := db.Create(&list).RowsAffected
	if err > 0 {
		db.Model(&list).Where("user_id = ?", userid).First(&result)
		return result
	} else {
		return nil
	}
}

func GetSelfInfo(userid string) *models.User {
	result := new(models.User)
	err := db.Model(&result).Where("user_id = ?", userid).First(&result).RowsAffected
	if err > 0 {
		return result
	} else {
		return nil
	}
}

func GetClassmateList(college string, major string, class string) *models.User {
	result := new(models.User)
	err := db.Model(&result).Where("college = ? AND major = ? AND class=?", college, major, class).Find(&result).RowsAffected
	if err > 0 {
		return result
	} else {
		return nil
	}
}
