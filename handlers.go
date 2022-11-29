package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// to check if the Json is Empty or not
func (e *Employee) IsEmpty() bool {
	return e.EmpId == 0 && e.EmpName == ""
}

// To create Employees
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content - Type", "application/json")

	/*

		--- could be double Using Unmarshal

		reqBody, _ := ioutil.ReadAll(r.Body)
		_= json.Unmarshal(reqBody, &newEmp)
	*/
	var newEmp Employee
	_ = json.NewDecoder(r.Body).Decode(&newEmp)

	// checking if the Json is Empty
	if newEmp.IsEmpty() {
		json.NewEncoder(w).Encode("Please Provide the Data")
		return
	}

	var employees []Employee
	DataBase.Find(&employees)

	// checking if the Emp Id Already Exits, If yes the sends the msg
	for _, emp := range employees {
		if emp.EmpId == newEmp.EmpId {
			json.NewEncoder(w).Encode("Employee Id Already Exits")
			return
		}
	}

	// if No then creates the new Enployee
	DataBase.Create(&newEmp)
	json.NewEncoder(w).Encode("Employee Created !")
	json.NewEncoder(w).Encode(newEmp)

}

// to get all the Employees Details
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content - Type", "application/json")
	var employees []Employee
	DataBase.Find(&employees)
	json.NewEncoder(w).Encode(employees)

}

// To get Employee Details by Id
func GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content - Type", "application/json")
	employees := []Employee{}
	params := mux.Vars(r)
	DataBase.Find(&employees) // transfering data from table to employees var

	// DataBase.First(&emp, mux.Vars(r)["id"])
	// DataBase.First(&emp, params["empid"])

	for _, emp := range employees {
		// string to int
		e_id, err := strconv.Atoi(params["id"])
		if err == nil {
			if emp.EmpId == e_id {
				json.NewEncoder(w).Encode(emp)
				return
			}
		}

	}
	json.NewEncoder(w).Encode("Employee Id Not Found")

}

func DeleteAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content - Type", "application/json")
	var employees []Employee
	DataBase.Delete(&employees)
	json.NewEncoder(w).Encode("Employees Deleted !")

}

func DeleteEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content - Type", "application/json")

	employees := []Employee{}
	params := mux.Vars(r)
	DataBase.Find(&employees) // transfering data from table to employees var

	for _, emp := range employees {
		// string to int
		e_id, err := strconv.Atoi(params["id"])
		if err == nil {
			if emp.EmpId == e_id {
				DataBase.Delete(&emp)
				DataBase.Save(&employees)
				json.NewEncoder(w).Encode("Employee Deleted !")
				return
			}
		}
	}
	json.NewEncoder(w).Encode("Employee Id Not Found")

}

func UpdateById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(" Content Type ", "application/json")

	params := mux.Vars(r)

	var updatedEmp Employee
	// DataBase.First(&updatedEmp, params["id"])
	// if employee.IsEmpty() {
	// 	json.NewEncoder(w).Encode("Employee Id Not Found")
	// 	return
	// }

	// var oldemp Employee
	var employees []Employee
	// var oldemp Employee
	DataBase.Find(&employees)
	_ = json.NewDecoder(r.Body).Decode(&updatedEmp)
	// DataBase.Where("id=?", params["id"]).Delete(&employees)

	// DataBase.First(&oldemp, params["id"])
	// oldemp.Email = updatedEmp.Email
	// oldemp.EmpId = updatedEmp.EmpId
	// oldemp.Gender = updatedEmp.Gender
	// oldemp.EmpName = updatedEmp.EmpName
	// oldemp.EmpSalary = updatedEmp.EmpSalary
	// DataBase.Save(&updatedEmp)
	// DataBase.Create(&employee)
	// json.NewEncoder(w).Encode("Employee Updated !")
	// json.NewEncoder(w).Encode(&employee)

	// var employees []Employee\
	// var oldemp Employee
	// DataBase.Find(&employees)
	// DataBase.Where("emp_id =?", params["id"])
	// fmt.Println(oldemp)
	// json.NewDecoder(r.Body).Decode(&oldemp)
	// DataBase.InstantSet("empid", )
	// fmt.Println(oldemp)

	for _, emp := range employees {
		// string to int
		e_id, err := strconv.Atoi(params["id"])
		if err == nil {
			if emp.EmpId == e_id {
				DataBase.Delete(&emp)
	// oldemp.Email = updatedEmp.Email
	// oldemp.EmpId = updatedEmp.EmpId
	// oldemp.Gender = updatedEmp.Gender
	// oldemp.EmpName = updatedEmp.EmpName
	// oldemp.EmpSalary = updatedEmp.EmpSalary

	// DataBase.Update()
	// DataBase.Delete(&oldemp)
				DataBase.Create(&updatedEmp)
				json.NewEncoder(w).Encode("Employee Updated !")

			}
			// DataBase.Save(&employees)
		}

	}
	// json.NewEncoder(w).Encode("Employee Id Not Found")
	defer r.Body.Close()
}
