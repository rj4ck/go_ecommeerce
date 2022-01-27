package routes

import (
	"github.com/gofiber/fiber/v2"
	"server/src/controllers"
)

func Categories(app *fiber.App) {
	api := app.Group("/categories")

	api.Get("/", controllers.GetAllCategories)
}
