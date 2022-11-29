package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	fmt.Println("Welcome to the DataBase Connectivity")
	r := mux.NewRouter()

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/employee", CreateEmployee).Methods("POST")
	r.HandleFunc("/employees", GetEmployees).Methods("GET")
	r.HandleFunc("/employee/{id}", GetEmployeeById).Methods("GET")
	r.HandleFunc("/employees", DeleteAllEmployees).Methods("DELETE")

	r.HandleFunc("/employee/{id}", DeleteEmployeeByID).Methods("DELETE")
	r.HandleFunc("/employee/{id}", UpdateById).Methods("PUT")

	log.Fatal(http.ListenAndServe(":6000", r))
	// log.Fatal(http.ListenAndServe(":8080", r))

	fmt.Println("CONNECTED !")

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to DataBase Connectivity Session</h1>"))
}
