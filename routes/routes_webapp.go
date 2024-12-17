package routes

import (
	c "golangCourse/controllers"

	"github.com/gofiber/fiber/v2"
)

func WebAppRoutes(app *fiber.App) {
	app.Get("/", c.HelloTestV3)

	api := app.Group("/api")

	v1 := api.Group("/v1")
	v2 := api.Group("/v2")

	v1.Get("/", c.HelloTest)
	v1.Post("/", c.PersonTest)
	v1.Get("/:name", c.ParamsTest)

	// v2.Get("/", c.HelloTestV2)
	v2.Post("/", c.CreateDog)
	v2.Delete("/:id", c.DeleteDog)
	v2.Get("/ddog", c.GetDeleteDogs)
	v2.Get("/", c.GetDogs)
}
