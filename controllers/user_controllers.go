package controllers

import "github.com/kataras/iris"

func Insert_personal_info(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}

func Update_personal_image(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}

func Get_friends_list(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}

func Get_self_info(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}

func Make_friends(ctx iris.Context) {

}
