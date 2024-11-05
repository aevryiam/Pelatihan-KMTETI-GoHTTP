// management.go
package main

import "github.com/gofiber/fiber/v2"

func createApp() *fiber.App {
    app := fiber.New()

    // API routes for books
    app.Get("/books", getAllBooks)
    app.Post("/books", addBook)
    app.Put("/books/:id", updateBook)
    app.Delete("/books/:id", deleteBook)

    // API routes for employees
    app.Get("/employees", getAllEmployees)
    app.Post("/employees", addEmployee)

    return app
}

func Handler(w http.ResponseWriter, r *http.Request) {
    app := createApp()
    app.Handler()(w, r)
}
