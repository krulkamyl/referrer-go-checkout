package user

import (
	"referrer/app/database"
	"referrer/app/http/middlewares"
	"referrer/app/models"

	"strconv"

	"github.com/go-faker/faker/v4"
	"github.com/gofiber/fiber/v2"
)

type CreateLinkRequest struct {
	Products []int
}

func LinkIndex(c *fiber.Ctx) error {
	var links []models.Link

	id, _ := strconv.Atoi(c.Params("id"))

	database.DB.Where("user_id = ?", id).Find(&links)

	for i, link := range links {
		var orders []models.Order
		database.DB.Where("code = ? and complete = true", link.Code).Find(&orders)

		links[i].Orders = orders
	}

	return c.JSON(links)
}

func LinkStore(c *fiber.Ctx) error {
	var request CreateLinkRequest

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	id, _ := middlewares.GetUserId(c)

	link := models.Link{
		UserId: id,
		Code:   faker.Username(),
	}

	for _, productId := range request.Products {
		product := models.Product{}
		product.Id = uint(productId)
		link.Products = append(link.Products, product)
	}

	database.DB.Create(&link)

	return c.JSON(link)
}

func LinkShow(c *fiber.Ctx) error {
	code := c.Params("code")

	link := models.Link{
		Code: code,
	}

	database.DB.Preload("User").Preload("Products").First(&link)

	return c.JSON(link)
}
