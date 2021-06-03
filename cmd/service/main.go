package main

import (
	"github.com/forinterns/srv/internal/handlers/get"
	"net/http"
)

func main() {
	addEmployee := add.New()
	//getEmployee := get.New()
	addEmployeeHandler := http.HandlerFunc(addEmployee.Handle)
	//getEmployeeHandler := http.HandlerFunc(getEmployee.Handle)
	http.Handle("/employee/add", addEmployeeHandler)
	//http.Handle("/employee/get", getEmployeeHandler)
	http.ListenAndServe(":8080", nil)
}
