package transactions

import (
	"context"
	"pay-with-transfer/cache"
	"pay-with-transfer/config"
	paylog "pay-with-transfer/shared/logger"
	"pay-with-transfer/store"
)

type Handler struct {
	store store.Store
	cache cache.Cache
	cfg   config.Config
}

func New(store store.Store, cache cache.Cache, cfg config.Config) *Handler {
	return &Handler{
		store: store,
		cache: cache,
		cfg:   cfg,
	}
}

func (h *Handler) FetchSingleTransaction(ctx context.Context, transactionID string) (*store.Transaction, error) {
	logger := paylog.WithTrace(ctx).With(paylog.LOG_FIELD_FUNCTION_NAME, "FetchSingleTransaction")

	logger.Infof("fetching transaction with id: %s", transactionID)

	transaction, err := h.store.GetTransactionByID(ctx, transactionID)
	if err != nil {
		logger.With(paylog.LOG_FIELD_ERROR, err).Error("failed to find transaction: %s", transactionID)
		return nil, err
	}

	return transaction, nil
}
