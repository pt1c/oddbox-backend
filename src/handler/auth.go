package handler

import (
	"errors"
	"oddbox/src/config"
	"oddbox/src/database"
	"oddbox/src/models"
	"time"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByUsernameOrEmail(identity string) (*models.User, error) {
	var user models.User

	if err := database.DB.Where("username = ? OR email = ?", identity, identity).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Login get user and password
func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}
	type UserData struct {
		Id       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	input := new(LoginInput)
	var userData UserData

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"status": "error",
				"message": "Error on login request",
				"data": err,
			},
		)
	}

	identity := input.Identity
	pass := input.Password

	userModel, _ := getUserByUsernameOrEmail(identity)

	if userModel == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"status": "error",
				"message": "User not found",
			},
		)
	} else {
		userData = UserData{
			Id:       userModel.Id,
			Username: userModel.Username,
			Email:    userModel.Email,
			Password: userModel.Password,
		}
	}

	if !CheckPasswordHash(pass, userData.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"status": "error",
				"message": "Invalid password",
			},
		)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = userData.Username
	claims["user_id"] = userData.Id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(
		fiber.Map{
			"id": userData.Id,
			"username": userData.Username,
			"email": userData.Email,
			"token": t,
		},
	)
}

func CheckLogin(c *fiber.Ctx) error {
	var user models.User

	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	if err := database.DB.Where("id = ?", uid).Find(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"status": "error",
				"message": "User not found",
			},
		)
	}

	return c.JSON(&user)
}