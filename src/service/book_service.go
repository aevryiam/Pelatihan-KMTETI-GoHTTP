package service

import (
    "context"
    "Pelatihan-KMTETI-GoHTTP/src/db"
    "Pelatihan-KMTETI-GoHTTP/src/models"
    "go.mongodb.org/mongo-driver/bson"
)

func GetAllBooks() ([]models.Book, error) {
    collection := db.Client.Database("go-http-server").Collection("books")
    cursor, err := collection.Find(context.TODO(), bson.M{})
    if err != nil {
        return nil, err
    }
    var books []models.Book
    if err = cursor.All(context.TODO(), &books); err != nil {
        return nil, err
    }
    return books, nil
}

func GetBookDetail(id string) (models.Book, error) {
    collection := db.Client.Database("go-http-server").Collection("books")
    var book models.Book
    err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&book)
    if err != nil {
        return book, err
    }
    return book, nil
}

func AddBook(book models.Book) error {
    collection := db.Client.Database("go-http-server").Collection("books")
    _, err := collection.InsertOne(context.TODO(), book)
    return err
 
}

func UpdateBook(book models.Book) error {
    collection := db.Client.Database("go-http-server").Collection("books")
    _, err := collection.UpdateOne(context.TODO(), bson.M{"_id": book.ID}, bson.M{"$set": bson.M{"stock": book.Stock, "price": book.Price}})
    return err
}

func DeleteBook(id string) error {
    collection := db.Client.Database("go-http-server").Collection("books")
    _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
    return err
}