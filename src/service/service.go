package service

import (
    "encoding/json"
    "net/http"
    "Pelatihan-KMTETI-GoHTTP/src/db"
    "Pelatihan-KMTETI-GoHTTP/src/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "context"
)

var bookCollection *mongo.Collection
var employeeCollection *mongo.Collection

func init() {
    db.ConnectDB()
    bookCollection = db.Client.Database("go-http-server").Collection("books")
    employeeCollection = db.Client.Database("go-http-server").Collection("employees")
}

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

func GetBookDetail(w http.ResponseWriter, r *http.Request) {
}

func AddBook(w http.ResponseWriter, r *http.Request) {
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
}

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

func AddEmployee(w http.ResponseWriter, r *http.Request) {
}