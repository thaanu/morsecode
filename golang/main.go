package main

import (
	Routes "golang/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func init() {
	
}

func main() {
	app := fiber.New()
	Routes.Setup(app)
    log.Fatal(app.Listen(":3000"))
}

