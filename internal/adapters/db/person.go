package db

import (
	"database/sql"
	"github.com/a-korkin/dossier/internal/models"
	"log"
)

type PersonRepo struct {
	DB *sql.DB
}

func NewPersonRepo(db *sql.DB) *PersonRepo {
	return &PersonRepo{
		DB: db,
	}
}

func (repo *PersonRepo) Add(in *models.PersonDTO) *models.Person {
	sql := `
insert into public.persons(last_name, first_name, middle_name, age)
values($1, $2, $3, $4)
returning id;`
	row, err := repo.DB.Query(
		sql, in.LastName, in.FirstName, in.MiddleName, in.Age)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	out := models.Person{
		LastName:   in.LastName,
		FirstName:  in.FirstName,
		MiddleName: in.MiddleName,
		Age:        in.Age,
	}
	if row.Next() {
		err = row.Scan(&out.ID)
		if err != nil {
			log.Fatal(err)
			return nil
		}
	} else {
		return nil
	}

	return &out
}

func (repo *PersonRepo) Update(id uint32, in *models.PersonDTO) *models.Person {
	sql := `
update public.persons
set last_name = $2,
	first_name = $3,
	middle_name = $4,
	age = $5
where id = $1
returning id, last_name, first_name, middle_name, age;`
	row, err := repo.DB.Query(
		sql, id, in.LastName, in.FirstName, in.MiddleName, in.Age)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	out := models.Person{}
	if row.Next() {
		err = row.Scan(
			&out.ID, &out.LastName, &out.FirstName, &out.MiddleName, &out.Age)
		if err != nil {
			log.Fatal(err)
			return nil
		}
	} else {
		return nil
	}

	return &out
}

func (repo *PersonRepo) GetAll() []*models.Person {
	sql := `
select id, last_name, first_name, middle_name, age 
from public.persons;`
	rows, err := repo.DB.Query(sql)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	persons := make([]*models.Person, 0)
	for rows.Next() {
		person := models.Person{}
		err = rows.Scan(
			&person.ID, &person.LastName, &person.FirstName,
			&person.MiddleName, &person.Age)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		persons = append(persons, &person)
	}
	return persons
}

func (repo *PersonRepo) getPayments(personID uint32) []*models.Payment {
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

func (repo *PersonRepo) GetByID(id uint32) *models.Person {
	sql := `
select id, last_name, first_name, middle_name, age
from public.persons
where id = $1`
	row, err := repo.DB.Query(sql, id)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	person := models.Person{}
	if row.Next() {
		err = row.Scan(
			&person.ID, &person.LastName,
			&person.FirstName, &person.MiddleName, &person.Age)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		person.Payments = repo.getPayments(id)
	} else {
		return nil
	}
	return &person
}

func (repo *PersonRepo) Delete(id uint32) {
	sql := `
delete from public.persons
where id = $1`
	_, err := repo.DB.Exec(sql, id)
	if err != nil {
		log.Fatal(err)
	}
}
