// management.service.go
package main

import (
    "context"
    "net/http"
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func getAllBooks(c *fiber.Ctx) error {
    collection := client.Database("library").Collection("books")
    cursor, err := collection.Find(context.TODO(), bson.M{})
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    defer cursor.Close(context.TODO())

    var books []Book
    for cursor.Next(context.TODO()) {
        var book Book
        if err := cursor.Decode(&book); err != nil {
            return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }
        books = append(books, book)
    }

    return c.JSON(books)
}

func addBook(c *fiber.Ctx) error {
    var book Book
    if err := c.BodyParser(&book); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    book.ID = primitive.NewObjectID()
    collection := client.Database("library").Collection("books")
    _, err := collection.InsertOne(context.TODO(), book)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(http.StatusCreated).JSON(book)
}

func updateBook(c *fiber.Ctx) error {
    id := c.Params("id")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
    }

    var updateData struct {
        Stock int     `json:"stock"`
        Price float64 `json:"price"`
    }

    if err := c.BodyParser(&updateData); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    collection := client.Database("library").Collection("books")
    _, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, bson.M{"$set": bson.M{"stock": updateData.Stock, "price": updateData.Price}})
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Book updated successfully"})
}

func deleteBook(c *fiber.Ctx) error {
    id := c.Params("id")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
    }

    collection := client.Database("library").Collection("books")
    _, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Book deleted successfully"})
}

func getAllEmployees(c *fiber.Ctx) error {
    collection := client.Database("library").Collection("employees")
    cursor, err := collection.Find(context.TODO(), bson.M{})
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    defer cursor.Close(context.TODO())

    var employees []Employee
    for cursor.Next(context.TODO()) {
        var employee Employee
        if err := cursor.Decode(&employee); err != nil {
            return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }
        employees = append(employees, employee)
    }

    return c.JSON(employees)
}

func addEmployee(c *fiber.Ctx) error {
    var employee Employee
    if err := c.BodyParser(&employee); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    employee.ID = primitive.NewObjectID()
    collection := client.Database("library").Collection("employees")
    _, err := collection.InsertOne(context.TODO(), employee)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(http.StatusCreated).JSON(employee)
}
