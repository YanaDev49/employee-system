package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"system/database"
	"system/models"

)

func AddEmployees(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
    http.Error(w,"Method Not Allowed!", http.StatusMethodNotAllowed)
	return
	}

	var employee models.Employee

	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}



	err = database.Db.Create(&employee).Error; // Always remember to add the .Error;  so it can access the error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	w.Header().Set("content-type", "application/json")

	err = json.NewEncoder(w).Encode(employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)



}

func AddDepartment(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Incorrect method!", http.StatusMethodNotAllowed)
		return
	}


	var department models.Department

	err := json.NewDecoder(r.Body).Decode(&department)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.Db.Create(&department).Error;
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


    w.Header().Set("content-type", "application/json")

	err = json.NewEncoder(w).Encode(department)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	w.WriteHeader(http.StatusOK)

}

func ReturnEmployees(w http.ResponseWriter, r *http.Request) {

 var employee []models.Employee

	if r.Method != http.MethodGet {
		log.Fatal("Incorrect method!", http.StatusMethodNotAllowed)
	}
    
    err := database.Db.Find(&employee).Error;
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	err = json.NewEncoder(w).Encode(employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func ReturnDepartment(w http.ResponseWriter, r *http.Request) {

var department []models.Department

    if r.Method != http.MethodGet {
		http.Error(w, "Incorrect method", http.StatusMethodNotAllowed)
	}

    err := database.Db.Find(&department).Error;
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

    err = json.NewEncoder(w).Encode(department)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}



