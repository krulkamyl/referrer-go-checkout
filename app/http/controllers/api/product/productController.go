package product

import (
	"referrer/app/config"
	"referrer/app/database"
	"referrer/app/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	var products []models.Product

	database.DB.Find(&products)

	return c.JSON(products)
}

func Show(c *fiber.Ctx) error {
	var product models.Product

	id, _ := strconv.Atoi(c.Params("id"))

	product.Id = uint(id)

	database.DB.Find(&product)

	return c.JSON(product)
}

func Store(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Create(&product)
	go database.ClearCache(config.Getenv("REDIS_PRODUCTS_KEY"))

	return c.JSON(product)
}

func Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{}
	product.Id = uint(id)

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Model(&product).Updates(&product)

	go database.ClearCache(config.Getenv("REDIS_PRODUCTS_KEY"))

	return c.JSON(product)
}

func Destroy(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{}
	product.Id = uint(id)

	database.DB.Delete(&product)
	go database.ClearCache(config.Getenv("REDIS_PRODUCTS_KEY"))

	return nil
}
