package controllers

import (
	"MindaZepeto/config"
	"MindaZepeto/repohandler"
	"MindaZepeto/util"
	"strconv"

	"github.com/kataras/iris"
)

//AddBackground 添加背景图片（是不是还得保证不重复，别忘了哈）
func AddBackground(ctx iris.Context) {
	value := ctx.GetCookie("usertoken")
	if value == config.Sysconfig.UserToken {
		name := ctx.URLParam("name")
		url := ctx.URLParam("url")
		repohandler.AddBackground(name, url)
		ctx.WriteString("200")
	} else {

	}

}

//DeleteBackground 删除背景图片
func DeleteBackground(ctx iris.Context) {
	value := ctx.GetCookie("usertoken")
	if value == config.Sysconfig.UserToken {
		id := ctx.URLParam("id")
		idMirr, err := strconv.Atoi(id) //string 类型转换成int
		if err != nil {
			panic(err)
		}
		repohandler.DeleteBackground(idMirr)
		ctx.WriteString("200")
	} else {

	}

}

// ModifyBackground 修改背景图片
func ModifyBackground(ctx iris.Context) {
	value := ctx.GetCookie("usertoken")
	if value == config.Sysconfig.UserToken {
		id := ctx.URLParam("id")
		idMirr, err := strconv.Atoi(id) //string 类型转换成int
		if err != nil {
			panic(err)
		}
		name := ctx.URLParam("name")
		url := ctx.URLParam("url")
		repohandler.ModifyBackground(name, url, idMirr)
		ctx.WriteString("200")
	} else {

	}

}

//GetBackgroundList 获取背景图片列表
func GetBackgroundList(ctx iris.Context) {
	list := repohandler.GetBackgroundList()
	result := util.GetReturnData(list, "SUCCESS")
	ctx.JSON(result)
}
