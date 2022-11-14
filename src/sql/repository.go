package sql

import (
	"context"
	"database/sql"
	"fmt"
	"golang-training-repo/employee"
)

type Database struct {
	dbName   string
	host     string
	port     string
	username string
	password string
}

func NewDB(dbName, host, port, username, password string) *Database {
	return &Database{
		dbName:   dbName,
		host:     host,
		port:     port,
		username: username,
		password: password,
	}
}

func (r *Database) ConnectionString() (string, error) {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", r.username, r.password, r.host, r.port, r.dbName), nil
}

func (r *Database) InsertEmployee(c context.Context, e *employee.Employee) error {
	connectionString, err := r.ConnectionString()
	if err != nil {
		return err
	}

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO employee(full_name, `position`, salary, joined, on_probation) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (r *Database) HandleEmployeeByPosition(c context.Context, position employee.Position) ([]employee.Employee, error) {
	res := make([]employee.Employee, 0)

	connectionString, err := r.ConnectionString()

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query(
		"SELECT * FROM employee e WHERE e.`position`=?", position)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		nextEmployee := new(employee.Employee)
		err := rows.Scan(&nextEmployee.ID, &nextEmployee.FullName, &nextEmployee.Position, &nextEmployee.Salary, &nextEmployee.Joined, &nextEmployee.OnProbation, &nextEmployee.CreatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, *nextEmployee)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *Database) HandleEmployeeById(c context.Context, id int) (*employee.Employee, error) {
	var res *employee.Employee

	connectionString, err := r.ConnectionString()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(
		"SELECT * FROM employee WHERE e.id=?", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		nextEmployee := new(employee.Employee)

		err := rows.Scan(&nextEmployee.ID, &nextEmployee.FullName, &nextEmployee.Position, &nextEmployee.Salary, &nextEmployee.Joined, &nextEmployee.OnProbation, &nextEmployee.CreatedAt)
		if err != nil {
			return nil, err
		}
		res = nextEmployee
	}

	return res, nil
}
