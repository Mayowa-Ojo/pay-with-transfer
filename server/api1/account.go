package api1

import (
	"net/http"
	"pay-with-transfer/accounts"
	"pay-with-transfer/shared"
	paylog "pay-with-transfer/shared/logger"

	"github.com/gin-gonic/gin"
)

type AccountFacade struct {
	svc accounts.AccountService
}

func NewAccountFacade(svc accounts.AccountService) *AccountFacade {
	return &AccountFacade{
		svc,
	}
}

func (f *AccountFacade) FetchSingleAccount(c *gin.Context) {
	logger := paylog.WithTrace(c).With(paylog.LOG_FIELD_FUNCTION_NAME, "FetchSingleAccount")

	accountID := c.Param("id")
	if accountID == "" {
		logger.Error("missing param [id] in request path")
		resp := shared.GetResponse(shared.ResponseCodeError, shared.ErrorMissingParam.String(), nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	account, err := f.svc.FetchSingleAccount(c, accountID)
	if err != nil {
		logger.With(paylog.LOG_FIELD_ERROR, err).Error("failed to fetch single account")
		resp := shared.GetResponse(shared.ResponseCodeError, err.Error(), nil)
		c.JSON(http.StatusPreconditionFailed, resp)
		return
	}

	c.JSON(http.StatusOK, shared.GetResponse(shared.ResponseCodeOk, "success", account))
}

func (f *AccountFacade) GeneratePoolAccounts(c *gin.Context) {
	logger := paylog.WithTrace(c).With(paylog.LOG_FIELD_FUNCTION_NAME, "GeneratePoolAccounts")

	var dto accounts.GeneratePoolAccountsDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		logger.Error("failed to bind request body")
		resp := shared.GetResponse(shared.ResponseCodeError, shared.ErrorInvalidRequest.String(), nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if err := f.svc.GeneratePoolAccounts(c, dto.Count); err != nil {
		logger.With(paylog.LOG_FIELD_ERROR, err).Error("failed to generate pool accounts")
		resp := shared.GetResponse(shared.ResponseCodeError, err.Error(), nil)
		c.JSON(http.StatusPreconditionFailed, resp)
		return
	}

	c.JSON(http.StatusOK, shared.GetResponse(shared.ResponseCodeOk, "success", nil))
}
