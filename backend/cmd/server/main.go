package main

import (
	"log"
	"net/http"

	"github.com/RaphaelEvanW/job-application-tracker/backend/internal/database"
	"github.com/RaphaelEvanW/job-application-tracker/backend/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	database.Connect()
	routes.AuthRoutes(r)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.Run(":8080")
}
