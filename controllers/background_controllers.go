package controllers

import (
	"MindaZepeto/models"
	"MindaZepeto/repohandler"
	"strconv"

	"github.com/kataras/iris"
)

func Add_background(ctx iris.Context) {
	name := ctx.PostValue("name")
	url := ctx.PostValue("url")
	repohandler.AddBackground(name, url)
}

func Delete_background(ctx iris.Context) {
	id := ctx.PostValue("id")
	id_mirr, err := strconv.Atoi(id) //string 类型转换成int
	if err != nil {
		panic(err)
	}
	repohandler.DeleteBackground(id_mirr)
}

func Modify_background(ctx iris.Context) {
	id := ctx.PostValue("id")
	id_mirr, err := strconv.Atoi(id) //string 类型转换成int
	if err != nil {
		panic(err)
	}
	name := ctx.PostValue("name")
	url := ctx.PostValue("url")
	result := repohandler.ModifyBackground(name, url, id_mirr)
	ctx.WriteString(result)
}

func Get_background_List(ctx iris.Context) {
	list := repohandler.GetBackgroundList()
	result := new(models.Result)
	result.Data = list
	result.Code = 200
	result.Msg = "SUCCESS"
	ctx.JSON(result)
}
