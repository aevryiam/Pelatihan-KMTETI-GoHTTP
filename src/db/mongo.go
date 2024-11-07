package db

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
)

var Client *mongo.Client

func ConnectDB() {
    clientOptions := options.Client().ApplyURI("mongodb+srv://ilhamyusufwiam:uC9Yn8WHnVuJKAin@go-http-server.2sq7e.mongodb.net/?retryWrites=true&w=majority&appName=go-http-server")
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    Client = client
}

func DisconnectDB() {
    err := Client.Disconnect(context.TODO())
    if err != nil {
        log.Fatal(err)
    }
}
