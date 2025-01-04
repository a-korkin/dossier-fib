package main

import (
	"github.com/gofiber/fiber/v2"
)

type Person struct {
	ID         int    `json:"id"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	Age        uint8  `json:"age"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from fiber!")
	})

	app.Get("/persons", func(c *fiber.Ctx) error {
		person := &Person{
			ID:         1,
			LastName:   "Bickle",
			FirstName:  "Travis",
			MiddleName: "Ivanovich",
			Age:        37,
		}
		return c.JSON(person)
	})

	app.Listen(":8000")
}
