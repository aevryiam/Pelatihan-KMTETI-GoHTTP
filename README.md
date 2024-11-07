# Project Final Pelatihan KMTETI Web Development - Backend with Go
1. Buat file .env
``` bash
MONGODB="mongodb_connection_string"
```
2. Install package yang dibutuhkan
``` bash
go mod tidy
go mod init Pelatihan-KMTETI-GoHTTp
go install github.com/gorilla/mux
go get go.mongodb.org/mongo-driver/mongo
```
3. Jalankan program
``` bash
go run src/main.go
```
4. Cek API endpoint menggunakan curl
``` bash
curl -X GET http://localhost:8080/books
curl -X POST http://localhost:8080/books/add -d '{"title": "New Book", "author": "Author Name"}' -H "Content-Type: application/json"
curl -X GET http://localhost:8080/employees
curl -X POST http://localhost:8080/employees/add -d '{"name": "John Doe", "position": "Developer"}' -H "Content-Type: application/json"
```
5. Bisa juga menggunakan software Postman untuk mengecek API endpoint
```
http://localhost:8080/books
http://localhost:8080/books/{id}
http://localhost:8080/employees
http://localhost:8080/employees/{id}
```
