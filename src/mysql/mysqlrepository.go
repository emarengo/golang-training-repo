package mysql

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang-training-repo/employee"
)

type MySQLRepository struct {
	dbName   string
	host     string
	port     string
	name     string
	password string
}

func SetDB(dbName, host, port, name, password string) *MySQLRepository {
	return &MySQLRepository{
		dbName:   dbName,
		host:     host,
		port:     port,
		name:     name,
		password: password,
	}
}

func (r *MySQLRepository) ConnectionString() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", r.name, r.password, r.host, r.port, r.dbName)
}

func (r *MySQLRepository) InsertEmployee(ctx context.Context, e *employee.Employee) error {
	db, err := sql.Open("mysql", r.ConnectionString())
	if err != nil {
		return err
	}

	defer db.Close()

	prepared, err := db.Prepare("INSERT INTO employee(full_name, `position`, salary, joined, on_probation) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer prepared.Close()

	_, err = prepared.Exec(e.FullName, e.Position, e.Salary, e.Joined, e.OnProbation)
	if err != nil {
		return err
	}

	return nil
}
