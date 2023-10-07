package paystack

import (
	"fmt"
	"strconv"
	"time"
)

const (
	ASSIGN_VIRTUAL_ACCOUNT_PATH = "/dedicated_account/assign"
)

const BANK_NAME_WEMA = "Wema Bank"
const BANK_ID_WEMA = 20780
const BANK_SLUG_WEMA = "wema-bank"

const CURRENCY_NGN = "NGN"

type AssignVirtualAccountRequest struct {
	Email         string `json:"email" validate:"required"`
	FirstName     string `json:"first_name" validate:"required"`
	MiddleName    string `json:"middle_name" validate:"required"`
	LastName      string `json:"last_name" validate:"required"`
	Phone         string `json:"phone" validate:"required"`
	PreferredBank string `json:"preferred_bank" validate:"required"`
	Country       string `json:"country" validate:"required"`
	AccountNumber string `json:"account_number"`
	BankCode      string `json:"bank_code"`
	BVN           string `json:"bvn"`
	Subaccount    string `json:"subaccount"`
	SplitCode     string `json:"split_code"`
}

type Customer struct {
	ID                       int         `json:"id"`
	FirstName                string      `json:"first_name"`
	LastName                 string      `json:"last_name"`
	Email                    string      `json:"email"`
	CustomerCode             string      `json:"customer_code"`
	Phone                    string      `json:"phone"`
	Metadata                 interface{} `json:"metadata"`
	RiskAction               string      `json:"risk_action"`
	InternationalFormatPhone string      `json:"international_format_phone"`
}
type Bank struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
	Slug string `json:"slug"`
}
type Assignment struct {
	Integration  int         `json:"integration"`
	AssigneeID   int         `json:"assignee_id"`
	AssigneeType string      `json:"assignee_type"`
	Expired      bool        `json:"expired"`
	AccountType  string      `json:"account_type"`
	AssignedAt   time.Time   `json:"assigned_at"`
	ExpiredAt    interface{} `json:"expired_at"`
}
type DedicatedAccount struct {
	Bank          Bank        `json:"bank"`
	AccountName   string      `json:"account_name"`
	AccountNumber string      `json:"account_number"`
	Assigned      bool        `json:"assigned"`
	Currency      string      `json:"currency"`
	Metadata      interface{} `json:"metadata"`
	Active        bool        `json:"active"`
	ID            int         `json:"id"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	Assignment    Assignment  `json:"assignment"`
}
type Identification struct {
	Status string `json:"status"`
}
type Data struct {
	Customer         Customer         `json:"customer"`
	DedicatedAccount DedicatedAccount `json:"dedicated_account"`
	Identification   Identification   `json:"identification"`
}
type AssignVirtualAccountResponse struct {
	Status  bool   `json:"status"`
	Data    Data   `json:"data"`
	Message string `json:"message"`
}

func (r *AssignVirtualAccountResponse) WithAccountID() {
	if r == nil {
		return
	}
	r.Data.DedicatedAccount.ID, _ = strconv.Atoi(strconv.Itoa(int(time.Now().UnixMilli()))[7:])
}

func (r *AssignVirtualAccountResponse) WithCustomerID() {
	if r == nil {
		return
	}
	r.Data.Customer.ID, _ = strconv.Atoi(strconv.Itoa(int(time.Now().UnixMilli()))[7:])
}

func (r *AssignVirtualAccountResponse) WithCustomerCode() {
	if r == nil {
		return
	}
	r.Data.Customer.CustomerCode = fmt.Sprintf("cus-%s", strconv.Itoa(int(time.Now().UnixMilli()))[3:])
}
