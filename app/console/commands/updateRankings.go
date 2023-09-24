package main

import (
	"context"
	"referrer/app/config"
	"referrer/app/database"
	"referrer/app/models"

	"github.com/go-redis/redis/v8"
)

func main() {
	database.Connect()
	database.SetupRedis()

	ctx := context.Background()

	var users []models.User

	database.DB.Find(&users, models.User{
		IsReferrer: true,
	})

	for _, user := range users {
		referrer := models.Referrer(user)
		referrer.CalculateRevenue(database.DB)

		database.Cache.ZAdd(ctx, config.Getenv("REDIS_RANKING_KEY"), &redis.Z{
			Score:  *referrer.Revenue,
			Member: user.Name(),
		})
	}
}
