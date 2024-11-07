package api

import (
    //"net/http"
    "github.com/gorilla/mux"
    "Pelatihan-KMTETI-GoHTTP/src/service"
)

func RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/books", service.GetAllBooks).Methods("GET")
    r.HandleFunc("/books/{id}", service.GetBookDetail).Methods("GET")
    r.HandleFunc("/books", service.AddBook).Methods("POST")
    r.HandleFunc("/books/{id}", service.UpdateBook).Methods("PUT")
    r.HandleFunc("/books/{id}", service.DeleteBook).Methods("DELETE")
    r.HandleFunc("/employees", service.GetAllEmployees).Methods("GET")
    r.HandleFunc("/employees", service.AddEmployee).Methods("POST")
}