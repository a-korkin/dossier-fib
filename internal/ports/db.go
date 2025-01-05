package ports

import "github.com/a-korkin/dossier/internal/models"

type Repo interface {
	Add(*models.PersonDTO) *models.Person
	Update(uint32, *models.PersonDTO) *models.Person
	GetAll() []*models.Person
	GetByID(id uint32) *models.Person
	Delete(uint32)
}

type PaymentRepo interface {
	Add(uint32, *models.PaymentDTO) *models.Payment
}
