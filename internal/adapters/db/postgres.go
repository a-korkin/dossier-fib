package db

import (
	"database/sql"
	"github.com/a-korkin/dossier/internal/models"
	_ "github.com/lib/pq"
	"log"
)

type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresDB(conn string) (*PostgresDB, error) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	pg := PostgresDB{
		DB: db,
	}

	return &pg, nil
}

func (pg *PostgresDB) Close() error {
	if err := pg.DB.Close(); err != nil {
		return err
	}
	return nil
}

type PersonRepo struct {
	DB *PostgresDB
}

func NewPersonRepo(db *PostgresDB) *PersonRepo {
	return &PersonRepo{
		DB: db,
	}
}

func (repo *PersonRepo) GetAll() []*models.Person {
	sql := `
select id, last_name, first_name, middle_name, age 
from public.persons;`
	rows, err := repo.DB.DB.Query(sql)
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

func (repo *PersonRepo) GetByID(id uint32) *models.Person {
	sql := `
select id, last_name, first_name, middle_name, age
from public.persons
where id = $1`
	row, err := repo.DB.DB.Query(sql, id)
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
	} else {
		return nil
	}
	return &person
}
