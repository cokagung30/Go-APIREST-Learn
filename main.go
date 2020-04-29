package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Employee struct {
	Employee_id    int    `form:"employee_id" json:"employee_id"`
	Employee_name  string `form:"employee_name" json:"employee_name"`
	Employee_email string `form:"employee_email" json:"employee_email"`
}

type ResponseEmployee struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Employee
}

func returnAllEmployee(w http.ResponseWriter, r *http.Request) {
	var employee Employee
	var arrEmployee []Employee
	var responseProd ResponseEmployee
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/golang_api_learn")

	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

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

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employee", returnAllEmployee).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
