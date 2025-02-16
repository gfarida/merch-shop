package tests

import (
	"merch-shop/internal/models"
	"merch-shop/internal/repository"
	"merch-shop/internal/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockUserRepo struct{}

func (m *mockUserRepo) GetUserByID(userID string) (*models.User, error) {
	if userID == "valid_user" {
		return &models.User{ID: "valid_user", Balance: 1000}, nil
	}
	return nil, repository.ErrUserNotFound
}

func TestGetUserInfo(t *testing.T) {
	repo := &mockUserRepo{}
	services.SetUserRepository(repo)

	user, err := services.GetUserInfo("valid_user")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, int64(1000), user.Balance)

	_, err = services.GetUserInfo("invalid_user")
	assert.Error(t, err)
}
