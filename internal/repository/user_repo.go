package repository

import (
	"log/slog"
	"merch-shop/internal/db"
	"merch-shop/internal/models"
	"merch-shop/pkg/logger"
)

func GetUserByID(userID string) (*models.User, error) {
	var user models.User
	err := db.DB.Get(&user, "SELECT * FROM users WHERE id = $1 AND deleted_at IS NULL", userID)
	if err != nil {
		logger.Log().Error("Failed to get user by ID", slog.String("user_id", userID), slog.String("error", err.Error()))
		return nil, err
	}
	return &user, nil
}
