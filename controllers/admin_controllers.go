package controllers

import (
	"MindaZepeto/config"
	"MindaZepeto/util"
	"fmt"

	"github.com/kataras/iris"
)

//AdminLogin 管理端登录页面
func AdminLogin(ctx iris.Context) {
	value := ctx.GetCookie("usertoken")
	if value == config.Sysconfig.UserToken {
		ctx.Redirect("/admin")
	} else {
		ctx.View("login.html")
	}
}

//Admin 管理页面
func Admin(ctx iris.Context) {
	value := ctx.GetCookie("usertoken")
	if value == config.Sysconfig.UserToken {
		ctx.View("admin.html")
	} else {
		ctx.Redirect("/admin/login")
	}
}

// Loginverity 登录验证
func Loginverity(ctx iris.Context) {
	username := ctx.URLParam("username")
	pwd := ctx.URLParam("password")
	fmt.Println(pwd)
	value := util.GetSHAEncode(util.GetSHAEncode(username + pwd))
	ctx.SetCookieKV("usertoken", value)
	if value == config.Sysconfig.UserToken {
		ctx.WriteString("200")
	} else {
		ctx.WriteString("400")
	}
}
