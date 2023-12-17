package temporal

import "pay-with-transfer/store"

const (
	NAMESPACE                 = "pay-with-transfer"
	TASK_QUEUE                = "pay-with-transfer-task-queue"
	ACCOUNT_RESET_SCHEDULE_ID = "account-reset-schedule"
	ACCOUNT_RESET_WORKFLOW_ID = "account-reset-workflow"
)

//NB: struct fields should be exported because it's going to serialized when moving between worker/server

type GetActiveAccountsResponse struct {
	Accounts []*store.EphemeralAccount
}

type UpdateEphemeralAccountParam struct {
	Account store.EphemeralAccount
}
