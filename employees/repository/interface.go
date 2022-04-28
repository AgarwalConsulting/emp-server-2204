package repository

import "algogrit.com/emp-server/entities"

//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE.go -package=$GOPACKAGE

type EmployeeRepository interface {
	List() ([]entities.Employee, error)
	Create(entities.Employee) (*entities.Employee, error)
}
