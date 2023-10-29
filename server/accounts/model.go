package accounts

import (
	"context"
	"pay-with-transfer/store"
)

type AccountService interface {
	FetchSingleAccount(ctx context.Context, accountID string) (*store.Account, error)
	FetchSingleEphemeralAccount(ctx context.Context, accountID string) (*store.EphemeralAccount, error)
	GeneratePoolAccounts(ctx context.Context, count int) error
	CreateEphemeralAccount(ctx context.Context, dto CreateEphemeralAccountDTO) (*store.EphemeralAccount, error)
}

type GeneratePoolAccountsDTO struct {
	Count int `json:"count"`
}

type CreateEphemeralAccountDTO struct {
	Amount    float64 `json:"amount" binding:"required"`
	SessionID string  `json:"session_id" binding:"required"`
}
