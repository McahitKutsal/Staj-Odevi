package controller

import (
	"modules/database"
	"modules/models"

	"github.com/gofiber/fiber/v2"
)

func SendMessage(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user = models.User{
		Name:    data["name"],
		Email:   data["email"],
		Message: data["message"],
	}
	database.DB.Create(&user)

	return c.JSON(user)
}
