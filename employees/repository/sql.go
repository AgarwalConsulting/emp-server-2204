package repository

import (
	"database/sql"
	"log"

	"algogrit.com/emp-server/entities"
	_ "github.com/lib/pq"
)

type sqlRepo struct {
	db *sql.DB
}

func (repo sqlRepo) List() ([]entities.Employee, error) {
	rows, err := repo.db.Query("SELECT * FROM employees")

	if err != nil {
		return nil, err
	}

	emps := []entities.Employee{}

	for rows.Next() {
		var emp entities.Employee

		err := rows.Scan(&emp.ID, &emp.Name, &emp.Department, &emp.ProjectID)
		if err != nil {
			return nil, err
		}

		emps = append(emps, emp)
	}

	return emps, nil
}

func (repo sqlRepo) Create(newEmp entities.Employee) (*entities.Employee, error) {
	rows, err := repo.db.Query("SELECT max(id) FROM employees")

	if err != nil {
		return nil, err
	}

	var maxID *int // Default Value: <nil>

	rows.Next()
	err = rows.Scan(&maxID)
	if err != nil {
		return nil, err
	}

	newEmp.ID = 1
	if maxID != nil {
		newEmp.ID = *maxID + 1
	}

	_, err = repo.db.Exec("INSERT INTO employees (id, name, department, project) VALUES ($1, $2, $3, $4)", newEmp.ID, newEmp.Name, newEmp.Department, newEmp.ProjectID)

	if err != nil {
		return nil, err
	}

	return &newEmp, nil
}

func NewSQLRepository(dbURL string) EmployeeRepository {
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatalln("Unable to connect:", err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS employees (id numeric primary key, name text, department text, project numeric)")

	if err != nil {
		log.Fatalln("Unable to create table:", err)
	}

	repo := sqlRepo{db}

	return &repo
}
