package api1

import (
	"net/http"
	"pay-with-transfer/healthcheck"
	"pay-with-transfer/shared/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	traceID := uuid.NewString()
	// ctx := context.WithValue(c, logger.LOG_FIELD_TRACE_ID, traceID)
	logger := logger.With(logger.LOG_FIELD_FUNCTION_NAME, "GetHealthStatus").With(logger.LOG_FIELD_TRACE_ID, traceID)

	logger.Info("checking health status...")

	c.JSON(http.StatusOK, gin.H{"data": nil})
}
