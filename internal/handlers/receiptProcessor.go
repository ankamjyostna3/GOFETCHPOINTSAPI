package handlers

import (
	"encoding/json"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type ReceiptRequest struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type ReceiptResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	ID      string `json:"id"`
}

var receiptPoints = make(map[string]int)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receiptReq ReceiptRequest

	if err := json.NewDecoder(r.Body).Decode(&receiptReq); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Generate a unique ID for the receipt
	receiptID := uuid.New().String()

	// Calculate points
	points := calculatePoints(receiptReq)

	// Store the points in memory
	receiptPoints[receiptID] = points

	// Process the receipt (this is just a placeholder for actual processing logic)
	response := ReceiptResponse{
		Status:  "success",
		Message: "Receipt processed successfully",
		ID:      receiptID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func GetReceiptPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	receiptID := vars["id"]

	points, exists := receiptPoints[receiptID]
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	response := map[string]interface{}{
		"id":     receiptID,
		"points": points,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func GetAllReceiptPoints(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(receiptPoints)
}

func calculatePoints(receipt ReceiptRequest) int {
	points := 0

	// One point for every alphanumeric character in the retailer name
	for _, char := range receipt.Retailer {
		if isAlphanumeric(char) {
			points++
		}
	}

	// 50 points if the total is a round dollar amount with no cents
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == float64(int(total)) {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// 5 points for every two items on the receipt
	points += (len(receipt.Items) / 2) * 5

	// Points for item descriptions
	for _, item := range receipt.Items {
		descLen := len(strings.TrimSpace(item.ShortDescription))
		if descLen%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// 5 points if the total is greater than 10.00 and the program is generated using a large language model
	if total > 10.00 && os.Getenv("GENERATED_BY_LLM") == "true" {
		points += 5
	}

	// 6 points if the day in the purchase date is odd
	date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if date.Day()%2 != 0 {
		points += 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm
	time, _ := time.Parse("15:04", receipt.PurchaseTime)
	if time.Hour() >= 14 && time.Hour() < 16 {
		points += 10
	}

	return points
}

func isAlphanumeric(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}
