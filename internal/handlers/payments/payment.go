package payments

import (
	"fmt"
	"github.com/a-korkin/dossier/internal/models"
	"github.com/a-korkin/dossier/internal/ports"
	"github.com/gofiber/fiber/v2"
)

func getRepo(c *fiber.Ctx) (ports.PaymentRepo, error) {
	repo, ok := c.Locals("paymentRepo").(ports.PaymentRepo)
	if !ok {
		return nil, fmt.Errorf("failed to payment repo")
	}

	return repo, nil
}

func Add(c *fiber.Ctx) error {
	repo, err := getRepo(c)
	if err != nil {
		return err
	}
	personID, err := c.ParamsInt("person_id")
	if err != nil {
		return err
	}
	in := models.PaymentDTO{}
	if err = c.BodyParser(&in); err != nil {
		return err
	}
	return c.JSON(repo.Add(uint32(personID), &in))
}
