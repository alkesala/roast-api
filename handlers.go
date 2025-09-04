package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)


func handleRoast (w http.ResponseWriter, r *http.Request, db *pgx.Conn) {
	w.Header().Set("Content-Type", "application/json")

	var roast string
	query := "SELECT content FROM roasts ORDER BY RANDOM() LIMIT 1"
	err := db.QueryRow(context.Background(), query).Scan(&roast)
	if err != nil {
		log.Printf("db query failed: %v", err)
		http.Error(w,`{"error": Failed to fetch roast}`, http.StatusInternalServerError)
		return
	}
	 response := fmt.Sprintf(`{"roast": "%s"}`, roast)
    w.Write([]byte(response))
}
