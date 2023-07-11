package handler

import (
	"ambassador/src/database"
	"ambassador/src/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// GetUser get a user
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User

	database.DB.Find(&user, id)
	if user.Username == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status": "error", 
				"message": "No user found with ID",
			},
		)
	}

	return c.JSON(
		fiber.Map{
			"status": "success",
			"message": "User found",
			"data": user,
		},
	)
}

// CreateUser new user
func CreateUser(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"status": "error", 
				"message": "Review your input", 
				"data": err,
			},
		)
	}

	user := models.User{
		Username: data["username"],
		Email:    data["email"],
		Password:	data["password"],
	}

	//validate fields
	if err := validator.New().Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"status": "error",
				"message": err.Error(),
			},
		)
	}

	//validate passwords match
	if data["password"] != data["password_confirm"] {
		return c.Status(500).JSON(
			fiber.Map{
				"status": "error",
				"message": "Passwords are not same",
			},
		)
	}

	//Check if user already presents
	result := database.DB.Where("username = ? OR email = ?", user.Username, user.Email).First(&user)
	if result.RowsAffected > 0 {
		return c.Status(500).JSON(
			fiber.Map{
				"status": "error",
				"message": "User Already Present",
			},
		)
	}
	
	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"status": "error", 
				"message": "Couldn't hash password", 
				"data": err,
			},
		)
	}
	user.Password = hash
	
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"status": "error",
				"message": "Couldn't create user",
				"data": err,
			},
		)
	}

	return c.JSON(
		fiber.Map{
			"status": "success",
			"message": "Created user",
			"data": user,
		},
	)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}