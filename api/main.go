package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

var client *mongo.Client

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

func connectDB() {
    var err error
    mongoURI := os.Getenv("MONGODB_URI")
    clientOptions := options.Client().ApplyURI(mongoURI)
    client, err = mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    // Check the connection
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to MongoDB!")
}

func createApp() *fiber.App {
    app := fiber.New()

    // API routes for books
    app.Get("/books", getAllBooks)
    app.Get("/books/:id", getBookDetails)
    app.Post("/books", addBook)
    app.Put("/books/:id", updateBook)
    app.Delete("/books/:id", deleteBook)

    // API routes for employees
    app.Get("/employees", getAllEmployees)
    app.Post("/employees", addEmployee)

    return app
}

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

func getBookDetails(c *fiber.Ctx) error {
    id := c.Params("id")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
    }

    collection := client.Database("library").Collection("books")
    var book Book

    err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&book)
    if err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
    }

    return c.JSON(book)
}

func addBook(c *fiber.Ctx) error {
    var book Book
    if err := c.BodyParser(&book); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    book.ID = primitive.NewObjectID() // Generate a new ObjectID
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

    employee.ID = primitive.NewObjectID() // Generate a new ObjectID
    collection := client.Database("library").Collection("employees")
    _, err := collection.InsertOne(context.TODO(), employee)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(http.StatusCreated).JSON(employee)
}

func Handler(w http.ResponseWriter, r *http.Request) {
    app := createApp()
    app.Handler()(w, r)
}

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