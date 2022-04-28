package repository

import "algogrit.com/emp-server/entities"

type EmployeeRepository interface {
	List() ([]entities.Employee, error)
	Create(entities.Employee) (*entities.Employee, error)
}
