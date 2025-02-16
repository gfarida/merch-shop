package api

import (
	"net/http"

	"merch-shop/internal/services"
	"merch-shop/pkg/logger"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		logger.Log().Error("Missing user_id in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	info, err := services.GetUserFullInfo(userID.(string))
	if err != nil {
		logger.Log().Error("Failed to retrieve user info", "user_id", userID, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user information"})
		return
	}

	c.JSON(http.StatusOK, info)
}
