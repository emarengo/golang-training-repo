package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"golang-training-repo/employee"
	"golang-training-repo/mysql"
	"log"
	"os"
	"time"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	var repository employee.Repository = mysql.SetDB(dbName, dbHost, dbPort, dbUsername, dbPassword)
	ctx := context.Background()

	createEmployee := &employee.Employee{
		FullName:    "Ernesto",
		Position:    employee.Senior,
		Salary:      999,
		Joined:      time.Now(),
		OnProbation: false,
	}

	fmt.Println(createEmployee)

	err = repository.InsertEmployee(ctx, createEmployee)
	if err != nil {
		fmt.Println(err)
	}

}
