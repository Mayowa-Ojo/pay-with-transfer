package api1

import (
	"net/http"
	"pay-with-transfer/healthcheck"
	"pay-with-transfer/shared/logger"

	"github.com/gin-gonic/gin"
)

type HealthFacade struct {
	svc healthcheck.HealthService
}

func NewHealthFacade(svc healthcheck.HealthService) *HealthFacade {
	return &HealthFacade{
		svc,
	}
}

func (f *HealthFacade) GetHealthStatus(c *gin.Context) {
	logger.WithTrace(c).With(logger.LOG_FIELD_FUNCTION_NAME, "GetHealthStatus")

	logger.Info("checking health status...")

	c.JSON(http.StatusOK, gin.H{"data": nil})
}
