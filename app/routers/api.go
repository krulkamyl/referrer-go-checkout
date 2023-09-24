package routers

import (
	"referrer/app/http/controllers/api/auth"
	"referrer/app/http/controllers/api/order"
	"referrer/app/http/controllers/api/product"
	"referrer/app/http/controllers/api/user"
	"referrer/app/http/controllers/frontend/products"
	"referrer/app/http/controllers/frontend/rankings"
	"referrer/app/http/controllers/frontend/stats"
	"referrer/app/http/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// /api
	api := app.Group("api")

	// /api/admin
	admin := api.Group("admin")
	admin.Post("/login", auth.Login)
	admin.Post("/register", auth.Register)

	adminAuthenticated := admin.Use(middlewares.IsAuthenticated)
	adminAuthenticated.Get("/user", auth.Show)
	adminAuthenticated.Put("/user/info", auth.Update)
	adminAuthenticated.Put("/user/password", auth.UpdatePassword)
	adminAuthenticated.Post("/logout", auth.Logout)

	// /api/admin/users
	usersGroup := adminAuthenticated.Group("users")
	usersGroup.Get("/", user.Index)
	usersGroup.Get("/:id/links", user.LinkIndex)

	// /api/admin/products
	productsGroup := adminAuthenticated.Group("products")
	productsGroup.Get("/", product.Index)
	productsGroup.Post("/store", product.Store)
	productsGroup.Get("/:id", product.Show)
	productsGroup.Put("/:id", product.Update)
	productsGroup.Delete("/:id", product.Destroy)

	// /api/admin/orders
	ordersGroup := adminAuthenticated.Group("orders")
	ordersGroup.Get("/", order.Index)

	// /api/referrer
	referrer := api.Group("referrer")
	referrer.Post("/login", auth.Login)
	referrer.Post("/register", auth.Register)

	refererAuthenticated := referrer.Use(middlewares.IsAuthenticated)
	refererAuthenticated.Get("/user", auth.Show)
	refererAuthenticated.Put("/user/info", auth.Update)
	refererAuthenticated.Put("/user/password", auth.UpdatePassword)
	refererAuthenticated.Post("/logout", auth.Logout)
	refererAuthenticated.Post("/links", user.LinkStore)

	// /api/referrer/frontend
	referrerFrontend := referrer.Group("frontend")
	referrerFrontend.Get("/stats", stats.Index)
	referrerFrontend.Get("/rankings", rankings.Index)

	// /api/referrer/frontend/products
	referrerFrontendProduct := referrerFrontend.Group("products")
	referrerFrontendProduct.Get("/", products.Index)

	// /api/checkout
	checkout := api.Group("checkout")
	checkout.Get("links/:code", user.LinkShow)
	checkout.Post("orders", order.Store)
	checkout.Post("orders/confirm", order.Update)
}
