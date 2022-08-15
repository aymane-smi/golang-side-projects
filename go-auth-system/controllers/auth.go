package controllers

import (
	"aymane/config"
	"aymane/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var user model.User

func Hello(c *fiber.Ctx) error {
	response := map[string]string{"message": "Hello Page"}
	return c.JSON(response)
}

func SignUp(c *fiber.Ctx) error {
	var body map[string]string

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	//name check
	if body["name"] == "" {
		return fiber.NewError(fiber.StatusBadRequest, "please enter a name")
	} else {
		user.Name = body["name"]
	}

	//email check
	if body["email"] == "" {
		return fiber.NewError(fiber.StatusBadRequest, "please enter a valid email")
	} else {
		user.Email = body["email"]
	}

	//password check
	if body["password"] == "" {
		return fiber.NewError(fiber.StatusBadRequest, "please enter a password")
	} else {
		password, _ := bcrypt.GenerateFromPassword([]byte(body["password"]), 12)
		user.Password = password
	}

	fmt.Println(user)

	config.DB.Create(&user)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "user created",
		"user":    user,
	})

}
