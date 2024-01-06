package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Membuat instance dari Fiber
	app := fiber.New()

	// Mengatur endpoint untuk route GET "/"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Mengatur endpoint untuk route GET "/hello/:name"
	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.SendString("Hello, " + name + "!")
	})

	// Menjalankan aplikasi pada port 3000
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
