package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RoastResponse struct {
	Roast string `json:"roast"`
}

func handleRoast(w http.ResponseWriter, r *http.Request, db *pgxpool.Pool) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	var roast string
	query := "SELECT content FROM roasts ORDER BY RANDOM() LIMIT 1"
	err := db.QueryRow(context.Background(), query).Scan(&roast)
	if err != nil {
		log.Printf("db query failed: %v", err)
		http.Error(w, `{"error": "Failed to fetch roast"}`, http.StatusInternalServerError)
		return
	}
	response := RoastResponse{Roast: roast}
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, `{"error": "Failed to encode json"}`, http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(jsonData); err != nil {
		log.Printf("Failed to write the Roast %v", err)
	}
}
