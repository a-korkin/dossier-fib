package ports

import "github.com/a-korkin/dossier/internal/models"

type Repo interface {
	Add(*models.PersonDto) *models.Person
	Update(uint32, *models.PersonDto) *models.Person
	GetAll() []*models.Person
	GetByID(id uint32) *models.Person
}
