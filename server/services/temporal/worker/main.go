package main

import (
	"os"
	"pay-with-transfer/config"
	"pay-with-transfer/services/temporal"
	"pay-with-transfer/shared/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	logger.Info("starting temporal worker...")

	cfg, err := config.Load()
	if err != nil {
		logger.Error("failed to load config data")
		os.Exit(1)
	}

	db, err := sqlx.Open(config.DATABASE_DRIVER, cfg.Database.GetURI())
	if err != nil {
		logger.WithError(err).Error("failed to connect to database")
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		logger.WithError(err).Error("failed to connect to database")
		os.Exit(1)
	}

	logger.Info("connected to database ✅")

	// Initialize a Temporal Client
	clientOptions := client.Options{
		Namespace: temporal.NAMESPACE,
	}
	c, err := client.Dial(clientOptions)
	if err != nil {
		logger.WithError(err).Error("unable to create Temporal Client")
		os.Exit(1)
	}
	defer c.Close()
	// Create a new Worker
	w := worker.New(c, temporal.TASK_QUEUE, worker.Options{})
	// Register Workflows
	w.RegisterWorkflow(temporal.HandleAccountResetWorkflow)
	// Register Acivities
	a := &temporal.AccountActivity{}
	a.WithDatastore(db)
	w.RegisterActivity(a)
	// Start the the Worker Process
	err = w.Run(worker.InterruptCh())
	if err != nil {
		logger.WithError(err).Error("unable to start the Worker Process")
		os.Exit(1)
	} else {
		logger.Info("started temporal worker ✅")
	}
}
