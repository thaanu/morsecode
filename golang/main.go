package main

import (
	routes "golang/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.Setup(app)
    log.Fatal(app.Listen(":3000"))
}

