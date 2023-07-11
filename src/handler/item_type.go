package handler

import (
	"oddbox/src/database"
	"oddbox/src/models"

	"github.com/gofiber/fiber/v2"
)

// CreateItemType new ItemType
func CreateItemType(c *fiber.Ctx) error {
	db := database.DB

	itemType := new(models.ItemsType)

	if err := c.BodyParser(itemType); err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"status": "error", 
				"message": "Couldn't create ItemType", 
				"data": err,
			},
		)
	}

	db.Create(&itemType)

	return c.JSON(
		fiber.Map{
			"status": "success", 
			"message": "Created ItemType", 
			"data": itemType,
		},
	)

}