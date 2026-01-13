package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/RaphaelEvanW/job-application-tracker/backend/internal/database"
	"github.com/RaphaelEvanW/job-application-tracker/backend/internal/models"
	"github.com/RaphaelEvanW/job-application-tracker/backend/internal/utils"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errror": "failed to process password",
		})
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email already exists",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}
