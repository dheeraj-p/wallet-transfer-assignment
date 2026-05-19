package dtos

type TransferRequest struct {
	IdempotencyKey string `json:"idempotencyKey"`
	FromWalletId   string `json:"fromWalletId"`
	ToWalletId     string `json:"toWalletId"`
	Amount         int64  `json:"amount"`
}
