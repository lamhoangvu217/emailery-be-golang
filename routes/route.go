package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lamhoangvu217/emailery-be-golang/controller"
)

func Setup(app *fiber.App) {
	// app.Use(middleware.IsAuthenticate)

	api := app.Group("/api")
	api.Post("/register", controller.Register)
	api.Post("/login", controller.Login)
	api.Get("/user", controller.GetUserDetail)
	api.Post("/logout", controller.Logout)
	api.Post("/update-plans", controller.UpdateUserType)
	api.Post("/user", controller.UpdateUserInfo)

	api.Post("/generate-email", controller.GenerateEmail)
	api.Get("/emails/:address", controller.GetEmails)
	api.Get("/email/:id", controller.GetEmail)

	api.Get("/all-temp-emails", controller.GetAllTempEmails)
	api.Delete("/temp-email/:id", controller.DeleteTempEmail)
	api.Post("/send-email", controller.SendEmail)
}
