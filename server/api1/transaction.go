package api1

import (
	"net/http"
	"pay-with-transfer/shared"
	paylog "pay-with-transfer/shared/logger"
	"pay-with-transfer/transactions"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionFacade struct {
	svc transactions.TransactionService
}

func NewTransactionFacade(svc transactions.TransactionService) *TransactionFacade {
	return &TransactionFacade{
		svc,
	}
}

func (f *TransactionFacade) FetchSingleTransaction(c *gin.Context) {
	logger := paylog.WithTrace(c).With(paylog.LOG_FIELD_FUNCTION_NAME, "FetchSingleTransaction")

	transactionId := c.Param("id")
	if transactionId == "" {
		logger.Error("missing param [id] in request path")
		resp := shared.GetResponse(shared.ResponseCodeError, shared.ErrorMissingParam.String(), nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	time.Sleep(time.Millisecond * 3000)

	txn, err := f.svc.FetchSingleTransaction(c, transactionId)
	if err != nil {
		logger.With(paylog.LOG_FIELD_ERROR, err).Error("failed to fetch single transaction")
		resp := shared.GetResponse(shared.ResponseCodeError, err.Error(), nil)
		c.JSON(http.StatusPreconditionFailed, resp)
		return
	}

	c.JSON(http.StatusOK, shared.GetResponse(shared.ResponseCodeOk, "success", txn))
}
