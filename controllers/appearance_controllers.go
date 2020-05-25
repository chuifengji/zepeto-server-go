package controllers

import (
	"MindaZepeto/models"
	"MindaZepeto/repohandler"
	"strconv"

	"github.com/kataras/iris"
)

//Add_appearance 变量不能命名为ｔｙｐｅ，
func Add_appearance(ctx iris.Context) {
	typeAppearance := ctx.PostValue("type")
	url := ctx.PostValue("url")
	repohandler.AddAppearance(typeAppearance, url)
}

func Delete_appearance(ctx iris.Context) {
	id := ctx.PostValue("id")
	id_mirr, err := strconv.Atoi(id) //string 类型转换成int
	if err != nil {
		panic(err)
	}
	repohandler.DeleteAppearance(id_mirr)
}

func Modify_appearance(ctx iris.Context) {
	id := ctx.PostValue("id")
	id_mirr, err := strconv.Atoi(id) //string 类型转换成int
	if err != nil {
		panic(err)
	}
	typeAppearance := ctx.PostValue("type")
	url := ctx.PostValue("url")
	result := repohandler.ModifyAppearance(typeAppearance, url, id_mirr)
	ctx.WriteString(result)
}

func Get_appearance_List(ctx iris.Context) {
	list := repohandler.GetAppearanceList()
	result := new(models.Result)
	result.Data = list
	result.Code = 200
	result.Msg = "SUCCESS"
	ctx.JSON(result) //ctx.JSON用来返回之前ｒｅｓｕｌｔ接口保存的json类型数据，其实这里应该再写一个ｓｅｒｖｉｃｅ层专门处理ｒｅｓｕｌｔ
}
