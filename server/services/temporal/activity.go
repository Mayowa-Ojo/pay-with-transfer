package temporal

import (
	"context"
	"pay-with-transfer/store"

	"github.com/jmoiron/sqlx"
	"go.temporal.io/sdk/activity"
)

type AccountActivity struct {
	store store.Store
}

func (a *AccountActivity) WithDatastore(db *sqlx.DB) error {
	a.store = store.New(db)
	return nil
}

func (a *AccountActivity) GetActiveAccounts(ctx context.Context) (*GetActiveAccountsResponse, error) {
	logger := activity.GetLogger(ctx)

	accounts, err := a.store.FindActiveEphemeralAccounts(ctx, 10)
	if err != nil {
		logger.Error("error finding ephemeral account", "error", err.Error())
		return nil, err
	}

	return &GetActiveAccountsResponse{
		Accounts: accounts,
	}, nil
}

func (a *AccountActivity) UpdateEphemeralAccount(ctx context.Context, param UpdateEphemeralAccountParam) error {
	logger := activity.GetLogger(ctx)

	err := a.store.UpdateEphemeralAccount(ctx, param.Account)
	if err != nil {
		logger.Error("error while updating ephemeral account", "error", err.Error())
		return err
	}
	return nil
}
