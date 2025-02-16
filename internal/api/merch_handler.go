package api

import (
	"net/http"

	"merch-shop/internal/services"
	"merch-shop/pkg/logger"

	"github.com/gin-gonic/gin"
)

func BuyItem(c *gin.Context) {
	item := c.Param("item")
	if item == "" {
		logger.Log().Error("Item name is missing in request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Item name is required"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		logger.Log().Error("Missing user_id in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := services.BuyItem(userID.(string), item); err != nil {
		logger.Log().Error("Failed to process purchase", "user_id", userID, "item", item, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Purchase failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Purchase successful"})
}
