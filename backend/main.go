package main

import (
	"context"
	"log"
	"net/http"
	"playPal-backend/internal/config"
	"playPal-backend/internal/db"
	"playPal-backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Connect to MongoDB
	mongoClient, err := db.Connect(cfg.DbUri)
	if err != nil {
		log.Fatalf("could not connect to MongoDB: %v", err)
	}
	defer mongoClient.Disconnect(context.Background())

	// Set up routes
	router := gin.Default()
	routes.SetupRoutes(router)

	// Start server
	log.Printf("Starting server on port %s...", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
