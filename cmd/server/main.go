package main

import (
	"GoFetchPointsAPI/internal/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.SetupRoutes(router)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
