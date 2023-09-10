package accounts

import (
	"context"
	paylog "pay-with-transfer/shared/logger"
	"pay-with-transfer/store"
)

type Handler struct {
	store store.Store
}

func New(store store.Store) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) FetchSingleAccount(ctx context.Context, accountID string) (*store.Account, error) {
	logger := paylog.WithTrace(ctx).With(paylog.LOG_FIELD_FUNCTION_NAME, "FetchSingleAccount")

	logger.Infof("fetching account with id: %s", accountID)

	account, err := h.store.GetAccountByID(ctx, accountID)
	if err != nil {
		logger.With(paylog.LOG_FIELD_ERROR, err).Error("failed to find account: %s", accountID)
		return nil, err
	}

	return account, nil
}
