package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/lamhoangvu217/emailery-be-golang/database"
	"github.com/lamhoangvu217/emailery-be-golang/routes"
)

func main() {
	database.Connect()
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load(".env.local")
		if err != nil {
			log.Fatal("Error loading .env.local file")
		}
	}
	port := os.Getenv("PORT")
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000,https://www.emailery.online", // Replace with the actual frontend origin(s)
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowCredentials: true,
	}))
	routes.Setup(app)
	app.Listen(":" + port)

}
