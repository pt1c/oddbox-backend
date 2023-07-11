package main

import (
	"ambassador/src/config"
	"ambassador/src/database"
	"ambassador/src/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	database.AutoMigrate()

	app := fiber.New()
	app.Use(cors.New())

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":" + config.Config("SERVER_PORT")))
}