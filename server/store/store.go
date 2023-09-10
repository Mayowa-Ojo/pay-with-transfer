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
}

type DataStore struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Store {
	return &DataStore{db}
}
