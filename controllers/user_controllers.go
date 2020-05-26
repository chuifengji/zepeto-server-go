package controllers

import (
	"MindaZepeto/models"
	"MindaZepeto/util"
	"MindaZepeto/wxlogin"

	"github.com/kataras/iris"
)

func Login(ctx iris.Context) {
	code := ctx.URLParam("code")
	// 根据code获取 openID 和 session_key
	wxLoginResp, err := wxlogin.WXLogin(code)
	if err != nil {
		result := new(models.Result)
		result.Data = nil
		result.Code = 200
		result.Msg = "FAILED"
		ctx.JSON(result)
	}
	// 保存登录态
	USER_ID := util.GetMD5Encode(wxLoginResp.OpenID)
	//接下来需要验证数据库中是否存在该USER_ID
	ctx.WriteString(USER_ID)
}
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
