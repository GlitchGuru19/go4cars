package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Get PORT from environment variables
	// Fly.io automatically sets the PORT environment variable
	port := os.Getenv("PORT")
	if port == "" {
		// Fallback if PORT is not set
		port = "8080"
	}

	// Initialize Gin router
	router := gin.Default()

	// Define a simple GET route for the root path
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, `
	🚗  Go4Cars Server  🧑‍🤝‍🧑
	=========================
	Status: 🚀 Running
	Visitors: 👤 42
	Enjoy your ride! 🌟
	`)
	})
	

	// Start the server on the specified port
	log.Printf("Starting server on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
