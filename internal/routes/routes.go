package routes

import (
	"GoFetchPointsAPI/internal/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {

	router.HandleFunc("/receipt/process", handlers.ProcessReceipt).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", handlers.GetReceiptPoints).Methods("GET")
	router.HandleFunc("/receipts/points", handlers.GetAllReceiptPoints).Methods("GET")
}
