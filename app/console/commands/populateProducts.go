package main

import (
	"referrer/app/database"
	"referrer/app/models"

	"math/rand"

	"github.com/go-faker/faker/v4"
)

func main() {
	database.Connect()

	for i := 0; i < 30; i++ {
		product := models.Product{
			Title:       faker.Username(),
			Description: faker.Username(),
			Image:       faker.URL(),
			Price:       rand.Float64() * 100,
		}

		database.DB.Create(&product)
	}
}
