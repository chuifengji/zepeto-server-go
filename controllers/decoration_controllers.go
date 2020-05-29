package controllers

import (
	"MindaZepeto/repohandler"
	"MindaZepeto/util"
	"strconv"

	"github.com/kataras/iris"
)

//AddDecoration 添加饰品
func AddDecoration(ctx iris.Context) {

	name := ctx.PostValue("name")
	url := ctx.PostValue("url")
	repohandler.AddDecoration(name, url)

}

//DeleteDecoration 删除饰品
func DeleteDecoration(ctx iris.Context) {
	id := ctx.PostValue("id")
	idMirr, err := strconv.Atoi(id) //string 类型转换成int
	if err != nil {
		panic(err)
	}
	repohandler.DeleteDecoration(idMirr)
}

//ModifyDecoration 修改饰品信息
func ModifyDecoration(ctx iris.Context) {
	id := ctx.PostValue("id")
	idMirr, err := strconv.Atoi(id) //string 类型转换成int
	if err != nil {
		panic(err)
	}
	name := ctx.PostValue("name")
	url := ctx.PostValue("url")
	result := repohandler.ModifyDecoration(name, url, idMirr)
	ctx.WriteString(result)
}

//GetDecorationList 获取饰品列表
func GetDecorationList(ctx iris.Context) {
	list := repohandler.GetDecorationList()
	result := util.GetReturnData(list, "SUCCESS")
	ctx.JSON(result)
}
