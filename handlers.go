package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RoastResponse struct {
	Roast string `json:"roast"`
}

func handleRoast(w http.ResponseWriter, r *http.Request, db *pgxpool.Pool) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	log.Printf("Roast request from: %s", r.RemoteAddr)
	if r.Method != http.MethodGet {
		http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var roast string
	query := "SELECT content FROM roasts ORDER BY RANDOM() LIMIT 1"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := db.QueryRow(ctx, query).Scan(&roast)
	if err == pgx.ErrNoRows {
		http.Error(w, `{"error": "No roasts found"}`, http.StatusNotFound)
	}

	if err != nil {
		log.Printf("db query failed: %v", err)
		http.Error(w, `{"error": "Failed to fetch roast"}`, http.StatusInternalServerError)
		return
	}
	response := RoastResponse{Roast: roast}
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, `{"error": "Failed to encode json"}`, http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(jsonData); err != nil {
		log.Printf("Failed to write the Roast %v", err)
		return
	}
}
