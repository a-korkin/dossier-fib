package persons

import (
	"github.com/a-korkin/dossier/internal/models"
	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {
	person := models.Person{
		ID:         1,
		LastName:   "Ivanov",
		FirstName:  "Ivan",
		MiddleName: "Ivanovich",
		Age:        25,
	}

	return c.JSON(&person)
}
