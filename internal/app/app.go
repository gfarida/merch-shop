package app

import (
	"log/slog"
	"merch-shop/internal/api"
	"merch-shop/internal/middleware"

	"github.com/gin-gonic/gin"
)

type App struct {
	HTTPServer *gin.Engine
}

func New(log *slog.Logger) (*App, error) {
	router := gin.Default()

	router.POST("/api/auth/login", api.Login)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/api/info", api.GetUserInfo)
	protected.POST("/api/sendCoin", api.SendCoins)
	protected.GET("/api/buy/:item", api.BuyItem)
	protected.GET("/api/transactions", api.GetUserTransactions)

	return &App{HTTPServer: router}, nil
}
