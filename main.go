package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func returnAllEmployee(w http.ResponseWriter, r *http.Request) {
	var employee Employee
	var arrEmployee []Employee
	var responseProd ResponseEmployee

	db := connect()
	defer db.Close()

	rows, err := db.Query("Select employee_id, employee_name, employee_email From employee")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&employee.Employee_id, &employee.Employee_name, &employee.Employee_email); err != nil {
			log.Fatal(err.Error())
		} else {
			arrEmployee = append(arrEmployee, employee)
		}
	}

	responseProd.Status = 1
	responseProd.Message = "Get Data Employee Success"
	responseProd.Data = arrEmployee

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseProd)
}

func insertEmployeeMultipart(w http.ResponseWriter, r *http.Request) {
	var response ResponseEmployee

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	employeeName := r.FormValue("name")
	employeeEmail := r.FormValue("email")
	employeePhone := r.FormValue("phone")
	employeeAddress := r.FormValue("address")

	_, err = db.Exec("Insert into employee (employee_name, employee_email, employee_phone, employee_address) values (?, ?, ?, ?)", employeeName, employeeEmail, employeePhone, employeeAddress)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Data Employee Berhasil Ditambahkan"
	log.Print("Insert Data to Database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func updateEmployeeMultipart(w http.ResponseWriter, r *http.Request) {
	var response ResponseEmployee

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	employeeId := r.FormValue("id")
	employeeName := r.FormValue("name")
	employeeEmail := r.FormValue("email")
	employeePhone := r.FormValue("phone")
	employeeAddress := r.FormValue("address")

	_, err = db.Exec("update employee set employee_name = ?, employee_email = ?, employee_phone = ?, employee_address = ? where employee_id = ?", employeeName, employeeEmail, employeePhone, employeeAddress, employeeId)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Pembaharuan Data Employee Berhasil"
	log.Print("Insert Data to Database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func deleteEmployeeMultipart(w http.ResponseWriter, r *http.Request) {
	var response ResponseEmployee

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	employeeID := r.FormValue("id")

	_, err = db.Exec("DELETE from employee where employee_id = ?",
		employeeID,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Delete Data Employee Success"
	log.Print("Delete data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employee", returnAllEmployee).Methods("GET")
	router.HandleFunc("/employee", insertEmployeeMultipart).Methods("POST")
	router.HandleFunc("/employee", updateEmployeeMultipart).Methods("PUT")
	router.HandleFunc("/employee", deleteEmployeeMultipart).Methods("DELETE")
	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
