package cmd

import (
	"fmt"
	"net/http"
	"os"
	"pay-with-transfer/api1"
	"pay-with-transfer/config"
	"pay-with-transfer/shared/logger"

	"github.com/gin-gonic/gin"
)

func Execute() {
	cfg, err := config.Load()
	if err != nil {
		logger.Error("failed to load config data")
		os.Exit(1)
	}
	server := gin.Default()

	api1.Init(server)

	server.NoRoute(func(ctx *gin.Context) {
		ctx.Status(http.StatusNotFound)
		ctx.Writer.Write([]byte("Requested resource not found"))
	})

	if err := server.Run(fmt.Sprintf(":%s", cfg.App.Port)); err != nil {
		logger.Error("failed to start server")
		panic(err)
	}
}
