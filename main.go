package main

import (
	"context"
	"log"
	"net/http"
	"os"
"github.com/joho/godotenv"
)

func main () {
	if err := godotenv.Load(); err != nil {
	 log.Println("Env file not found")
	}
	database, err := Connect()
	if err != nil {
		log.Fatal("failed to connect db", err)
	}
		defer database.Close(context.Background())

	// port config, fallbacks into 8080 when local development.
	// deployment server handles the PORT
	port := os.Getenv("PORT")
	if port  == "" {
	port = "8081"
	}

	http.HandleFunc("/api/roasts", func(w http.ResponseWriter, r *http.Request){
	handleRoast(w, r, database)

	})

log.Printf("server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}


