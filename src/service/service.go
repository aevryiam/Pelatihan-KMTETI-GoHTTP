package service

import (
    "encoding/json"
    "net/http"
    "Pelatihan-KMTETI-GoHTTP/src/db"
    "Pelatihan-KMTETI-GoHTTP/src/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    //"go.mongodb.org/mongo-driver/mongo/options"
    "context"
)

var bookCollection *mongo.Collection
var employeeCollection *mongo.Collection

func init() {
    db.ConnectDB()
    bookCollection = db.Client.Database("go-http-server").Collection("books")
    employeeCollection = db.Client.Database("go-http-server").Collection("employees")
}

// GetAllBooks retrieves all books
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
    var books []models.Book
    cursor, err := bookCollection.Find(context.TODO(), bson.D{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(context.TODO())
    
    for cursor.Next(context.TODO()) {
        var book models.Book
        cursor.Decode(&book)
        books = append(books, book)
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

// GetBookDetail retrieves a specific book by ID
func GetBookDetail(w http.ResponseWriter, r *http.Request) {
    // Implement the logic to get book details by ID
}

// AddBook adds a new book to the database
func AddBook(w http.ResponseWriter, r *http.Request) {
    // Implement the logic to add a new book
}

// UpdateBook updates stock and price of a book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
    // Implement the logic to update stock and price of a book
}

// DeleteBook deletes a specific book by ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
    // Implement the logic to delete a book by ID
}

// GetAllEmployees retrieves all employees
func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
    var employees []models.Employee
    cursor, err := employeeCollection.Find(context.TODO(), bson.D{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(context.TODO())
    
    for cursor.Next(context.TODO()) {
        var employee models.Employee
        cursor.Decode(&employee)
        employees = append(employees, employee)
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(employees)
}

// AddEmployee adds a new employee to the database
func AddEmployee(w http.ResponseWriter, r *http.Request) {
    // Implement the logic to add a new employee
}