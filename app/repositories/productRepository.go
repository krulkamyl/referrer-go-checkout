package repositories

import (
	"context"
	"encoding/json"
	"sort"
	"strings"
	"time"

	"referrer/app/config"
	"referrer/app/database"
	"referrer/app/models"

	"github.com/gofiber/fiber/v2"
)

func FindProducts(c *fiber.Ctx, searchQuery, sortParam string, page, perPage int) error {

	var ctx = context.Background()
	// todo: convert to search in database and cache by search term
	// todo: convert sort to database and cache by sort term

	var products []models.Product

	result, err := database.Cache.Get(ctx, config.Getenv("REDIS_PRODUCTS_KEY")).Result()
	if err != nil {
		err = database.DB.Find(&products).Error
		if err != nil {
			return nil
		}

		bytes, err := json.Marshal(products)
		if err != nil {
			return nil
		}

		database.Cache.Set(ctx, config.Getenv("REDIS_PRODUCTS_KEY"), bytes, 30*time.Minute).Err()
	} else {
		json.Unmarshal([]byte(result), &products)
	}

	if searchQuery != "" {
		var searchedProducts []models.Product
		lower := strings.ToLower(searchQuery)
		for _, product := range products {
			if strings.Contains(strings.ToLower(product.Title), lower) || strings.Contains(strings.ToLower(product.Description), lower) {
				searchedProducts = append(searchedProducts, product)
			}
		}
		products = searchedProducts
	}

	if sortParam != "" {
		sortLower := strings.ToLower(sortParam)
		if sortLower == "asc" {
			sort.Slice(products, func(i, j int) bool {
				return products[i].Price < products[j].Price
			})
		} else if sortLower == "desc" {
			sort.Slice(products, func(i, j int) bool {
				return products[i].Price > products[j].Price
			})
		}
	}

	startIdx := (page - 1) * perPage
	endIdx := startIdx + perPage
	lastPage := (len(products) / perPage) + 1

	if endIdx > len(products) {
		endIdx = len(products)
	}

	if page > lastPage {
		return nil
	}

	paginatedProducts := products[startIdx:endIdx]

	return c.JSON(fiber.Map{
		"data":      paginatedProducts,
		"total":     len(products),
		"page":      page,
		"last_page": lastPage,
	})
}
