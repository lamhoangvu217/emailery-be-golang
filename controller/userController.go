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

func UpdateUserType(c *fiber.Ctx) error {
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
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request body"})
	}
	if newUserType, ok := data["user_type"].(string); ok {
		user.UserType = newUserType
	}
	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to update user type"})
	}
	return c.JSON(fiber.Map{
		"message": "Your plans is updated successfully!!",
		"user":    user,
	})
}

func UpdateUserInfo(c *fiber.Ctx) error {
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
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request body"})
	}
	// Update user info
	if err := database.DB.Model(&user).Updates(data).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to update user info",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "User information updated successfully!",
		"user":    user,
	})
}
