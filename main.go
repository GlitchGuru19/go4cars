package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := os.Getenv("ENV")
	var Port string
	if env == "production" {
		Port = "80"
	} else {
		Port = "8080"
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, People!")
	})

	router.Run(":" + Port)

}
