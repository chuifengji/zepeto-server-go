package main

import (
	"MindaZepeto/config"
	"MindaZepeto/controllers"

	"github.com/kataras/iris"
)

func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}
func main() {
	app := iris.New()
	app.Use(Cors)
	app.RegisterView(iris.HTML("./views", ".html"))
	app.Get("/admin", controllers.Admin)
	app.Get("/admin/login", controllers.AdminLogin)
	app.Post("/admin/loginverity", controllers.Loginverity)

	app.PartyFunc("/decoration", func(decoration iris.Party) {
		decoration.Post("/add", controllers.AddDecoration)
		decoration.Post("/delete", controllers.DeleteDecoration)
		decoration.Post("/modify", controllers.ModifyDecoration)
		decoration.Get("/get-list", controllers.GetDecorationList)
	})
	app.PartyFunc("/background", func(background iris.Party) {
		background.Post("/add", controllers.AddBackground)
		background.Post("/delete", controllers.DeleteBackground)
		background.Post("/modify", controllers.ModifyBackground)
		background.Get("/get-list", controllers.GetBackgroundList)
	})

	app.PartyFunc("/appearance", func(appearance iris.Party) {
		appearance.Post("/add", controllers.AddAppearance)
		appearance.Post("/delete", controllers.DeleteAppearance)
		appearance.Post("/modify", controllers.ModifyAppearance)
		appearance.Get("/get-list", controllers.GetAppearanceList)
	})

	app.PartyFunc("/personal", func(personal iris.Party) {
		personal.Get("/login", controllers.Login)
		personal.Post("/insert-personal-info", controllers.UpdatePersonalInfo)
		personal.Post("/update-personal-image", controllers.UpdatePersonalImage)
		personal.Post("/makeFriends", controllers.MakeFriends)
		personal.Post("/get-search-list", controllers.SearchUser)
		personal.Get("/get-friends-list", controllers.GetFriendsList)
		personal.Get("/get-self-info", controllers.GetSelfInfo)
		personal.Get("/get-classmates-list", controllers.GetClassmateList)
		personal.Get("/get-all-user-info", controllers.GetAllUserInfo)
	})

	app.Run(iris.Addr(config.Sysconfig.Port))
}
