package accounts

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"pay-with-transfer/cache"
	"pay-with-transfer/config"
	"pay-with-transfer/services/paystack"
	"pay-with-transfer/shared"
	"pay-with-transfer/shared/backoff"
	paylog "pay-with-transfer/shared/logger"
	"pay-with-transfer/store"

	"github.com/charmbracelet/log"

	"github.com/volatiletech/null/v8"
)

type Handler struct {
	store    store.Store
	cache    cache.Cache
	cfg      config.Config
	paystack paystack.Service
}

func New(store store.Store, cache cache.Cache, cfg config.Config) *Handler {
	return &Handler{
		store: store,
		cache: cache,
		cfg:   cfg,
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

func (h *Handler) FetchSingleEphemeralAccount(ctx context.Context, accountID string) (*store.EphemeralAccount, error) {
	logger := paylog.WithTrace(ctx).With(paylog.LOG_FIELD_FUNCTION_NAME, "FetchSingleEphemeralAccount")

	logger.Infof("fetching ephemeral account with id: %s", accountID)

	ea, err := h.store.GetEphemeralAccountByID(ctx, accountID)
	if err != nil {
		logger.With(paylog.LOG_FIELD_ERROR, err).Error("failed to find ephemeral account: %s", accountID)
		return nil, err
	}

	account, err := h.store.GetAccountByID(ctx, ea.AccountID.String())
	if err != nil {
		logger.With(paylog.LOG_FIELD_ERROR, err).Error("failed to find account: %s", accountID)
		return nil, err
	}

	ea.AccountName = account.AccountName
	ea.BankName = account.BankName.String
	ea.AccountNumber = account.AccountNumber
	ea.PaymentAmount = float64(ea.Amount / 100)

	return ea, nil
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

	resp, err := h.paystack.CreateAndAssignVirtualAccount(ctx, paystack.AssignVirtualAccountRequest{})
	if err != nil {
		logger.With(paylog.LOG_FIELD_ERROR, err).Error("failed to create virtual account")
		return err
	}

	b, err := json.Marshal(resp)
	if err != nil {
		logger.With(paylog.LOG_FIELD_ERROR, err).Error("JSONError: failed to parse response")
		return err
	}

	ah.ProviderCode = resp.Data.Customer.CustomerCode
	ah.ProviderID = strconv.Itoa(resp.Data.Customer.ID)
	ah.ProviderResponse = null.StringFrom(string(b))

	if err := h.store.CreateAccountHolder(ctx, ah); err != nil {
		logger.With(paylog.LOG_FIELD_ERROR, err).Error("failed to create account holder")
		return err
	}

	logger.Infof("generated single pool account: %s", ah.ID)

	return nil
}

func (h *Handler) CreateEphemeralAccount(ctx context.Context, dto CreateEphemeralAccountDTO) (*store.EphemeralAccount, error) {
	logger := paylog.WithTrace(ctx).With(paylog.LOG_FIELD_FUNCTION_NAME, "CreateEphemeralAccount")

	var account *store.EphemeralAccount
	var err error
	_ = backoff.WithRetry(ctx, func() error {
		ea, isRetryable, cErr := h.createEphemeralAccountWithBackoff(ctx, dto, logger)
		account = ea
		err = cErr
		if cErr != nil && isRetryable {
			return cErr
		}
		return nil
	})
	if err != nil {
		logger.With(paylog.LOG_FIELD_ERROR, err).Error("failed to create ephemeral account")
		return nil, err
	}

	txn := store.Transaction{}
	txn.WithDefaults()
	txn.EphemeralAccountID = account.ID
	txn.AccountID = account.AccountID
	txn.Amount = account.Amount
	txn.Currency = paystack.CURRENCY_NGN
	txn.AccountName = null.StringFrom(account.AccountName)
	txn.BankName = null.StringFrom(account.BankName)
	txn.AccountNumber = null.StringFrom(account.AccountNumber)
	txn.Provider = null.StringFrom(account.Provider)

	if err = h.store.CreateTransaction(ctx, txn); err != nil {
		logger.With(paylog.LOG_FIELD_ERROR, err).Error("failed to create transaction")
		return nil, err
	}

	account.TransactionID = txn.ID.String()

	return account, nil
}

func (h *Handler) createEphemeralAccountWithBackoff(ctx context.Context, dto CreateEphemeralAccountDTO, logger *log.Logger) (*store.EphemeralAccount, bool, error) {
	logger.Infof("creating ephemeral account with amount: %.2f", dto.Amount)

	//find a dormant account
	account, err := h.store.FindDormantAccount(ctx)
	if err != nil {
		logger.With(paylog.LOG_FIELD_ERROR, err).Error("error finding dormant account")
		return nil, false, err
	}

	//ensure the last ephemeral account is expired
	ephemeralAccount, err := h.store.FindEphemeralAccountByAccountID(ctx, account.ID.String())
	if err != nil && !shared.IsErrNoRows(err) {
		logger.With(paylog.LOG_FIELD_ERROR, err).Error("error finding ephemeral account")
		return nil, false, err
	}

	if ephemeralAccount != nil && ephemeralAccount.Status != store.EphemeralAccountExpired {
		err := fmt.Errorf("last ephemeral account is still active")
		logger.Error(err.Error())
		return nil, false, err
	}

	//acquire lock on account
	key := account.ID.String()
	val := fmt.Sprintf("%s:%d", key, time.Now().Unix())
	r := h.cache.SetNX(ctx, key, val, time.Hour*6)
	if !r.Val() {
		logger.Infof("setnx: %t for %s | account is locked", r.Val(), key)
		return nil, true, fmt.Errorf("failed to create ephemeral account")
	}

	//create ephemeral account
	ea := store.EphemeralAccount{}
	ea.WithDefaults()
	ea.AccountID = account.ID
	ea.Amount = shared.ToBaseUnitAmount(dto.Amount)
	ea.ExpiresAt = time.Now().Add(h.cfg.App.EphemeralAccountTTL)
	if err := h.store.CreateEphemeralAccount(ctx, ea); err != nil {
		logger.With(paylog.LOG_FIELD_ERROR, err).Error("failed to create account holder")
		return nil, false, err
	}

	ea.AccountName = account.AccountName
	ea.BankName = account.BankName.String
	ea.AccountNumber = account.AccountNumber
	ea.PaymentAmount = dto.Amount
	ea.Provider = account.Provider

	//update account dormant status
	account.IsDormant = false
	if err := h.store.UpdateAccount(ctx, *account); err != nil {
		logger.With(paylog.LOG_FIELD_ERROR, err).Error("failed to update account")
		return nil, false, err
	}

	logger.Infof("completed ephemeral account creation with amount: %.2f", dto.Amount)

	return &ea, false, nil
}
