package main

import (
	"encoding/json"
	"golang-training-repo/employee"
	"golang-training-repo/methods"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/CreateEmployee", createEmployee)
	mux.HandleFunc("/UpdateEmployee", updateEmployee)
	mux.HandleFunc("/GetEmployeeById", getEmployeeById)
	mux.HandleFunc("/GetEmployeesByPosition", getEmployeesByPosition)

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

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	body := new(employee.Employee)
	json.NewDecoder(r.Body).Decode(&body)

	c := r.Context()

	res, error := methods.UpdateEmployee(c, body)
	if error != nil {
		log.Fatal("error")
	}
	json.NewEncoder(w).Encode(res)

}

func getEmployeeById(w http.ResponseWriter, r *http.Request) {

	rawId := r.URL.Query().Get("id")
	id, error := strconv.ParseInt(rawId, 10, 32)
	if error != nil {
		log.Fatal("error")
	}

	c := r.Context()
	res, error := methods.GetEmployeeById(c, int(id))
	if error != nil {
		log.Fatal("error")
	}
	json.NewEncoder(w).Encode(res)
}

func getEmployeesByPosition(w http.ResponseWriter, r *http.Request) {

	rawPosition := r.URL.Query().Get("position")
	position, error := strconv.ParseInt(rawPosition, 10, 32)
	if error != nil {
		log.Fatal("error")
	}

	c := r.Context()
	res, error := methods.GetEmployeesByPosition(c, int(position))
	if error != nil {
		log.Fatal("error")
	}
	json.NewEncoder(w).Encode(res)
}
