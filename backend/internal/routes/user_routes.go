package routes

import (
	"github.com/RaphaelEvanW/job-application-tracker/backend/internal/handlers"
	"github.com/RaphaelEvanW/job-application-tracker/backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	user := r.Group("/user")
	user.Use(middleware.JWTAuthMiddleware())
	{
		user.GET("/me", handlers.GetMe)
	}
}
