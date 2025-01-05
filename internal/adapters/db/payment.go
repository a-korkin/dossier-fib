package db

import (
	"github.com/a-korkin/dossier/internal/models"
	"log"
)

type PaymentRepo struct {
	DB *PostgresDB
}

func NewPaymentRepo(db *PostgresDB) *PaymentRepo {
	return &PaymentRepo{
		DB: db,
	}
}

func (repo *PaymentRepo) Add(personID uint32, in *models.PaymentDTO) *models.Payment {
	sql := `
insert into public.payments(person_id, sum)
values($1, $2)
returning id, person_id, sum, created;`
	row, err := repo.DB.DB.Query(sql, personID, in.Sum)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	if row.Next() {
		out := models.Payment{}
		err = row.Scan(&out.ID, &out.PersonID, &out.Sum, &out.Created)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		return &out
	} else {
		return nil
	}
}
