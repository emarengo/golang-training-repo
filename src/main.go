package main

import (
	"encoding/json"
	"golang-training-repo/employee"
	"golang-training-repo/methods"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/CreateEmployee", createEmployee)
	mux.HandleFunc("/UpdateEmployee", methods.UpdateEmployee)
	mux.HandleFunc("/GetEmployeeById", methods.GetEmployeeById)
	mux.HandleFunc("/GetEmployeesByPosition", methods.GetEmployeesByPosition)

	server := &http.Server{
		Addr:    ":5050",
		Handler: mux,
	}

	server.ListenAndServe()
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	body := new(employee.Employee)
	json.NewDecoder(r.Body).Decode(&body)

	c := r.Context()

	res, error := methods.CreateEmployee(c, body)
	if error != nil {
		log.Fatal("error")
	}
	json.NewEncoder(w).Encode(res)

}
