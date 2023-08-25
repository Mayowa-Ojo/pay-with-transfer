package store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type DataStore struct {
	db *sqlx.DB
}

func New() AccountStore {
	return &DataStore{}
}

func (d *DataStore) GetAccountByID(ctx context.Context, id string) (*Account, error) {
	rows, err := d.db.Queryx("SELECT * FROM service.accounts WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var account Account
	for rows.Next() {
		err = rows.StructScan(&account)
		if err != nil {
			return nil, err
		}
	}
	return &account, nil
}
