package controllers

import (
	"MindaZepeto/config"
	"MindaZepeto/repohandler"
	"MindaZepeto/util"
	"strconv"

	"github.com/kataras/iris"
)

//AddDecoration 添加饰品
func AddDecoration(ctx iris.Context) {
	value := ctx.GetCookie("usertoken")
	if value == config.Sysconfig.UserToken {
		name := ctx.URLParam("name")
		url := ctx.URLParam("url")
		repohandler.AddDecoration(name, url)
		ctx.WriteString("200")
	} else {

	}

}

//DeleteDecoration 删除饰品
func DeleteDecoration(ctx iris.Context) {
	value := ctx.GetCookie("usertoken")
	if value == config.Sysconfig.UserToken {
		id := ctx.URLParam("id")
		idMirr, err := strconv.Atoi(id) //string 类型转换成int
		if err != nil {
			panic(err)
		}
		repohandler.DeleteDecoration(idMirr)
		ctx.WriteString("200")
	}

}

//ModifyDecoration 修改饰品信息
func ModifyDecoration(ctx iris.Context) {
	value := ctx.GetCookie("usertoken")
	if value == config.Sysconfig.UserToken {
		id := ctx.URLParam("id")
		idMirr, err := strconv.Atoi(id) //string 类型转换成int
		if err != nil {
			panic(err)
		}
		name := ctx.URLParam("name")
		url := ctx.URLParam("url")
		repohandler.ModifyDecoration(name, url, idMirr)
		ctx.WriteString("200")
	} else {

	}
}

//GetDecorationList 获取饰品列表
func GetDecorationList(ctx iris.Context) {
	list := repohandler.GetDecorationList()
	result := util.GetReturnData(list, "SUCCESS")
	ctx.JSON(result)
}
