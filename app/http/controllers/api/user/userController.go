package user

import (
	"referrer/app/database"
	"referrer/app/models"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Where("is_referrer = true").Find(&users)

	return c.JSON(users)
}
