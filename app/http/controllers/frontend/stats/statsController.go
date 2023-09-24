package stats

import (
	"referrer/app/database"
	"referrer/app/http/middlewares"
	"referrer/app/models"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	id, _ := middlewares.GetUserId(c)

	var links []models.Link

	database.DB.Find(&links, models.Link{
		UserId: id,
	})

	var result []interface{}
	var orders []models.Order

	for _, link := range links {
		database.DB.Preload("OrderItems").Find(&orders, &models.Order{
			Code:     link.Code,
			Complete: true,
		})
		revenue := 0.0

		for _, order := range orders {
			revenue += order.GetTotal()
		}

		result = append(result, fiber.Map{
			"code":    link.Code,
			"count":   len(orders),
			"revenue": revenue,
		})
	}

	return c.JSON(result)
}
