package routes

import (
	"playPal-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/items", handlers.GetItems)
	router.POST("/items", handlers.CreateItem)
}
