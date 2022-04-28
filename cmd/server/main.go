package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/employees/service"
	"algogrit.com/emp-server/entities"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// MarshalJSON is a receiver functions => methods
// func (e Employee) MarshalJSON() ([]byte, error) {
// 	jsonString := fmt.Sprintf(`{"name": "%s", "speciality": "%s", "project": %d}`, e.Name, e.Department, e.ProjectID)

// 	return []byte(jsonString), nil
// }

var repo = repository.NewInMemRepository()
var svcV1 = service.NewV1(repo)

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// jsonData, _ := json.Marshal(employees)
	// w.Write(jsonData)

	employees, err := svcV1.Index()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	json.NewEncoder(w).Encode(employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	var newEmployee entities.Employee
	err := json.NewDecoder(req.Body).Decode(&newEmployee)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	createdEmp, err := svcV1.Create(newEmployee)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(createdEmp)
}

// func EmployeesHandler(w http.ResponseWriter, req *http.Request) {
// 	if req.Method == "POST" {
// 		EmployeeCreateHandler(w, req)
// 	} else {
// 		EmployeesIndexHandler(w, req)
// 	}
// }

func LoggingMiddleWare(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		begin := time.Now()

		h.ServeHTTP(w, req)

		log.Printf("%s %s took %s\n", req.Method, req.URL, time.Since(begin))
	}
}

func main() {
	r := mux.NewRouter()
	// r := http.NewServeMux()
	// r := http.DefaultServeMux

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!" // Type: string

		// w.Write([]byte(msg))
		fmt.Fprintln(w, msg)
	})

	r.HandleFunc("/employees", EmployeesIndexHandler).Methods("GET")
	r.HandleFunc("/employees", EmployeeCreateHandler).Methods("POST")
	// r.HandleFunc("/employees", EmployeesHandler)
	// r.HandleFunc("/employees/{id}", EmployeeShowHandler)

	// r.HandleFunc("/employees", EmployeesIndexHandler)
	// r.HandleFunc("/employees", EmployeeCreateHandler)

	log.Println("Starting server on port: 8000...")
	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r))
	// http.ListenAndServe(":8000", LoggingMiddleWare(r))
}
