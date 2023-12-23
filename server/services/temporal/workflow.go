package temporal

import (
	"pay-with-transfer/store"
	"time"

	"go.temporal.io/sdk/workflow"
)

func HandleAccountResetWorkflow(ctx workflow.Context) error {
	workflow.GetLogger(ctx).Info("AccountResetWorkflow workflow started", "StartTime", workflow.Now(ctx))

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}

	ctx = workflow.WithActivityOptions(ctx, ao)

	var a *AccountActivity
	var accounts []*store.EphemeralAccount
	{
		var resp GetActiveAccountsResponse
		err := workflow.ExecuteActivity(ctx, a.GetActiveAccounts).Get(ctx, &resp)
		if err != nil {
			workflow.GetLogger(ctx).Error("failed to execute activity", "Error", err)
			return err
		}
		accounts = resp.Accounts
	}

	for _, v := range accounts {
		v.Status = store.EphemeralAccountExpired
		err := workflow.ExecuteActivity(ctx, a.UpdateEphemeralAccount, UpdateEphemeralAccountParam{
			Account: *v,
		}).Get(ctx, nil)
		if err != nil {
			workflow.GetLogger(ctx).Error("failed to execute activity", "Error", err)
			continue
		}

		//check if transaction is still pending
		var resp GetEphemeralAccountTransactionResponse
		err = workflow.ExecuteActivity(ctx, a.GetEphemeralAccountTransaction, GetEphemeralAccountTransactionParam{
			AccountID: v.ID.String(),
		}).Get(ctx, &resp)
		if err != nil {
			workflow.GetLogger(ctx).Error("failed to execute activity", "Error", err)
			continue
		}

		if resp.Transaction.Status == store.TransactionPending {
			resp.Transaction.Status = store.TransactionCanceled
			err = workflow.ExecuteActivity(ctx, a.UpdateTransaction, UpdateTransactionParam{
				Transaction: resp.Transaction,
			}).Get(ctx, nil)
			if err != nil {
				workflow.GetLogger(ctx).Error("failed to execute activity", "Error", err)
				continue
			}
		}
	}

	workflow.GetLogger(ctx).Info("AccountResetWorkflow workflow completed", "CompletedTime", workflow.Now(ctx))

	return nil
}
