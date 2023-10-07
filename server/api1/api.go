package api1

import (
	"pay-with-transfer/accounts"
	paycache "pay-with-transfer/cache"
	"pay-with-transfer/config"
	"pay-with-transfer/healthcheck"
	"pay-with-transfer/store"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Init(router *gin.Engine, db *sqlx.DB, cache paycache.Cache, cfg *config.Config) {
	v1 := router.Group("/api/v1")

	store := store.New(db)

	healthsvc := healthcheck.New()
	healthFacade := NewHealthFacade(healthsvc)

	accountsvc := accounts.New(store, cache, *cfg)
	accountFacade := NewAccountFacade(accountsvc)

	v1.GET("/health/check", healthFacade.GetHealthStatus)
	v1.GET("/accounts/:id", accountFacade.FetchSingleAccount)
	v1.POST("/accounts/pool/generate", accountFacade.GeneratePoolAccounts)
	v1.POST("/accounts/ephemeral", accountFacade.CreateEphemeralAccount)
}
