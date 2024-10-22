package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lamhoangvu217/emailery-be-golang/controller"
)

func Setup(app *fiber.App) {
	// app.Use(middleware.IsAuthenticate)
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Get("/api/user", controller.GetUserDetail)
	app.Post("/api/logout", controller.Logout)
	app.Post("/api/update-plans", controller.UpdateUserType)
	app.Post("/api/user", controller.UpdateUserInfo)
}
