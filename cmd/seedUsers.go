package main

import (
	"oddbox/src/database"
	"oddbox/src/models"

	"github.com/bxcodec/faker/v4"
)

func main() {
	database.Connect()

	for i := 0; i < 30; i++ {
		user := models.User{
			Username: faker.FirstName(),
			Email: faker.Email(),
		}
		user.Password, _ = user.HashPassword("12345678")

		database.DB.Create(&user)
	}
}