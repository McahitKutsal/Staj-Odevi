package main

import (
	"modules/database"
	"modules/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()

	app := fiber.New()
	//Front end'den veri gönderebilmek için Cors altında Credentials'ı açmamız gerekiyor açmayınca hata veriyor
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routers.Setup(app)

	app.Listen(":8000")
}
