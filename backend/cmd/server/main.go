package main

import (
	"log"
	"net/http"
	"os"

	"github.com/RaphaelEvanW/job-application-tracker/backend/internal/database"
	"github.com/RaphaelEvanW/job-application-tracker/backend/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	log.Println("JWT_SECRET =", os.Getenv("JWT_SECRET"))

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

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
