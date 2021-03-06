package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"algogrit.com/emp-server/entities"
	"github.com/go-playground/validator/v10"
)

func (h Handler) indexV1(w http.ResponseWriter, req *http.Request) {
	employees, err := h.svcV1.Index()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	json.NewEncoder(w).Encode(employees)
}

func (h Handler) createV1(w http.ResponseWriter, req *http.Request) {
	var newEmployee entities.Employee
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&newEmployee)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(newEmployee)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	createdEmp, err := h.svcV1.Create(newEmployee)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(createdEmp)
}
