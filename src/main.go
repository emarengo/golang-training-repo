package main

import (
	"golang-training-repo/methods"
	"net/http"
)

const keyServerAddr = "serverAddr"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/CreateEmployee", methods.CreateEmployee)
	mux.HandleFunc("/UpdateEmployee", methods.UpdateEmployee)
	mux.HandleFunc("/GetEmployeeById", methods.GetEmployeeById)
	mux.HandleFunc("/GetEmployeesByPosition", methods.GetEmployeesByPosition)

	server := &http.Server{
		Addr:    ":5050",
		Handler: mux,
	}

	server.ListenAndServe()
}
