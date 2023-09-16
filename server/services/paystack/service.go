package paystack

import (
	"context"
	"net/http"
	"pay-with-transfer/config"
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
	return nil, nil
}
