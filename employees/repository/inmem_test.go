package repository_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/entities"
)

func TestConsistency(t *testing.T) {
	sut := repository.NewInMemRepository()

	emps, err := sut.List()

	assert.Nil(t, err)
	assert.NotNil(t, emps)

	initialEmpCount := len(emps)

	assert.NotEqual(t, 0, initialEmpCount)

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		newEmployee := entities.Employee{Name: "Gaurav", Department: "LnD", ProjectID: 10001}

		go func() {
			defer wg.Done()
			emp, err := sut.Create(newEmployee)

			assert.Nil(t, err)
			assert.NotNil(t, emp)
			assert.NotEqual(t, 0, emp.ID)
		}()
	}

	wg.Wait()

	emps, err = sut.List()

	assert.Nil(t, err)
	assert.NotNil(t, emps)

	finalEmpCount := len(emps)

	assert.NotEqual(t, 0, finalEmpCount)
	assert.Equal(t, 100+initialEmpCount, finalEmpCount)
}
