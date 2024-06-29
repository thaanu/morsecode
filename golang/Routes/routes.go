package routes

import (
	Controllers "golang/Controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup( app *fiber.App) {
	app.Get("/", Controllers.Home)
	app.Post("/encode", Controllers.EncodeMorseCode)
	app.Post("/decode", Controllers.DecodeMorseCode)
}