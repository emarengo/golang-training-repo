package methods

import (
	"context"
	"errors"
	"github.com/joho/godotenv"
	"golang-training-repo/employee"
	"golang-training-repo/sql"
	"os"
)

type Response struct {
	Ok      bool
	Message string
}

func CreateEmployee(c context.Context, e *employee.Employee) (*employee.Employee, error) {
	if e.FullName == "" {
		return nil, errors.New("full name cannot be empty")
	}
	if e.Salary == 0 {
		return nil, errors.New("salary cannot be 0")
	}
	godotenv.Load()

	dbName, host, port, username, password := os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD")
	var repository employee.Repository = sql.NewDB(dbName, host, port, username, password)

	repository.InsertEmployee(c, e)

	return e, nil
}

func UpdateEmployee(c context.Context, e *employee.Employee) (*employee.Employee, error) {
	if e.FullName == "" {
		return nil, errors.New("full name cannot be empty")
	}
	if e.Position == employee.Undefined {
		return nil, errors.New("position cannot be undefined")
	}
	godotenv.Load()

	dbName, host, port, username, password := os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD")
	var repository employee.Repository = sql.NewDB(dbName, host, port, username, password)

	repository.InsertEmployee(c, e)

	return e, nil
}

func GetEmployeeById(c context.Context, id int) (*employee.Employee, error) {

	godotenv.Load()

	dbName, host, port, username, password := os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD")
	var repository employee.Repository = sql.NewDB(dbName, host, port, username, password)

	employees, err := repository.HandleEmployeeById(c, id)
	if err != nil {
		return nil, errors.New("employee not found")
	}
	return employees, nil
}

func GetEmployeesByPosition(c context.Context, position int) ([]employee.Employee, error) {

	godotenv.Load()

	dbName, host, port, username, password := os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD")
	var repository employee.Repository = sql.NewDB(dbName, host, port, username, password)

	employees, err := repository.HandleEmployeeByPosition(c, employee.Position(position))
	if err != nil {
		return nil, errors.New("employee not found")
	}

	return employees, nil
}
