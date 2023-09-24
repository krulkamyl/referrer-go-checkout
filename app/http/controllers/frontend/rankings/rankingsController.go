package rankings

import (
	"context"
	"referrer/app/config"
	"referrer/app/database"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	rankings, err := database.Cache.ZRevRangeByScoreWithScores(context.Background(), config.Getenv("REDIS_RANKING_KEY"), &redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()

	if err != nil {
		panic(err)
	}

	result := make(map[string]float64)

	for _, ranking := range rankings {
		result[ranking.Member.(string)] = ranking.Score
	}

	return c.JSON(result)
}
