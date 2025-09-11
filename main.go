package main

import (
	"log"
	"net/http"
)

func main() {
	config := Load()
	if config.Port == "" {
		log.Fatal("Could not load PORT")
	}
	if config.Database == "" {
		log.Fatal("Could not load Database config")
	}
	database, err := Connect()
	if err != nil {
		log.Fatal("failed to connect db", err)
	}
	defer database.Close()

	http.HandleFunc("/api/roasts", func(w http.ResponseWriter, r *http.Request) {
		handleRoast(w, r, database)

	})

	log.Printf("server starting on port %s", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
