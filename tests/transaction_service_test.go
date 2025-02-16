package tests

import (
	"merch-shop/internal/repository"
	"merch-shop/internal/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockTransactionRepo struct{}

func (m *mockTransactionRepo) SaveTransaction(senderID, receiverID string, amount int64) error {
	if senderID == receiverID {
		return repository.ErrInvalidTransaction
	}
	if amount <= 0 {
		return repository.ErrInvalidTransaction
	}
	return nil
}

func TestSendCoins_ValidTransaction(t *testing.T) {
	repo := &mockTransactionRepo{}
	services.SetTransactionRepository(repo)

	err := services.SendCoins("valid_sender", "valid_receiver", 50)
	assert.NoError(t, err)
}

func TestSendCoins_SelfTransfer(t *testing.T) {
	repo := &mockTransactionRepo{}
	services.SetTransactionRepository(repo)

	err := services.SendCoins("valid_sender", "valid_sender", 50)
	assert.Error(t, err)
	assert.Equal(t, repository.ErrInvalidTransaction, err)
}

func TestSendCoins_ZeroAmount(t *testing.T) {
	repo := &mockTransactionRepo{}
	services.SetTransactionRepository(repo)

	err := services.SendCoins("valid_sender", "valid_receiver", 0)
	assert.Error(t, err)
	assert.Equal(t, repository.ErrInvalidTransaction, err)
}

func TestSendCoins_NegativeAmount(t *testing.T) {
	repo := &mockTransactionRepo{}
	services.SetTransactionRepository(repo)

	err := services.SendCoins("valid_sender", "valid_receiver", -10)
	assert.Error(t, err)
	assert.Equal(t, repository.ErrInvalidTransaction, err)
}
