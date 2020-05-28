package main

import (
	"MindaZepeto/controllers"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.PartyFunc("/decoration", func(decoration iris.Party) {
		decoration.Post("/add", controllers.Add_decoration)
		decoration.Post("/delete", controllers.Delete_decoration)
		decoration.Post("/modify", controllers.Modify_decoration)
		decoration.Get("/get-list", controllers.Get_decoration_List)
	})
	app.PartyFunc("/background", func(background iris.Party) {
		background.Post("/add", controllers.Add_background)
		background.Post("/delete", controllers.Delete_background)
		background.Post("/modify", controllers.Modify_background)
		background.Get("/get-list", controllers.Get_background_List)
	})

	app.PartyFunc("/appearance", func(appearance iris.Party) {
		appearance.Post("/add", controllers.Add_appearance)
		appearance.Post("/delete", controllers.Delete_appearance)
		appearance.Post("/modify", controllers.Modify_appearance)
		appearance.Get("/get-list", controllers.Get_appearance_List)
	})

	app.PartyFunc("/personal", func(personal iris.Party) {
		personal.Get("/login", controllers.Login)
		personal.Post("/insert-personal-info", controllers.UpdatePersonalInfo)
		personal.Post("/update-personal-image", controllers.UpdatePersonalImage)
		personal.Get("/get-friends-list", controllers.GetFriendsList)
		personal.Get("/get-self-info", controllers.GetSelfInfo)
	})

	app.Run(iris.Addr(":8080"))
}
