// management.model.go
package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
    ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Title  string             `json:"title" bson:"title"`
    Author string             `json:"author" bson:"author"`
    Year   int                `json:"year" bson:"year"`
    Stock  int                `json:"stock" bson:"stock"`
    Price  float64            `json:"price" bson:"price"`
}

type Employee struct {
    ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Name             string             `json:"name" bson:"name"`
    NIK              string             `json:"nik" bson:"nik"`
    LastEducation    string             `json:"lastEducation" bson:"lastEducation"`
    EntryDate        string             `json:"entryDate" bson:"entryDate"`
    EmploymentStatus string             `json:"employmentStatus" bson:"employmentStatus"`
}
