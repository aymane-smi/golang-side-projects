package middleware

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Authorization(c *fiber.Ctx) error {
	godotenv.Load()
	cookie := c.Cookies("jwt")

	_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Next()

}
