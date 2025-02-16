package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"merch-shop/pkg/jwt"
	"merch-shop/pkg/logger" // Импортируем логгер

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// AuthMiddleware проверяет JWT и передает информацию в контекст запроса
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			logger.Log().Error("Authorization header missing or invalid", slog.String("request", c.Request.URL.Path))
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неавторизованный доступ"})
			c.Abort()
			return
		}

		// Извлекаем токен из заголовка
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		parsedToken, err := jwt.Parse(tokenString)
		if err != nil {
			logger.Log().Error("Invalid token", slog.String("error", err.Error()))
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
			c.Abort()
			return
		}

		// Извлекаем данные из токена (например, user_id и is_admin)
		claims, ok := parsedToken.Claims.(*jwt.RegisteredClaims)
		if !ok {
			logger.Log().Error("Invalid token claims", slog.String("request", c.Request.URL.Path))
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неудовлетворительный токен"})
			c.Abort()
			return
		}

		// Добавляем данные в контекст
		ctx := context.WithValue(c.Request.Context(), "user_id", claims.Subject)
		ctx = context.WithValue(ctx, "is_admin", claims.Audience)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
