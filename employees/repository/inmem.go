package repository

import "algogrit.com/emp-server/entities"

type inmemRepo struct {
	employees []entities.Employee
}

func (repo *inmemRepo) List() ([]entities.Employee, error) {
	return repo.employees, nil
}

func (repo *inmemRepo) Create(newEmployee entities.Employee) (*entities.Employee, error) {
	newEmployee.ID = len(repo.employees) + 1
	repo.employees = append(repo.employees, newEmployee)

	return &newEmployee, nil
}

func NewInMemRepository() EmployeeRepository {
	repo := inmemRepo{}

	repo.employees = []entities.Employee{
		{1, "Gaurav", "LnD", 1001},
		{2, "Sai Teja", "DBA", 10001},
		{3, "Akshay", "SRE", 10002},
	}

	return &repo
}
