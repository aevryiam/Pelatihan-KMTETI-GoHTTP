// main.go
package main

import (
    "log"
    "net/http"
    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    connectDB() // Connect to the database

    // Start the server
    log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(Handler)))
}
