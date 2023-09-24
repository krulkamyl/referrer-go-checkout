package products

import (
	"referrer/app/repositories"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	searchQuery := c.Query("s")
	sortParam := c.Query("sort")

	page, _ := strconv.Atoi(c.Query("page", "1"))

	if page < 1 {
		page = 1
	}
	perPage := 9

	products := repositories.FindProducts(c, searchQuery, sortParam, page, perPage)

	if products != nil {
		panic("Cant fetch products")
	}

	return products
}
