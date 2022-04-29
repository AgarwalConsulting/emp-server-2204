package service

import "algogrit.com/emp-server/entities"

//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE.go -package=$GOPACKAGE

type EmployeeService interface {
	Index() ([]entities.Employee, error)
	Create(entities.Employee) (*entities.Employee, error)
}
