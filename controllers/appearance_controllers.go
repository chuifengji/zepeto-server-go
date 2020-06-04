package controllers

import (
	"MindaZepeto/repohandler"
	"MindaZepeto/util"

	"github.com/kataras/iris"
)

//AddAppearance 变量不能命名为ｔｙｐｅ，
// func AddAppearance(ctx iris.Context) {
// 	value := ctx.GetCookie("usertoken")
// 	if value == config.Sysconfig.UserToken {
// 		typeAppearance := ctx.URLParam("type")
// 		url := ctx.URLParam("url")
// 		repohandler.AddAppearance(typeAppearance, url)
// 		ctx.WriteString("200")
// 	} else {

// 	}
// }

// //DeleteAppearance 删除
// func DeleteAppearance(ctx iris.Context) {
// 	value := ctx.GetCookie("usertoken")
// 	if value == config.Sysconfig.UserToken {
// 		id := ctx.URLParam("id")
// 		idMirr, err := strconv.Atoi(id) //string 类型转换成int
// 		if err != nil {
// 			panic(err)
// 		}
// 		repohandler.DeleteAppearance(idMirr)
// 		ctx.WriteString("200")
// 	} else {

// 	}

// }

// //ModifyAppearance 修改
// func ModifyAppearance(ctx iris.Context) {
// 	value := ctx.GetCookie("usertoken")
// 	if value == config.Sysconfig.UserToken {
// 		id := ctx.URLParam("id")
// 		idMirr, err := strconv.Atoi(id) //string 类型转换成int
// 		if err != nil {
// 			panic(err)
// 		}
// 		typeAppearance := ctx.URLParam("type")
// 		url := ctx.URLParam("url")
// 		repohandler.ModifyAppearance(typeAppearance, url, idMirr)
// 		ctx.WriteString("200")
// 	} else {

// 	}

// }

//暂时来说，以上部分全部作废，我们只需要下面这个接口就好了——2020.6.4

//GetAppearanceList 获取列表信息
func GetAppearanceList(ctx iris.Context) {
	list := repohandler.GetAppearanceList()
	result := util.GetReturnData(list, "SUCCESS")
	ctx.JSON(result)
}
