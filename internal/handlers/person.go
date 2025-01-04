package persons

import (
	"github.com/a-korkin/dossier/internal/models"
	"github.com/gofiber/fiber/v2"
	"log"
)

var persons = []models.Person{
	{
		ID:         1,
		LastName:   "Ivanov",
		FirstName:  "Ivan",
		MiddleName: "Ivanovich",
		Age:        35,
	},
	{
		ID:         2,
		LastName:   "Petrov",
		FirstName:  "Petr",
		MiddleName: "Petrovich",
		Age:        23,
	},
}

func GetAll(c *fiber.Ctx) error {
	return c.JSON(&persons)
}

func GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("id: %s", id)

	return c.JSON(&persons[0])
}
