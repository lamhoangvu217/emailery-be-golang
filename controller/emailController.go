package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lamhoangvu217/emailery-be-golang/database"
	"github.com/lamhoangvu217/emailery-be-golang/models"
	"github.com/lamhoangvu217/emailery-be-golang/services"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type EmailSendingRequest struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randSource := rand.NewSource(time.Now().UnixNano()) // Create a new random source
	random := rand.New(randSource)                      // Create a new random generator
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}
	return string(result)
}

func GenerateEmail(c *fiber.Ctx) error {
	// Generate a random string of length 10 (you can adjust this length)
	emailPrefix := generateRandomString(10)
	email := models.Email{
		ID:      uuid.New().String(),
		Address: emailPrefix + "@temp-mail.com",
	}

	// Save to the database
	if err := database.DB.Create(&email).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate email"})
	}

	return c.JSON(email)
}

func GetEmails(c *fiber.Ctx) error {
	address := c.Params("address")
	var emails []models.Message

	// Query messages for the given email address.
	if err := database.DB.Where("email_id = ?", address).Find(&emails).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch emails"})
	}

	return c.JSON(emails)
}

func GetEmail(c *fiber.Ctx) error {
	id := c.Params("id")
	var email models.Message

	// Query the email.
	if err := database.DB.First(&email, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Email not found"})
	}

	return c.JSON(email)
}

func GetAllTempEmails(c *fiber.Ctx) error {
	emails, err := services.GetAllEmailsService()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "get all temp emails successfully",
		"emails":  emails,
	})
}

func DeleteTempEmail(c *fiber.Ctx) error {
	tempEmailStr := c.Params("id")
	if tempEmailStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "temp email id is required",
		})
	}
	var tempEmail models.Email
	if err := database.DB.Where("id = ?", tempEmailStr).First(&tempEmail).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "temp email id not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not retrieve temp email",
		})
	}

	if err := services.DeleteTempEmail(&tempEmail, tempEmailStr); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete temp email",
		})
	}
	return c.JSON(fiber.Map{
		"message": "delete temp email successfully",
	})
}

func SendEmail(c *fiber.Ctx) error {
	var emailRequest EmailSendingRequest

	// Parse the JSON body into the EmailRequest struct
	if err := c.BodyParser(&emailRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	m := gomail.NewMessage()
	m.SetHeader("From", emailRequest.From)
	m.SetHeader("To", emailRequest.To)
	m.SetHeader("Subject", emailRequest.Subject)
	m.SetBody("text/plain", emailRequest.Body)

	// Set up the SMTP server details
	d := gomail.NewDialer("localhost", 1025, "", "")

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to send email",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Email sent",
	})
}
