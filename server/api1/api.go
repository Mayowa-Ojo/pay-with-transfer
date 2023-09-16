package api1

import (
	"pay-with-transfer/accounts"
	"pay-with-transfer/healthcheck"
	"pay-with-transfer/store"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Init(router *gin.Engine, db *sqlx.DB) {
	v1 := router.Group("/api/v1")

	store := store.New(db)

	healthsvc := healthcheck.New()
	healthFacade := NewHealthFacade(healthsvc)

	accountsvc := accounts.New(store)
	accountFacade := NewAccountFacade(accountsvc)

	v1.GET("/health/check", healthFacade.GetHealthStatus)
	v1.GET("/accounts/:id", accountFacade.FetchSingleAccount)
	v1.POST("/accounts/pool/generate", accountFacade.GeneratePoolAccounts)
}
