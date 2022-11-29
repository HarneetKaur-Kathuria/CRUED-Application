package main

type Employee struct {
	EmpId     int     `json:"empid"`
	EmpName   string  `json:"empname"`
	EmpSalary float64 `json:"salary"`
	Email     string  `json:"email"`
	Gender    string  `json:"gender"`
}
