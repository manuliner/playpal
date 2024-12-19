package handlers

import (
	"net/http"
	"playPal-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {
	// Logic to retrieve items from the database
	items := []models.Item{} // Replace with actual database retrieval logic
	c.JSON(http.StatusOK, items)
}

func CreateItem(c *gin.Context) {
	var newItem models.Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to save newItem to the database
	c.JSON(http.StatusCreated, newItem)
}
