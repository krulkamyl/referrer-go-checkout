package middlewares

import (
	"referrer/app/config"
	"strconv"

	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type ClaimsWithScope struct {
	jwt.RegisteredClaims
	Scope string
}

func IsAuthenticated(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &ClaimsWithScope{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)

		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	payload := token.Claims.(*ClaimsWithScope)
	isReferrer := strings.Contains(c.Path(), "/api/referrer")

	if (payload.Scope == "admin" && isReferrer) || (payload.Scope == "referrer" && !isReferrer) {
		c.Status(fiber.StatusUnauthorized)

		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Next()
}

func GenerateJWT(id uint, scope string) (string, error) {
	now := time.Now()
	payload := ClaimsWithScope{}

	payload.Subject = strconv.Itoa(int(int(id)))
	payload.ExpiresAt = &jwt.NumericDate{now.Add(time.Hour * 24)}
	payload.Scope = scope

	return jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(config.Getenv("JWT_SECRET")))
}

func GetUserId(c *fiber.Ctx) (uint, error) {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &ClaimsWithScope{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	payload := token.Claims.(*ClaimsWithScope)

	id, _ := strconv.Atoi(payload.Subject)

	return uint(id), nil
}
