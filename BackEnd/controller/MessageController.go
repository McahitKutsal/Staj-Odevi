package controller

import (
	"modules/database"
	"modules/models"

	"github.com/gofiber/fiber/v2"
)

//Yapılacak tek bir işimiz olduğu için tek bir fonksiyonumuz var
func SendMessage(c *fiber.Ctx) error {

	var data map[string]string
	//Burada c isimli parametre olarak gelen fiber context nesnesini data map'ine okuyoruz
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	//User nesnesinin bir örneğini oluşturarak elimizdeki map'i üzerine yazıyoruz
	var user = models.User{
		Name:    data["name"],
		Email:   data["email"],
		Message: data["message"],
	}
	//
	database.DB.Create(&user)

	return c.JSON(user)
}
