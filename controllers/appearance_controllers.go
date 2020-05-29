package controllers

import (
	"MindaZepeto/repohandler"
	"MindaZepeto/util"
	"strconv"

	"github.com/kataras/iris"
)

//AddAppearance 变量不能命名为ｔｙｐｅ，
func AddAppearance(ctx iris.Context) {
	typeAppearance := ctx.PostValue("type")
	url := ctx.PostValue("url")
	repohandler.AddAppearance(typeAppearance, url)
}

//DeleteAppearance 删除
func DeleteAppearance(ctx iris.Context) {
	id := ctx.PostValue("id")
	idMirr, err := strconv.Atoi(id) //string 类型转换成int
	if err != nil {
		panic(err)
	}
	repohandler.DeleteAppearance(idMirr)
}

//ModifyAppearance 修改
func ModifyAppearance(ctx iris.Context) {
	id := ctx.PostValue("id")
	idMirr, err := strconv.Atoi(id) //string 类型转换成int
	if err != nil {
		panic(err)
	}
	typeAppearance := ctx.PostValue("type")
	url := ctx.PostValue("url")
	result := repohandler.ModifyAppearance(typeAppearance, url, idMirr)
	ctx.WriteString(result)
}

//GetAppearanceList 获取列表信息
func GetAppearanceList(ctx iris.Context) {
	list := repohandler.GetAppearanceList()
	result := util.GetReturnData(list, "SUCCESS")
	ctx.JSON(result)
}
