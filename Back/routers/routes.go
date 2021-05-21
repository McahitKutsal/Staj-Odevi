package routers

import (
	"modules/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/mesajgonder", controller.SendMessage)
}
