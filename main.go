package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type employee struct {
	ID         string `json: "id`
	Name       string `json: "name"`
	Department string `json: "department`
	RM         *employee
}

var employees []employee

func getAllEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside get all emp")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, employee := range employees {
		if employee.ID == params["id"] {
			json.NewEncoder(w).Encode(employee)
			return
		}
	}
	json.NewEncoder(w).Encode(&employee{})
}

func postEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newEmployee employee
	_ = json.NewDecoder(r.Body).Decode(&newEmployee)

	employees = append(employees, newEmployee)
	json.NewEncoder(w).Encode(newEmployee)
}

func putEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, emp := range employees {
		if emp.ID == params["id"] {
			employees = append(employees[:index], employees[index+1:]...)
			var newEmployee employee
			_ = json.NewDecoder(r.Body).Decode(&newEmployee)

			employees = append(employees, newEmployee)
			json.NewEncoder(w).Encode(newEmployee)
			return
		}
	}
	json.NewEncoder(w).Encode(employees)
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, employee := range employees {
		if employee.ID == params["id"] {
			employees = append(employees[:index], employees[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(employees)
}

func main() {
	r := mux.NewRouter()

	employees = append(employees, employee{ID: "1", Name: "ABC", Department: "A1", RM: &employee{ID: "100", Name: "AAA", Department: "A1", RM: nil}})
	employees = append(employees, employee{ID: "2", Name: "DEF", Department: "A1", RM: &employee{ID: "100", Name: "AAA", Department: "A1", RM: nil}})
	employees = append(employees, employee{ID: "3", Name: "GHI", Department: "A1", RM: &employee{ID: "100", Name: "AAA", Department: "A1", RM: nil}})

	// r.HandleFunc("/register", registerEmployee).Method("POST")
	// r.HandleFunc("/login", loginEmployee).Method("POST")
	r.HandleFunc("/employees", getAllEmployees).Methods("GET")
	r.HandleFunc("/employees/{id}", getEmployee).Methods("GET")
	r.HandleFunc("/employees", postEmployee).Methods("POST")
	r.HandleFunc("/employees/{id}", putEmployee).Methods("PUT")
	r.HandleFunc("/employees/{id}", deleteEmployee).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
