package middleware

import (
	"os"
	"pretest-golang-tdi/util"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

// Protected middleware untuk memeriksa apakah token valid
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtSecret,
		ErrorHandler: jwtError,
		Claims:       &util.Claims{},
	})
}

// Authorize middleware untuk memeriksa peran pengguna
func Authorize(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(*util.Claims)
		userRole := claims.Role

		for _, role := range allowedRoles {
			if role == userRole {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   "Forbidden",
			"message": "Anda tidak memiliki akses ke sumber daya ini.",
		})
	}
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}
