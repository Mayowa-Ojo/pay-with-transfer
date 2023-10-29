package transactions

import (
	"context"
	"pay-with-transfer/store"
)

type TransactionService interface {
	FetchSingleTransaction(ctx context.Context, transactionID string) (*store.Transaction, error)
}
