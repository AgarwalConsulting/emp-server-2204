package service_test

import (
	"testing"

	gomock "github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"

	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/employees/service"
	"algogrit.com/emp-server/entities"
)

func TestIndex(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockEmployeeRepository(ctrl)

	sut := service.NewV1(mockRepo)

	expectedEmp := []entities.Employee{{1, "Gaurav", "LnD", 10001}}

	mockRepo.EXPECT().List().Return(expectedEmp, nil)

	actualEmps, err := sut.Index()

	assert.Nil(t, err)
	assert.NotNil(t, actualEmps)

	assert.Equal(t, expectedEmp, actualEmps)
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockEmployeeRepository(ctrl)

	sut := service.NewV1(mockRepo)

	expectedNewEmp := entities.Employee{Name: "Gaurav", Department: "LnD", ProjectID: 10001}
	expectedEmp := &entities.Employee{1, "Gaurav", "LnD", 10001}

	mockRepo.EXPECT().Create(expectedNewEmp).Return(expectedEmp, nil)

	actualEmps, err := sut.Create(expectedNewEmp)

	assert.Nil(t, err)
	assert.NotNil(t, actualEmps)
	assert.Equal(t, expectedEmp, actualEmps)
}
