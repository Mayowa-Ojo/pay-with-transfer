package store

import "github.com/google/uuid"

type Account struct {
	ID              uuid.UUID `db:"id" json:"id"`
	AccountHolderID uuid.UUID `db:"account_holder_id" json:"account_holder_id"`
	AccountName     string    `db:"account_name" json:"account_name"`
	AccountNumber   string    `db:"account_number" json:"account_number"`
	BankName        string    `db:"bank_name" json:"bank_name"`
	Currency        string    `db:"currency" json:"currency"`
	ProviderID      string    `db:"provider_id" json:"provider_id"`
	Provider        string    `db:"provider" json:"provider"`
	IsActive        string    `db:"is_active" json:"is_active"`
	CreatedAt       string    `db:"created_at" json:"created_at"`
	UpdatedAt       string    `db:"updated_at" json:"updated_at"`
}

type AccountHolder struct {
	ID           uuid.UUID `db:"id" json:"id"`
	FirstName    string    `db:"first_name" json:"first_name"`
	LastNumber   string    `db:"last_name" json:"last_name"`
	Email        string    `db:"email" json:"email"`
	Phone        string    `db:"phone" json:"phone"`
	ProviderID   string    `db:"provider_id" json:"provider_id"`
	ProviderCode string    `db:"provider_code" json:"provider_code"`
	CreatedAt    string    `db:"created_at" json:"created_at"`
	UpdatedAt    string    `db:"updated_at" json:"updated_at"`
}
