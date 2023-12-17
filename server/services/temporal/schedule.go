package temporal

import (
	"context"
	"os"
	"pay-with-transfer/shared/logger"
	"time"

	"go.temporal.io/sdk/client"
)

func CreateAccountResetSchedule(c client.ScheduleClient) {
	action := &client.ScheduleWorkflowAction{
		ID:                 ACCOUNT_RESET_WORKFLOW_ID,
		Workflow:           nil,
		TaskQueue:          TASK_QUEUE,
		WorkflowRunTimeout: 30 * time.Second,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := c.Create(ctx, client.ScheduleOptions{
		ID:                 ACCOUNT_RESET_SCHEDULE_ID,
		TriggerImmediately: true,
		Spec:               client.ScheduleSpec{},
		Action:             action,
	})

	if err != nil {
		logger.WithError(err).Error("unable to create schedule: %s", ACCOUNT_RESET_SCHEDULE_ID)
		os.Exit(1)
	}
	logger.Info("create schedule: %s successfully", ACCOUNT_RESET_SCHEDULE_ID)
	return
}
