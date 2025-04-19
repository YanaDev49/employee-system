package server

import (
	"net/http"
	"fmt"
	"system/handlers"
)

func Handlers() {
	http.HandleFunc("/add/employee", handlers.AddEmployees)
	http.HandleFunc("/add/department", handlers.AddDepartment)
	http.HandleFunc("/return/employees", handlers.ReturnEmployees)
	http.HandleFunc("/return/department", handlers.ReturnDepartment)
}

func ServerRun() {

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}