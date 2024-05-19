package routes

import (
	c "go-workshop/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	//ข้อ 5.0
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566",
		},
	}))
	//ข้อ 5.1
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/fact/:number", c.FiveDotOne)

	//ข้อ 5.2
	v3 := api.Group("/v3")
	jim := v3.Group("/jim")
	jim.Post("/", c.FiveDotTwo)

	//ข้อ 6
	v1.Post("/register", c.Six)
}
