package service

import (
	"Pelatihan-KMTETI-GoHTTP/src/db"
	"Pelatihan-KMTETI-GoHTTP/src/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllEmployees() ([]models.Employee, error) {
	collection := db.Client.Database("go-http-server").Collection("employees")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	var employees []models.Employee
	if err = cursor.All(context.TODO(), &employees); err != nil {
		return nil, err
	}
	return employees, nil
}

func AddEmployee(employee models.Employee) error {
	collection := db.Client.Database("go-http-server").Collection("employees")
	_, err := collection.InsertOne(context.TODO(), employee)
	return err
}
