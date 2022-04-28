package http

import (
	"algogrit.com/emp-server/employees/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	*mux.Router
	svcV1 service.EmployeeService
	// Router *mux.Router
}

func (h *Handler) SetupRoutes(r *mux.Router) {
	r.HandleFunc("/v1/employees", h.indexV1).Methods("GET")
	r.HandleFunc("/v1/employees", h.createV1).Methods("POST")

	h.Router = r
}

func New(svcV1 service.EmployeeService) Handler {
	h := Handler{svcV1: svcV1}

	h.SetupRoutes(mux.NewRouter())

	return h
}
