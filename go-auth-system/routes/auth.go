package routes

import (
	"aymane/controllers"

	"github.com/gofiber/fiber/v2"
)

func UseAuth(app *fiber.App) {

	api := app.Group("/api/user")

	api.Get("/", controllers.Hello)

	api.Post("/signup", controllers.SignUp)

}
