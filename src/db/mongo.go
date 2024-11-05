package db

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
)

var Client *mongo.Client

func Connect() {
    clientOptions := options.Client().ApplyURI("mongodb+srv://ilhamyusufwiam:uC9Yn8WHnVuJKAin@go-http-server.2sq7e.mongodb.net/?retryWrites=true&w=majority&appName=go-http-server")
    var err error
    Client, err = mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    err = Client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Connected to MongoDB!")
}