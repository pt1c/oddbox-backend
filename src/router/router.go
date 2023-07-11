package router

import (
	"ambassador/src/handler"
	"ambassador/src/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	
	api.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello,dd W21ors11231")
	})

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)
	auth.Get("/login", middleware.Protected(), handler.CheckLogin)
	
	// User
	user := api.Group("/user")
	user.Post("/register", handler.CreateUser)
	user.Get("/:id", middleware.Protected(), handler.GetUser)
	

	// ItemType
	ItemType := api.Group("/item-type")
	ItemType.Post("/", handler.CreateItemType)
	
}