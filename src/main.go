package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "Pelatihan-KMTETI-GoHTTP/api"
)

func main() {
    r := mux.NewRouter()
    api.RegisterRoutes(r)

    log.Println("Server is running on port 8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}