package main

import (
	"go-workshop/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.InetRoutes(app)
	app.Listen(":3000")
}
