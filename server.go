package main

import (
	"github.com/a-korkin/dossier/configs"
	"github.com/a-korkin/dossier/internal/adapters/db"
	"github.com/a-korkin/dossier/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	dbConn := configs.GetEnv("GOOSE_DBSTRING")
	pg, err := db.NewPostgresDB(dbConn)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = pg.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	app := fiber.New()

	personRepo := db.NewPersonRepo(pg)
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("personRepo", personRepo)
		return c.Next()
	})

	app.Post("/persons", persons.Add)
	app.Put("/persons/:id", persons.Update)
	app.Get("/persons", persons.GetAll)
	app.Get("/persons/:id", persons.GetByID)
	app.Delete("/persons/:id", persons.Delete)

	if err = app.Listen(":8000"); err != nil {
		log.Fatal(err)
	}
}
