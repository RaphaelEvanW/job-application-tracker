package main

import (
	"log"
	"net/http"

	"github.com/RaphaelEvanW/job-application-tracker/backend/internal/database"
	"github.com/RaphaelEvanW/job-application-tracker/backend/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	r.Use(cors.Default())

	database.Connect()
	routes.AuthRoutes(r)
	routes.UserRoutes(r)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.Run(":8080")
}
