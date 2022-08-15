package main

import (
	"aymane/config"
	"aymane/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config.Connect()

	app := fiber.New()

	routes.UseAuth(app)

	log.Fatal(app.Listen(":3000"))
}
