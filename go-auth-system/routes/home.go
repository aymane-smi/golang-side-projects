package routes

import (
	"aymane/controllers"
	"aymane/middleware"

	"github.com/gofiber/fiber/v2"
)

func UseHome(app *fiber.App) {
	api := app.Group("/api/home", middleware.Authorization)

	api.Get("/", controllers.Home)
}
