package service

import (
	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/entities"
)

type empSvcV1 struct {
	repo repository.EmployeeRepository
}

func (svc empSvcV1) Index() ([]entities.Employee, error) {
	return svc.repo.List()
}

func (svc empSvcV1) Create(newEmployee entities.Employee) (*entities.Employee, error) {
	return svc.repo.Create(newEmployee)
}

func NewV1(repo repository.EmployeeRepository) EmployeeService {
	svc := empSvcV1{repo}

	return svc
}
