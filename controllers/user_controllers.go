package controllers

import (
	"MindaZepeto/repohandler"
	"MindaZepeto/util"
	"MindaZepeto/wxlogin"

	"github.com/kataras/iris"
)

//Login 注册/登录
func Login(ctx iris.Context) {
	code := ctx.URLParam("code")
	//根据code获取 openID 和 session_key
	wxLoginResp, err := wxlogin.WXLogin(code)
	if err != nil {
		result := util.GetReturnData(nil, "FAILED")
		ctx.JSON(result)
	}
	// // 保存登录态
	USERID := util.GetMD5Encode(wxLoginResp.OpenID)
	//接下来需要验证数据库中是否存在该USER_ID
	list := repohandler.VerifyHasThisUserID(USERID)
	if list.USERID == "" { //不存在该用户
		repohandler.CreateUser(USERID)
	} else {
		result := util.GetReturnData(list, "SUCCESS")
		ctx.JSON(result)
	}
}

//UpdatePersonalInfo 更新个人信息
func UpdatePersonalInfo(ctx iris.Context) {
	userid := ctx.PostValue("user_id")
	name := ctx.PostValue("name")
	college := ctx.PostValue("college")
	major := ctx.PostValue("major")
	class := ctx.PostValue("class")
	list := repohandler.UpdateInfo(userid, name, college, major, class)
	result := util.GetReturnData(list, "SUCCESS")
	ctx.JSON(result)

}

//UpdatePersonalImage 更新个人形象图片
func UpdatePersonalImage(ctx iris.Context) {
	userid := ctx.PostValue("user_id")
	url := ctx.PostValue("url")
	list := repohandler.UpdateSelfImage(userid, url)
	result := util.GetReturnData(list, "SUCCESS")
	ctx.JSON(result)
}

//GetFriendsList 获取朋友列表
func GetFriendsList(ctx iris.Context) {
	myid := ctx.URLParam("myid")
	list := repohandler.GetMyFriendsList(myid)
	result := util.GetReturnData(list, "SUCCESS")
	ctx.JSON(result)
}

//GetClassmateList 获取同班同学列表
func GetClassmateList(ctx iris.Context) {
	college := ctx.URLParam("college")
	major := ctx.URLParam("major")
	class := ctx.URLParam("class")
	list := repohandler.GetClassmateList(college, major, class)
	result := util.GetReturnData(list, "SUCCESS")
	ctx.JSON(result)
}

//GetSelfInfo 获取个人信息
func GetSelfInfo(ctx iris.Context) {
	userid := ctx.URLParam("user_id")
	list := repohandler.GetSelfInfo(userid)
	result := util.GetReturnData(list, "SUCCESS")
	ctx.JSON(result)
}

//GetAllUserInfo 获取所有用户的个人信息，管理端的东西
func GetAllUserInfo(ctx iris.Context) {

}

//MakeFriends 交个朋友呗
func MakeFriends(ctx iris.Context) {
	myid := ctx.PostValue("myid")
	friendid := ctx.PostValue("friendid")
	// list := repohandler.MakeFriends(myid, friendid)
	repohandler.MakeFriends(myid, friendid)
	result := util.GetReturnData(nil, "SUCCESS")
	ctx.JSON(result)
}
