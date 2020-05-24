package controllers

import (
	"github.com/kataras/iris"
)

func Add_background(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}

func Delete_background(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}

func Modify_background(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}

func Get_background_List(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}
