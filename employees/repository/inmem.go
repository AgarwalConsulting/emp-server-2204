package repository

import (
	"sync"

	"algogrit.com/emp-server/entities"
)

type inmemRepo struct {
	employees []entities.Employee
	mut       sync.RWMutex
}

func (repo *inmemRepo) List() ([]entities.Employee, error) {
	repo.mut.RLock()
	defer repo.mut.RUnlock()

	return repo.employees, nil
}

func (repo *inmemRepo) Create(newEmployee entities.Employee) (*entities.Employee, error) {
	repo.mut.Lock()
	defer repo.mut.Unlock()

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
