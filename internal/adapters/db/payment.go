package db

import (
	"database/sql"
	"github.com/a-korkin/dossier/internal/models"
	"log"
)

type PaymentRepo struct {
	DB *sql.DB
}

func NewPaymentRepo(db *sql.DB) *PaymentRepo {
	return &PaymentRepo{
		DB: db,
	}
}

func (repo *PaymentRepo) Add(personID uint32, in *models.PaymentDTO) *models.Payment {
	sql := `
insert into public.payments(person_id, sum)
values($1, $2)
returning id, person_id, sum, created;`
	row, err := repo.DB.Query(sql, personID, in.Sum)
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

func (repo *PaymentRepo) GetByPerson(personID uint32) []*models.Payment {
	sql := `
select id, person_id, sum, created 
from public.payments
where person_id = $1`
	rows, err := repo.DB.Query(sql, personID)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	payments := make([]*models.Payment, 0)
	for rows.Next() {
		payment := models.Payment{}
		err = rows.Scan(&payment.ID, &payment.PersonID,
			&payment.Sum, &payment.Created)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		payments = append(payments, &payment)
	}
	return payments
}
