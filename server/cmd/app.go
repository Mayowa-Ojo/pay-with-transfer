package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"pay-with-transfer/api1"
	"pay-with-transfer/cache"
	paycache "pay-with-transfer/cache"
	"pay-with-transfer/config"
	"pay-with-transfer/shared/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Execute() {
	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		logger.Error("failed to load config data")
		os.Exit(1)
	}

	server := gin.Default()
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

	cache := paycache.New(ctx, cfg.Redis.Host, cfg.Redis.Port,
		cache.WithNamespace(cfg.Redis.Namespace),
	)

	if _, err := cache.Ping(ctx); err != nil {
		logger.WithError(err).Error("failed to connect to redis")
		os.Exit(1)
	}

	logger.Info("connected to redis ✅")

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api1.Init(server, db, cache, cfg)

	server.NoRoute(func(ctx *gin.Context) {
		ctx.Status(http.StatusNotFound)
		ctx.Writer.Write([]byte("Requested resource not found"))
	})

	if err := server.Run(fmt.Sprintf(":%s", cfg.App.Port)); err != nil {
		logger.Error("failed to start server")
		panic(err)
	}
}
