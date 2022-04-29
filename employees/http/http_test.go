package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	empHTTP "algogrit.com/emp-server/employees/http"
	"algogrit.com/emp-server/employees/service"
	"algogrit.com/emp-server/entities"
)

func TestCreateV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service.NewMockEmployeeService(ctrl)
	sut := empHTTP.New(mockSvc)

	jsonBody := `{"name": "Gaurav", "speciality": "LnD", "project": 1001}`
	reqBody := strings.NewReader(jsonBody)
	req := httptest.NewRequest("POST", "/v1/employees", reqBody)

	resRec := httptest.NewRecorder()

	expectedEmp := entities.Employee{0, "Gaurav", "LnD", 1001}
	mockSvc.EXPECT().Create(expectedEmp).Return(&expectedEmp, nil)

	// sut.createV1(resRec, req)
	sut.ServeHTTP(resRec, req)

	res := resRec.Result()

	var actualEmp entities.Employee
	json.NewDecoder(res.Body).Decode(&actualEmp)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, expectedEmp, actualEmp)
}

func FuzzCreateV1(f *testing.F) {
	jsonBody := `{"name": "Gaurav", "speciality": "LnD}`
	f.Add(jsonBody)

	f.Fuzz(func(t *testing.T, in string) {
		t.Log("Input:", in)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockSvc := service.NewMockEmployeeService(ctrl)
		sut := empHTTP.New(mockSvc)

		reqBody := strings.NewReader(in)
		req := httptest.NewRequest("POST", "/v1/employees", reqBody)

		resRec := httptest.NewRecorder()

		expectedEmp := entities.Employee{0, "Gaurav", "LnD", 1001}
		mockSvc.EXPECT().Create(gomock.Any()).AnyTimes().Return(&expectedEmp, nil)

		// sut.createV1(resRec, req)
		sut.ServeHTTP(resRec, req)

		res := resRec.Result()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}
