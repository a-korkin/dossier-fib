package persons

import (
	"fmt"

	"github.com/a-korkin/dossier/internal/models"
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

func Add(c *fiber.Ctx) error {
	repo, err := getRepo(c)
	if err != nil {
		return err
	}
	in := models.PersonDto{}
	if err = c.BodyParser(&in); err != nil {
		return err
	}
	return c.JSON(repo.Add(&in))
}

func Update(c *fiber.Ctx) error {
	repo, err := getRepo(c)
	if err != nil {
		return err
	}
	in := models.PersonDto{}
	if err = c.BodyParser(&in); err != nil {
		return err
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	return c.JSON(repo.Update(uint32(id), &in))
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

func Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	repo, err := getRepo(c)
	if err != nil {
		return err
	}
	repo.Delete(uint32(id))
	c.Status(fiber.StatusNoContent)
	return nil
}
