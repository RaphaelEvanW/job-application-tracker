package handlers

import (
	"net/http"

	"github.com/RaphaelEvanW/job-application-tracker/backend/internal/database"
	"github.com/RaphaelEvanW/job-application-tracker/backend/internal/models"
	"github.com/gin-gonic/gin"
)

func GetMe(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}
