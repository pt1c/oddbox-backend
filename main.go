package main

import (
	"log"
	"oddbox/src/config"
	"oddbox/src/database"
	"oddbox/src/router"

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