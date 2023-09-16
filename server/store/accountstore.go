package store

import (
	"context"
	"fmt"
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

func (d *DataStore) CreateAccountHolder(ctx context.Context, ah AccountHolder) error {
	fmt.Printf("\n\n[LOG]: creating account...\n\n")
	stmt := `INSERT INTO service.account_holders 
	(id, first_name, last_name, email, phone, provider_id, provider_code, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := d.db.ExecContext(ctx, shared.BindReplacer(stmt),
		ah.ID.String(),
		ah.FirstName,
		ah.LastName,
		ah.Email,
		ah.Phone,
		ah.ProviderID,
		ah.ProviderCode,
		ah.CreatedAt,
		ah.UpdatedAt,
	)
	if err != nil {
		return err
	}

	result.RowsAffected()
	return nil
}
