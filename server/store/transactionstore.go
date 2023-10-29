package store

import (
	"context"
	"database/sql"
	"pay-with-transfer/shared"

	"github.com/volatiletech/null/v8"
)

func (d *DataStore) GetTransactionByID(ctx context.Context, id string) (*Transaction, error) {
	rows, err := d.db.Queryx(shared.BindReplacer(`SELECT * FROM service.transactions WHERE id = ?`), id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	txn := Transaction{}
	for rows.Next() {
		err = rows.StructScan(&txn)
		if err != nil {
			return nil, err
		}
	}
	if txn.IsEmpty() {
		return nil, sql.ErrNoRows
	}
	return &txn, nil
}

func (d *DataStore) CreateTransaction(ctx context.Context, t Transaction) error {
	if t.ProviderResponse.String == "" {
		t.ProviderResponse = null.NewString("{}", true)
	}
	stmt := `INSERT INTO service.transactions
	(id, account_id, ephemeral_account_id, external_id, amount, currency, account_name, account_number, bank_name, status, provider, provider_response, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := d.db.ExecContext(ctx, shared.BindReplacer(stmt),
		t.ID.String(),
		t.AccountID.String(),
		t.EphemeralAccountID.String(),
		t.ExternalID.String,
		t.Amount,
		t.Currency,
		t.AccountName.String,
		t.AccountNumber.String,
		t.BankName.String,
		t.Status.String(),
		t.Provider.String,
		t.ProviderResponse.String,
		t.CreatedAt,
		t.UpdatedAt,
	)
	if err != nil {
		return err
	}

	result.RowsAffected()
	return nil
}

func (d *DataStore) UpdateTransaction(ctx context.Context, t Transaction) error {
	if t.ProviderResponse.String == "" {
		t.ProviderResponse = null.NewString("{}", true)
	}
	stmt := `UPDATE service.transaction SET
	account_id = ?, ephemeral_account_id = ?, external_id = ?, amount = ?, currency = ?, account_name = ?,
	account_number = ?, bank_name = ?, status = ?, provider = ?, provider_response = ?, created_at = ?, updated_at = ?
	WHERE id = ?`
	result, err := d.db.ExecContext(ctx, shared.BindReplacer(stmt),
		t.AccountID.String(),
		t.EphemeralAccountID.String(),
		t.ExternalID.String,
		t.Amount,
		t.Currency,
		t.AccountName.String,
		t.AccountNumber.String,
		t.BankName.String,
		t.Status.String(),
		t.Provider.String,
		t.ProviderResponse.String,
		t.CreatedAt,
		t.UpdatedAt,
	)
	if err != nil {
		return err
	}

	result.RowsAffected()
	return nil
}
