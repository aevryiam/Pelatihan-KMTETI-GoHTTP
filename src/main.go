package main

import (
	"Pelatihan-KMTETI-GoHTTP/api"
	"log"
	"net/http"
)

func main() {
	router := api.SetupRouter()
	log.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
