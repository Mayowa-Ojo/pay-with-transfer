package store

import (
	"encoding/json"
	"pay-with-transfer/shared"
	"time"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
)

type Account struct {
	ID               uuid.UUID   `db:"id" json:"id"`
	AccountHolderID  uuid.UUID   `db:"account_holder_id" json:"account_holder_id"`
	AccountName      string      `db:"account_name" json:"account_name"`
	AccountNumber    string      `db:"account_number" json:"account_number"`
	BankName         null.String `db:"bank_name" json:"bank_name"`
	Currency         string      `db:"currency" json:"currency"`
	ProviderID       string      `db:"provider_id" json:"provider_id"`
	Provider         string      `db:"provider" json:"provider"`
	IsActive         bool        `db:"is_active" json:"is_active"`
	IsDormant        bool        `db:"is_dormant" json:"is_dormant"`
	ProviderResponse null.String `db:"provider_response" json:"provider_response"`
	CreatedAt        string      `db:"created_at" json:"created_at"`
	UpdatedAt        string      `db:"updated_at" json:"updated_at"`
}

func (r *Account) IsEmpty() bool {
	var empty Account
	b, _ := json.Marshal(r)
	bb, _ := json.Marshal(empty)
	return string(b) == string(bb)
}

type AccountHolder struct {
	ID               uuid.UUID   `db:"id" json:"id"`
	FirstName        string      `db:"first_name" json:"first_name"`
	LastName         string      `db:"last_name" json:"last_name"`
	Email            string      `db:"email" json:"email"`
	Phone            string      `db:"phone" json:"phone"`
	ProviderID       string      `db:"provider_id" json:"provider_id"`
	ProviderCode     string      `db:"provider_code" json:"provider_code"`
	ProviderResponse null.String `db:"provider_response" json:"provider_response"`
	CreatedAt        time.Time   `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time   `db:"updated_at" json:"updated_at"`
}

func (r *AccountHolder) WithDefaults() {
	r.ID = uuid.New()
	r.FirstName = shared.DEFAULT_ACCOUNT_HOLDER_FIRST_NAME
	r.Email = shared.GeneratePayEmail()
	r.Phone = shared.GeneratePayPhoneNumber()
	r.LastName = shared.DEFAULT_ACCOUNT_HOLDER_LAST_NAME
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}

type EphemeralAccountStatus string

const (
	EphemeralAccountActive  = EphemeralAccountStatus("ACTIVE")
	EphemeralAccountExpired = EphemeralAccountStatus("EXPIRED")
)

func (eas EphemeralAccountStatus) String() string {
	return string(eas)
}

type EphemeralAccount struct {
	ID        uuid.UUID              `db:"id" json:"id"`
	AccountID uuid.UUID              `db:"account_id" json:"-"`
	Amount    int64                  `db:"amount" json:"-"` //amount is in base units
	Status    EphemeralAccountStatus `db:"status" json:"-"`
	ExpiresAt time.Time              `db:"expires_at" json:"expires_at"`
	CreatedAt time.Time              `db:"created_at" json:"created_at"`
	UpdatedAt time.Time              `db:"updated_at" json:"updated_at"`

	AccountName   string  `json:"account_name"`
	AccountNumber string  `json:"account_number"`
	BankName      string  `json:"bank_name"`
	PaymentAmount float64 `json:"payment_amount"`
	Provider      string  `json:"-"`
	TransactionID string  `json:"transaction_id"`
}

func (r *EphemeralAccount) IsEmpty() bool {
	var empty EphemeralAccount
	b, _ := json.Marshal(r)
	bb, _ := json.Marshal(empty)
	return string(b) == string(bb)
}

func (r *EphemeralAccount) WithDefaults() {
	r.ID = uuid.New()
	r.Status = EphemeralAccountActive
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}

type TransactionStatus string

const (
	TransactionSuccessful = TransactionStatus("SUCCESSFUL")
	TransactionPending    = TransactionStatus("PENDING")
	TransactionFailed     = TransactionStatus("FAILED")
)

func (eas TransactionStatus) String() string {
	return string(eas)
}

type Transaction struct {
	ID                 uuid.UUID         `db:"id" json:"id"`
	AccountID          uuid.UUID         `db:"account_id" json:"account_id"`
	EphemeralAccountID uuid.UUID         `db:"ephemeral_account_id" json:"ephemeral_account_id"`
	ExternalID         null.String       `db:"external_id" json:"external_id"`
	Amount             int64             `db:"amount" json:"amount"`
	Currency           string            `db:"currency" json:"currency"`
	AccountName        null.String       `db:"account_name" json:"account_name"`
	AccountNumber      null.String       `db:"account_number" json:"account_number"`
	BankName           null.String       `db:"bank_name" json:"bank_name"`
	Status             TransactionStatus `db:"status" json:"status"`
	Provider           null.String       `db:"provider" json:"provider"`
	ProviderResponse   null.String       `db:"provider_response" json:"provider_response"`
	CreatedAt          time.Time         `db:"created_at" json:"created_at"`
	UpdatedAt          time.Time         `db:"updated_at" json:"updated_at"`
}

func (r *Transaction) IsEmpty() bool {
	var empty Transaction
	b, _ := json.Marshal(r)
	bb, _ := json.Marshal(empty)
	return string(b) == string(bb)
}

func (r *Transaction) WithDefaults() {
	r.ID = uuid.New()
	r.Status = TransactionPending
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}
