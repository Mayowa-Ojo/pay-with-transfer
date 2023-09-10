package accounts

import (
	"context"
	"pay-with-transfer/store"
)

type AccountService interface {
	FetchSingleAccount(ctx context.Context, accountID string) (*store.Account, error)
}
