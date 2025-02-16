package middleware

import (
	"context"
	"net/http"
	"strings"

	"merch-shop/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неавторизованный доступ"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		parsedToken, err := jwt.Parse(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
			c.Abort()
			return
		}

		claims, ok := parsedToken.Claims.(*jwt.RegisteredClaims) // Используем RegisteredClaims
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неудовлетворительный токен"})
			c.Abort()
			return
		}

		ctx := context.WithValue(c.Request.Context(), "user_id", claims.Subject) // Используем Subject для user_id
		ctx = context.WithValue(ctx, "is_admin", claims.Audience)                // Используем Audience для is_admin
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
