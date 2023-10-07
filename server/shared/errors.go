package shared

import (
	"database/sql"
	"strings"
)

type ErrorMessage string

func (e ErrorMessage) String() string {
	return string(e)
}

const (
	ErrorMissingParam   = ErrorMessage("invalid/missing param in request path")
	ErrorInvalidRequest = ErrorMessage("invalid/missing field in request payload")
)

func IsErrNoRows(err error) bool {
	return strings.Contains(err.Error(), sql.ErrNoRows.Error())
}
