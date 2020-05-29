package controllers

import (
	"MindaZepeto/repohandler"
	"MindaZepeto/util"
	"strconv"

	"github.com/kataras/iris"
)

//AddBackground 添加背景图片（是不是还得保证不重复，别忘了哈）
func AddBackground(ctx iris.Context) {
	name := ctx.PostValue("name")
	url := ctx.PostValue("url")
	repohandler.AddBackground(name, url)
}

//DeleteBackground 删除背景图片
func DeleteBackground(ctx iris.Context) {
	id := ctx.PostValue("id")
	idMirr, err := strconv.Atoi(id) //string 类型转换成int
	if err != nil {
		panic(err)
	}
	repohandler.DeleteBackground(idMirr)
}

// ModifyBackground 修改背景图片
func ModifyBackground(ctx iris.Context) {
	id := ctx.PostValue("id")
	idMirr, err := strconv.Atoi(id) //string 类型转换成int
	if err != nil {
		panic(err)
	}
	name := ctx.PostValue("name")
	url := ctx.PostValue("url")
	result := repohandler.ModifyBackground(name, url, idMirr)
	ctx.WriteString(result)
}

//GetBackgroundList 获取背景图片列表
func GetBackgroundList(ctx iris.Context) {
	list := repohandler.GetBackgroundList()
	result := util.GetReturnData(list, "SUCCESS")
	ctx.JSON(result)
}
