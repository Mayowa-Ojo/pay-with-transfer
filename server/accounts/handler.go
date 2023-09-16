package accounts

import (
	"context"
	"pay-with-transfer/services/paystack"
	paylog "pay-with-transfer/shared/logger"
	"pay-with-transfer/store"
	"time"
)

type Handler struct {
	store    store.Store
	paystack paystack.Service
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

func (h *Handler) GeneratePoolAccounts(ctx context.Context, count int) error {
	logger := paylog.WithTrace(ctx).With(paylog.LOG_FIELD_FUNCTION_NAME, "FetchSingleAccount")

	logger.Infof("generating %d pool accounts", count)

	for i := 0; i < count; i++ {
		if err := h.generateSinglePoolAccount(ctx); err != nil {
			logger.With(paylog.LOG_FIELD_ERROR, err).Error("error generating single pool account")
			continue
		}
		time.Sleep(time.Millisecond * 100)
	}

	return nil
}

func (h *Handler) generateSinglePoolAccount(ctx context.Context) error {
	logger := paylog.WithTrace(ctx).With(paylog.LOG_FIELD_FUNCTION_NAME, "generateSinglePoolAccount")

	ah := store.AccountHolder{}
	ah.WithDefaults()

	// resp, err := h.paystack.CreateAndAssignVirtualAccount(ctx, paystack.AssignVirtualAccountRequest{})
	// if err != nil {
	// 	logger.With(paylog.LOG_FIELD_ERROR, err).Error("failed to create virtual account")
	// 	return err
	// }

	// logger.Infof("response: %+v", resp)

	if err := h.store.CreateAccountHolder(ctx, ah); err != nil {
		logger.With(paylog.LOG_FIELD_ERROR, err).Error("failed to create account holder")
		return err
	}

	logger.Info("generated single pool account: %s", ah.ID)

	return nil
}
