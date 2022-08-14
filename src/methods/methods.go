package methods

import (
	"context"
	"encoding/json"
	"github.com/joho/godotenv"
	"golang-training-repo/employee"
	"golang-training-repo/sql"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Response struct {
	Ok      bool
	Message string
}

func CreateEmployee(c context.Context, e *employee.Employee) (*employee.Employee, error) {

	godotenv.Load()

	dbName, host, port, username, password := os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD")
	var repository employee.Repository = sql.NewDB(dbName, host, port, username, password)

	repository.InsertEmployee(c, e)

	return e, nil
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {

	godotenv.Load()

	dbName, host, port, username, password := os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD")
	var repository employee.Repository = sql.NewDB(dbName, host, port, username, password)

	newEmployee := new(employee.Employee)

	c := r.Context()
	repository.InsertEmployee(c, newEmployee)

	res := Response{Message: "Employee updated"}
	json.NewEncoder(w).Encode(res)
}

func GetEmployeeById(w http.ResponseWriter, r *http.Request) {

	godotenv.Load()

	dbName, host, port, username, password := os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD")
	var repository employee.Repository = sql.NewDB(dbName, host, port, username, password)

	idRaw := r.URL.Query().Get("id")

	parsedId, err := strconv.ParseInt(idRaw, 10, 32)
	if err != nil {
		log.Fatalf("There was an error")
	}

	c := r.Context()
	employees, err := repository.HandleEmployee(c, int(parsedId))
	if err != nil {
		log.Fatalf("There was an error")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func GetEmployeesByPosition(w http.ResponseWriter, r *http.Request) {

	godotenv.Load()

	dbName, host, port, username, password := os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD")
	var repository employee.Repository = sql.NewDB(dbName, host, port, username, password)

	positionRaw := r.URL.Query().Get("position")

	parsedPosition, err := strconv.ParseInt(positionRaw, 10, 32)
	if err != nil {
		log.Fatalf("There was an error")
	}

	c := r.Context()
	employees, err := repository.HandleEmployees(c, employee.Position(parsedPosition))
	if err != nil {
		log.Fatalf("There was an error")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}
