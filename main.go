package main

import (
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "restful-api/docs"
)

// @title Roast API
// @version 1.0
// @description A simple API that serves random roasts
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email your@email.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host roast-api.fly.dev
// @BasePath /api/v1
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

	// Swagger endpoint
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	http.HandleFunc("/api/v1/roasts", func(w http.ResponseWriter, r *http.Request) {
		handleRoast(w, r, database)
	})

	log.Printf("server starting on port %s", config.Port)
	log.Printf("Swagger docs available at: http://localhost:%s/swagger/index.html", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
