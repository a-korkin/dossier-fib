package ports

import "github.com/a-korkin/dossier/internal/models"

type Repo interface {
	GetAll() []*models.Person
	GetByID(id uint32) *models.Person
}
