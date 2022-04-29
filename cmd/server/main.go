package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	employeeHTTP "algogrit.com/emp-server/employees/http"
	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/employees/service"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func LoggingMiddleWare(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		begin := time.Now()

		h.ServeHTTP(w, req)

		log.Printf("%s %s took %s\n", req.Method, req.URL, time.Since(begin))
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!" // Type: string

		// w.Write([]byte(msg))
		fmt.Fprintln(w, msg)
	})

	// var repo = repository.NewInMemRepository()
	var repo = repository.NewSQLRepository()
	var svcV1 = service.NewV1(repo)
	var empHandler = employeeHTTP.New(svcV1)

	empHandler.SetupRoutes(r)

	log.Println("Starting server on port: 8000...")
	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r))
	// http.ListenAndServe(":8000", LoggingMiddleWare(r))
}
