package controllers

import "github.com/kataras/iris"

func Add_appearance(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}

func Delete_appearance(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}

func Modify_appearance(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}

func Get_appearance_List(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}
