package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"merch-shop/internal/api"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUserInfoAPI(t *testing.T) {
	router := gin.Default()
	router.GET("/api/info", api.GetUserInfo)

	req, _ := http.NewRequest("GET", "/api/info", nil)
	req.Header.Set("Authorization", "Bearer valid_token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSendCoinsAPI(t *testing.T) {
	router := gin.Default()
	router.POST("/api/sendCoin", api.SendCoins)

	req, _ := http.NewRequest("POST", "/api/sendCoin", nil)
	req.Header.Set("Authorization", "Bearer valid_token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
