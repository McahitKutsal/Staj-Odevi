package routers

import (
	"modules/controller"

	"github.com/gofiber/fiber/v2"
)

//Tek bir route yeterli
func Setup(app *fiber.App) {
	/*localhost 8000 portuna "api/mesajgonder" adresine post isteği yapıldıpında
	alttaki route devreye girip controller Package'ının altında SendMessage Fonksiyonu çalışacak*/
	app.Post("/api/mesajgonder", controller.SendMessage)
}
