package main

// import "github.com/jinzhu/gorm"

type Employee struct {
	EmpId     int     `json:"empid" gorm:"primary_key"`
	// gorm.Model // struct with ID as primary Key
	EmpName   string  `json:"empname"`
	EmpSalary float64 `json:"salary"`
	Email     string  `json:"email"`
	Gender    string  `json:"gender"`
}
