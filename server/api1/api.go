package api1

import (
	"pay-with-transfer/healthcheck"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	healthsvc := healthcheck.New()
	healthFacade := NewHealthFacade(healthsvc)

	v1.GET("/health/check", healthFacade.GetHealthStatus)
}
