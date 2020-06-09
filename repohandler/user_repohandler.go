package repohandler

import (
	"MindaZepeto/models"
)

//VerifyHasThisUserID 验证该ＵＳＥＲＩＤ是否已经注册过。
func VerifyHasThisUserID(USERID string) models.User {
	user := models.User{}
	db.Where("user_id = ?", USERID).First(&user)
	return user
}

//CreateUser 新注册用户
func CreateUser(USERID string) interface{} {
	user := models.User{USERID: USERID}
	db.Create(&user)
	return user
}

//UpdateInfo 更新个人信息
func UpdateInfo(userid string, name string, college string, major string, class string, canSearchMe string) *models.User {
	result := new(models.User) //更新之后并没有返回最新的值，而是返回的修改对象本身，所以还得自己重新查询。
	err := db.Table("user").Where("user_id = ?", userid).Updates(map[string]interface{}{"NAME": name, "MAJOR": major, "COLLEGE": college, "CLASS": class, "can_search_me": canSearchMe}).RowsAffected
	if err > 0 {
		db.Table("user").Where("user_id = ?", userid).First(&result)
		return result
	} else {
		return nil
	}
}
func CanSearchMe(userid string, canSearchMe string) *models.User {
	result := new(models.User)
	err := db.Table("user").Where("user_id = ?", userid).Update("can_search_me", canSearchMe).RowsAffected
	if err > 0 {
		db.Table("user").Where("user_id = ?", userid).First(&result)
		return result
	} else {
		return nil
	}
}

//UpdateSelfImage 更新个人形象照片
func UpdateSelfImage(userid string, url string) *models.User {
	result := new(models.User) //更新之后并没有返回最新的值，而是返回的修改对象本身，所以还得自己重新查询。
	err := db.Table("user").Where("user_id = ?", userid).Update("myimg", url).RowsAffected
	if err > 0 {
		db.Table("user").Where("user_id = ?", userid).First(&result)
		return result
	} else {
		return nil
	}
}

//MakeFriends 通过返回两个用户的id来建立朋友关系，现在是单向好友,需要先确定两者是否已经建立朋友关系，不能重复。
func MakeFriends(myid string, friendid string) *[]models.Users {
	num := db.Table("friends").Where("my_id = ? AND friend_id = ?", myid, friendid).RowsAffected
	if num > 0 { //该关系已经存在了，不能重复添加。
		return nil
	} else {
		list := new(models.Friends)
		list.MyID = myid
		list.FriendID = friendid
		err := db.Create(&list).RowsAffected
		if err > 0 {
			result := GetMyFriendsList(myid)
			return result
		} else {
			return nil
		}
	}
}

//GetSelfInfo 获取个人信息
func GetSelfInfo(userid string) *models.User {
	result := new(models.User)
	err := db.Model(&result).Where("user_id = ?", userid).First(&result).RowsAffected
	if err > 0 {
		return result
	} else {
		return nil
	}
}

//GetClassmateList 获取某一个指定班级的所有成员信息,不应返回user_id这样隐私的数据。
func GetClassmateList(college string, major string, class string) *[]models.Users {
	var result = new([]models.Users)
	err := db.Raw("SELECT id,name,college,major,class,myimg FROM user WHERE college =?  AND major = ? AND class= ? ", college, major, class).Scan(&result).RowsAffected
	if err > 0 {
		return result
	} else {
		return nil
	}
}

//SearchUserList 返回模糊搜索的列表
func SearchUserList(content string) *[]models.Users {
	var result = new([]models.Users)
	err := db.Raw("SELECT id, name,college,major,class,myimg FROM user WHERE concat(name,college,major)  like ?  AND can_search_me='true' ", "%"+content+"%").Scan(&result).RowsAffected
	if err > 0 {
		return result
	} else {
		return nil
	}
}

//
//GetMyFriendsList 获取某个用户的朋友列表
func GetMyFriendsList(myid string) *[]models.Users {
	var result = new([]models.Users)
	db.Raw("SELECT id,name, college,major,class,myimg FROM user WHERE id in (SELECT friend_id from friends WHERE my_id = ? )", myid).Scan(&result)
	return result
}
func GetAllUserList() *[]models.Users {
	var result = new([]models.Users)
	db.Raw("SELECT id,name,college,major,class,myimg FROM user").Scan(&result)
	return result
}

//AddGroupPhoto 向指定表中添加图片
func AddGroupPhoto(userid string, location string, time string, url string, tableName string) *[]models.Grouphotoend {
	photo := models.Grouphotoend{
		USERID:   userid,
		LOCATION: location,
		URL:      url,
		TIME:     time,
	}
	result := new([]models.Grouphotoend)
	err := db.Table(tableName).Create(&photo).RowsAffected
	if err > 0 {
		db.Table(tableName).Where("user_id = ?", userid).Find(&result)
		return result
	} else {
		return nil
	}
}

//DeleteGroupPhoto 删除某用户的某张图片
func DeleteGroupPhoto(tableName string, idimg string, userid string) *[]models.Grouphotoend {
	nouse := new([]models.Grouphotoend)
	err := db.Raw("DELETE FROM ? WHERE user_id = ? AND  id = ? ", tableName, userid, idimg).Scan(&nouse).RowsAffected
	if err > 0 {
		result := new([]models.Grouphotoend)
		db.Table(tableName).Where("user_id = ?", userid).Find(&result)
		return result
	} else {
		return nil
	}
}

//GetMyPhotos 获取某用户的全部图片
func GetMyPhotos(userid string, tableName string) *[]models.Grouphotoend {
	result := new([]models.Grouphotoend)
	err := db.Table(tableName).Where("user_id = ?", userid).Find(&result).RowsAffected
	if err > 0 {
		return result
	} else {
		return nil
	}
}

func DeleteFriend(myid string, friendid string) *[]models.Users {
	nouse := new(models.Friends)
	err := db.Raw("DELETE FROM friends WHERE my_id = ? AND friend_id = ? ", myid, friendid).Scan(&nouse).RowsAffected
	if err > 0 {
		result := GetMyFriendsList(myid)
		return result
	} else {
		return nil
	}
}

func DeleteAllPhotos(userid string, tableName string) string {
	photos := new(models.Grouphotoend)
	err := db.Table(tableName).Where("user_id = ?", userid).Delete(&photos).RowsAffected
	if err > 0 {
		return "success"
	} else {
		return "failed"
	}
}
