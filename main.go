package main

import (
	"log"
	"net/http"
	"os"
	"sync/atomic"   // for thread-safe visitor counter
	"time"

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

	// Visitor counter (atomic for safe concurrent access)
	var visitors uint64 = 0

	// Define a simple GET route for the root path
	router.GET("/", func(c *gin.Context) {
		// Increment visitor counter safely
		currentVisitors := atomic.AddUint64(&visitors, 1)

		// Get the current time in human-readable format
		currentTime := time.Now().Format("Mon Jan 2 15:04:05 MST 2006")

		// Respond with a formatted string showing status, visitor count, and current time
		c.String(http.StatusOK, `
🚗  Go4Cars Server  🧑‍🤝‍🧑
=========================
Status: 🚀 Running
Visitors: 👤 %d
Current Time: %s
Enjoy your ride! 🌟
`, currentVisitors, currentTime)
	})

	// Start the server on the specified port
	log.Printf("Starting server on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
