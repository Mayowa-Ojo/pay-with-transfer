package paystack

import (
	"context"
	"fmt"
	"net/http"
	"pay-with-transfer/config"
	"pay-with-transfer/shared"
	"time"
)

type Service interface {
	CreateAndAssignVirtualAccount(context.Context, AssignVirtualAccountRequest) (*AssignVirtualAccountResponse, error)
}

type Client struct {
	baseURL    string
	secretKey  string
	httpClient *http.Client
}

func New(cfg config.PaystackConfig, httpClient *http.Client) Service {
	return &Client{
		baseURL:    cfg.BaseURL,
		secretKey:  cfg.SecretKey,
		httpClient: httpClient,
	}
}

func (c *Client) CreateAndAssignVirtualAccount(ctx context.Context, req AssignVirtualAccountRequest) (*AssignVirtualAccountResponse, error) {
	resp := &AssignVirtualAccountResponse{}

	resp.Status = true
	resp.Message = "Assigned dedicated account"
	resp.Data.Customer.FirstName = shared.DEFAULT_ACCOUNT_HOLDER_FIRST_NAME
	resp.Data.Customer.LastName = shared.DEFAULT_ACCOUNT_HOLDER_LAST_NAME
	resp.Data.Customer.Email = shared.GeneratePayEmail()
	resp.Data.Customer.Phone = shared.GeneratePayPhoneNumber()
	resp.Data.DedicatedAccount.Bank.Name = BANK_NAME_WEMA
	resp.Data.DedicatedAccount.Bank.ID = BANK_ID_WEMA
	resp.Data.DedicatedAccount.Bank.Slug = BANK_SLUG_WEMA
	resp.Data.DedicatedAccount.AccountName = fmt.Sprintf("%s %s", shared.DEFAULT_ACCOUNT_HOLDER_FIRST_NAME, shared.DEFAULT_ACCOUNT_HOLDER_LAST_NAME)
	resp.Data.DedicatedAccount.AccountNumber = shared.GenerateAccountNumber()
	resp.Data.DedicatedAccount.Assigned = true
	resp.Data.DedicatedAccount.Currency = CURRENCY_NGN
	resp.Data.DedicatedAccount.Active = true
	resp.Data.DedicatedAccount.CreatedAt = time.Now()
	resp.Data.DedicatedAccount.UpdatedAt = time.Now()
	resp.Data.Identification.Status = "success"
	resp.WithCustomerID()
	resp.WithCustomerCode()
	resp.WithAccountID()

	time.Sleep(time.Millisecond * 700)

	return resp, nil
}
