package main

import (
	"charcha-api/database"
	"charcha-api/routes/productRoutes"
	"charcha-api/routes/userRoutes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectToDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	//user routes
	app.Post("/user", userRoutes.CreateUser)
	app.Get("/user/:id", userRoutes.GetService)
	app.Delete("/user/:id", userRoutes.DeleteService)

	//product routes
	app.Post("/product", productRoutes.CreateOrGetProduct)

	app.Listen(":3000")

}
