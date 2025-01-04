package persons

import (
	"fmt"

	"github.com/a-korkin/dossier/internal/ports"
	"github.com/gofiber/fiber/v2"
)

func getRepo(c *fiber.Ctx) (ports.Repo, error) {
	personRepo, ok := c.Locals("personRepo").(ports.Repo)
	if !ok {
		return nil, fmt.Errorf("failed to get person repo")
	}
	return personRepo, nil
}

func GetAll(c *fiber.Ctx) error {
	repo, err := getRepo(c)
	if err != nil {
		return err
	}
	return c.JSON(repo.GetAll())
}

func GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	repo, err := getRepo(c)
	if err != nil {
		return err
	}
	person := repo.GetByID(uint32(id))
	if person == nil {
		c.Status(fiber.StatusNotFound)
		return nil
	}
	return c.JSON(person)
}
