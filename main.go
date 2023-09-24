package main

import (
	"referrer/app/database"
	"referrer/app/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	database.AutoMigrate()
	database.SetupRedis()
	database.SetupCacheChannel()

	app := fiber.New()
	routers.Setup(app)

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Listen(":3000")
}
