package controllers

import (
	"github.com/kataras/iris"
)

func Add_decoration(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}

func Delete_decoration(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}

func Modify_decoration(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}

func Get_decoration_List(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}
