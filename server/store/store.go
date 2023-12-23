package store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=store.go -destination store_mock.go -package store . Store

type Store interface {
	AccountStore
}

type AccountStore interface {
	GetAccountByID(ctx context.Context, id string) (*Account, error)
	GetEphemeralAccountByID(ctx context.Context, id string) (*EphemeralAccount, error)
	UpdateAccount(ctx context.Context, ac Account) error
	CreateAccountHolder(ctx context.Context, ah AccountHolder) error
	FindDormantAccount(ctx context.Context) (*Account, error)
	FindActiveEphemeralAccounts(ctx context.Context, limit int) ([]*EphemeralAccount, error)
	FindEphemeralAccountByAccountID(ctx context.Context, accountID string) (*EphemeralAccount, error)
	CreateEphemeralAccount(ctx context.Context, ea EphemeralAccount) error
	UpdateEphemeralAccount(ctx context.Context, ea EphemeralAccount) error
	GetTransactionByID(ctx context.Context, id string) (*Transaction, error)
	CreateTransaction(ctx context.Context, t Transaction) error
	UpdateTransaction(ctx context.Context, t Transaction) error
	GetTransactionByEphemeralAccountID(ctx context.Context, eaID string) (*Transaction, error)
}

type DataStore struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Store {
	return &DataStore{db}
}
