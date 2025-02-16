package api

import (
	"net/http"

	"merch-shop/internal/services"
	"merch-shop/pkg/logger"

	"github.com/gin-gonic/gin"
)

// SendCoins handles coin transfer requests.
func SendCoins(c *gin.Context) {
	var req struct {
		ToUser string `json:"toUser"`
		Amount int64  `json:"amount"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Log().Error("Invalid JSON request", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	senderID, exists := c.Get("user_id")
	if !exists {
		logger.Log().Error("Missing user_id in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := services.SendCoins(senderID.(string), req.ToUser, req.Amount); err != nil {
		logger.Log().Error("Failed to send coins", "sender_id", senderID, "receiver_id", req.ToUser, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction successful"})
}
