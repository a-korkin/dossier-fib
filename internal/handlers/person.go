package persons

import (
	// "github.com/a-korkin/dossier/internal/adapters/db"
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
	// personRepo := db.PersonRepo{}
	// repo, _ := c.Locals(personRepo).(db.PersonRepo)
	// return c.JSON(repo.GetAll())
	// if ok {
	// 	log.Printf("repo: %v", repo)
	// }
	return c.JSON(&persons)
}

func GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("id: %s", id)

	return c.JSON(&persons[0])
}
