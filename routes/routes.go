package routes

import (
	c "go-workshop/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	//CRUD Project UserProfile
	user := v1.Group("/user")
	user.Get("/", c.GetUsers)
	user.Get("/filter", c.GetUser)

	user.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"testgo": "23012023",
		},
	}))

	user.Post("/", c.AddUser)
	user.Put("/:id", c.UpdateUser)
	user.Delete("/:id", c.DeleteUser)

	//ข้อ 5.0
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566",
		},
	}))
	//ข้อ 5.1
	v1.Post("/fact/:number", c.FiveDotOne)
	v1.Get("/TestParams", c.TestParams)

	//CRUD dogs
	dog := v1.Group("/dog")
	dog.Get("", c.GetDogs)
	dog.Get("/filter", c.GetDog)
	dog.Get("/json", c.GetDogsJson) //7.2
	dog.Post("/", c.AddDog)         //create
	dog.Put("/:id", c.UpdateDog)
	dog.Delete("/:id", c.RemoveDog)

	dog.Get("/lostdog", c.ShowDelete) //7.0.2 แสดงข้อมูลที่ถูกลบไปแล้ว
	dog.Get("/Scope", c.GetDogsScope) //7.1

	//CRUD Company
	company := v1.Group("/company")
	company.Get("", c.GetCompanies)
	company.Get("/filter", c.GetCompany)
	// company.Post("/json", c.GetCompaniesJson)
	company.Post("/", c.AddCompany)
	company.Put("/:id", c.UpdateCompany)
	company.Delete("/:id", c.RemoveCompany)

	//ข้อ 5.2
	v3 := api.Group("/v3")
	v3.Post("/jim", c.FiveDotTwo)

	//ข้อ 6
	v1.Post("/register", c.Six)
}
