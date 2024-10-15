package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lamhoangvu217/emailery-be-golang/database"
	"github.com/lamhoangvu217/emailery-be-golang/models"
	"github.com/lamhoangvu217/emailery-be-golang/utils"
)

func GetUserDetail(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	if cookie == "" {
		return c.Status(401).JSON(fiber.Map{"message": "Not authenticated"})
	}
	email, err := utils.ParseJwt(cookie)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid token"})
	}
	var user models.User
	database.DB.Where("email=?", email).First(&user)
	if user.Id == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}
	return c.JSON(fiber.Map{
		"message": "Token is correct",
		"user":    user,
	})
}
