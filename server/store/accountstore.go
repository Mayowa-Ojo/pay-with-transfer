package store

import (
	"context"
	"database/sql"
	"pay-with-transfer/shared"

	"github.com/volatiletech/null/v8"
)

func (d *DataStore) GetAccountByID(ctx context.Context, id string) (*Account, error) {
	rows, err := d.db.Queryx(shared.BindReplacer(`SELECT * FROM service.accounts WHERE id = ?`), id)
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
	if account.IsEmpty() {
		return nil, sql.ErrNoRows
	}
	return &account, nil
}

func (d *DataStore) UpdateAccount(ctx context.Context, ac Account) error {
	if ac.ProviderResponse.String == "" {
		ac.ProviderResponse = null.NewString("{}", true)
	}
	stmt := `UPDATE service.accounts SET
	account_holder_id = ?, account_name = ?, account_number = ?, bank_name = ?, currency = ?, provider_id = ?,
	provider = ?, is_active = ?, is_dormant = ?, provider_response = ?, updated_at = ?
	WHERE id = ?`
	result, err := d.db.ExecContext(ctx, shared.BindReplacer(stmt),
		ac.AccountHolderID.String(),
		ac.AccountName,
		ac.AccountNumber,
		ac.BankName.String,
		ac.Currency,
		ac.ProviderID,
		ac.Provider,
		ac.IsActive,
		ac.IsDormant,
		ac.ProviderResponse.String,
		ac.UpdatedAt,
		ac.ID.String(),
	)
	if err != nil {
		return err
	}

	result.RowsAffected()
	return nil
}

func (d *DataStore) CreateAccountHolder(ctx context.Context, ah AccountHolder) error {
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

func (d *DataStore) FindDormantAccount(ctx context.Context) (*Account, error) {
	rows, err := d.db.Queryx(`SELECT * FROM service.accounts WHERE is_active = TRUE AND is_dormant = TRUE LIMIT 1`)
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
	if account.IsEmpty() {
		return nil, sql.ErrNoRows
	}
	return &account, nil
}

func (d *DataStore) FindEphemeralAccountByAccountID(ctx context.Context, accountID string) (*EphemeralAccount, error) {
	query := `SELECT * FROM service.ephemeral_accounts WHERE account_id = ? ORDER BY created_at DESC LIMIT 1`
	rows, err := d.db.Queryx(shared.BindReplacer(query), accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	account := EphemeralAccount{}
	for rows.Next() {
		err = rows.StructScan(&account)
		if err != nil {
			return nil, err
		}
	}
	if account.IsEmpty() {
		return nil, sql.ErrNoRows
	}
	return &account, nil
}

func (d *DataStore) CreateEphemeralAccount(ctx context.Context, ea EphemeralAccount) error {
	stmt := `INSERT INTO service.ephemeral_accounts 
	(id, account_id, amount, status, expires_at, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?)`
	result, err := d.db.ExecContext(ctx, shared.BindReplacer(stmt),
		ea.ID.String(),
		ea.AccountID.String(),
		ea.Amount,
		ea.Status.String(),
		ea.ExpiresAt,
		ea.CreatedAt,
		ea.UpdatedAt,
	)
	if err != nil {
		return err
	}

	result.RowsAffected()
	return nil
}
