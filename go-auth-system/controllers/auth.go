package controllers

import (
	"aymane/config"
	"aymane/model"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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

func SignIn(c *fiber.Ctx) error {
	godotenv.Load()
	var body map[string]string

	c.BodyParser(&body)

	if body["email"] == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "please enter a valid email",
		})
	}
	if body["password"] == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "password missing",
		})
	}

	config.DB.Where("email = ?", body["email"]).First(&user)

	if user.Email == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "email not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(body["password"])); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, token_err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if token_err != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": token_err,
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "welcome",
		"user":    user,
		"token":   token,
	})

}
