package store

import "context"

type Store interface {
	Account() AccountStore
}

type AccountStore interface {
	GetAccountByID(ctx context.Context, id string) (*Account, error)
}
