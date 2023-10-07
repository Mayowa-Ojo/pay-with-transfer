package backoff

import (
	"context"
	"time"

	paylog "pay-with-transfer/shared/logger"
)

var backoffTimeouts = []time.Duration{200 * time.Millisecond, 400 * time.Millisecond, 400 * time.Millisecond, 600 * time.Millisecond, 600 * time.Millisecond}

// WithRetry executes a BackoffOperation and waits an increasing time before retrying the operation.
func WithRetry(ctx context.Context, op func() error) error {
	logger := paylog.WithTrace(ctx).With(paylog.LOG_FIELD_FUNCTION_NAME, "Backoff.WithRetry")
	var err error

	for attempts := 0; attempts < len(backoffTimeouts); attempts++ {
		err = op()
		if err == nil {
			return nil
		}

		logger.Infof("failed to execute operation, retrying with count: %d", attempts+1)
		time.Sleep(backoffTimeouts[attempts])
	}

	return err
}
