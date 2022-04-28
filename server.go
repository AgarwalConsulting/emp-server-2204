package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Employee struct {
	ID         int    `json:"-"`
	Name       string `json:"name"`
	Department string `json:"speciality"`
	ProjectID  int    `json:"project"`
}

// MarshalJSON is a receiver functions => methods
// func (e Employee) MarshalJSON() ([]byte, error) {
// 	jsonString := fmt.Sprintf(`{"name": "%s", "speciality": "%s", "project": %d}`, e.Name, e.Department, e.ProjectID)

// 	return []byte(jsonString), nil
// }

var employees = []Employee{
	{1, "Gaurav", "LnD", 1001},
	{2, "Sai Teja", "DBA", 10001},
	{3, "Akshay", "SRE", 10002},
}

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// jsonData, _ := json.Marshal(employees)
	// w.Write(jsonData)

	json.NewEncoder(w).Encode(employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	var newEmployee Employee
	err := json.NewDecoder(req.Body).Decode(&newEmployee)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	newEmployee.ID = len(employees) + 1
	employees = append(employees, newEmployee)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(newEmployee)
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

	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r))
	// http.ListenAndServe(":8000", LoggingMiddleWare(r))
}
