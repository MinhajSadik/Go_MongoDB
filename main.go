package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("value:" + c.Params("value"))
	})

	// GET http://localhost:8080/Minhaj

	app.Get("/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello " + c.Params("name"))
		}
		return c.SendString("Where Is Minhaj")
	})

	// http:localhost:8080/api/minhajsadik/documents/my-projects/Go-MongoDB

	app.Get("/api/*", func(c *fiber.Ctx) error {
		return c.SendString("API path: " + c.Params("*"))

	})

	app.Listen(":8080")
}
