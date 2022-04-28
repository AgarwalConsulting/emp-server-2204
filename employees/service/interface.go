package service

import "algogrit.com/emp-server/entities"

type EmployeeService interface {
	Index() ([]entities.Employee, error)
	Create(entities.Employee) (*entities.Employee, error)
}
