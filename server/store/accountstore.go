package store

import (
	"context"
	"pay-with-transfer/shared"
)

func (d *DataStore) GetAccountByID(ctx context.Context, id string) (*Account, error) {
	rows, err := d.db.Queryx(shared.BindReplacer("SELECT * FROM service.accounts WHERE id = ?"), id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	account := Account{}
	for rows.Next() {
		err = rows.StructScan(&account)
		if err != nil {
			return nil, err
		}
	}
	return &account, nil
}
