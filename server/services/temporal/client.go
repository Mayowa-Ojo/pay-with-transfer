package temporal

import (
	"context"
	"fmt"
	"os"
	"pay-with-transfer/shared/logger"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

func NewClient() client.Client {
	c, err := client.Dial(client.Options{
		Namespace: NAMESPACE,
	})
	if err != nil {
		logger.WithError(err).Error("unable to create temporal client")
		os.Exit(1)
	}

	logger.Info("connected to temporal client ✅")
	return c
}

func Init(ctx context.Context, tc client.Client) {
	_, err := tc.ExecuteWorkflow(ctx, client.StartWorkflowOptions{
		ID:           fmt.Sprintf("temporal.workflow.account-reset:%s", uuid.NewString()[:5]),
		TaskQueue:    TASK_QUEUE,
		CronSchedule: "0-59/2 * * * *",
	}, HandleAccountResetWorkflow)
	if err != nil {
		logger.WithError(err).Error("failed to start workflow execution")
		os.Exit(1)
	}

	logger.Info("started workflow execution ✅")
}
