package main

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
