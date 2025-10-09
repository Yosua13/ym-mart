package util

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

type Claims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int, role string) (string, error) {
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func GetUserClaims(c *fiber.Ctx) *Claims {
	user := c.Locals("user")
	if user == nil {
		return nil
	}
	token := user.(*jwt.Token)
	claims := token.Claims.(*Claims)
	return claims
}
