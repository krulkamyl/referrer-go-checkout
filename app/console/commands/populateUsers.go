package main

import (
	"referrer/app/database"
	"referrer/app/models"

	"github.com/go-faker/faker/v4"
)

func main() {
	database.Connect()

	for i := 0; i < 30; i++ {
		referrer := models.User{
			FirstName:  faker.FirstName(),
			LastName:   faker.LastName(),
			Email:      faker.Email(),
			IsReferrer: true,
		}

		referrer.SetPassword("1234")

		database.DB.Create(&referrer)
	}
}
