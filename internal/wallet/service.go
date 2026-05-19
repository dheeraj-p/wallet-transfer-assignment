package wallet

import (
	"context"

	"wallet-transfer-assignment/internal/api/dtos"
)

type WalletService interface {
	TransferFunds(ctx context.Context, req dtos.TransferRequest) error
}

type walletService struct {
	repo WalletRepository
}

func NewService(repo WalletRepository) WalletService {
	return &walletService{repo: repo}
}

func (s *walletService) TransferFunds(ctx context.Context, req dtos.TransferRequest) error {
	// Placeholder for actual transfer logic
	return nil
}
