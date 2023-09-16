package paystack

const (
	ASSIGN_VIRTUAL_ACCOUNT_PATH = "/dedicated_account/assign"
)

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

type AssignVirtualAccountResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
