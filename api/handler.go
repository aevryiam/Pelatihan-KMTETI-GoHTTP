package api

import (
    "encoding/json"
    "net/http"
    "Pelatihan-KMTETI-GoHTTP/src/models"
    "Pelatihan-KMTETI-GoHTTP/src/service"
)

func SetupRouter() *http.ServeMux {
    router := http.NewServeMux()

    router.HandleFunc("/books", getAllBooks)
    router.HandleFunc("/books/detail", getBookDetail)
    router.HandleFunc("/books/add", addBook)
    router.HandleFunc("/books/update", updateBook)
    router.HandleFunc("/books/delete", deleteBook)
    router.HandleFunc("/employees", getAllEmployees)
    router.HandleFunc("/employees/add", addEmployee)

    return router
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
    books, err := service.GetAllBooks()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(books)
}

func getBookDetail(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    book, err := service.GetBookDetail(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(book)
}

func addBook(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := service.AddBook(book); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := service.UpdateBook(book); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if err := service.DeleteBook(id); err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func getAllEmployees(w http.ResponseWriter, r *http.Request) {
    employees, err := service.GetAllEmployees()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(employees)
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
    var employee models.Employee
    if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := service.AddEmployee(employee); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}