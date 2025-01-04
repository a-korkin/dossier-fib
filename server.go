package main

import (
	"github.com/a-korkin/dossier/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/persons", persons.GetAll)
	app.Get("/persons/:id", persons.GetByID)

	app.Listen(":8000")
}
