package main

import (
	"github.com/a-korkin/dossier/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from fiber!")
	})

	app.Get("/persons", persons.GetAll)

	app.Listen(":8000")
}
