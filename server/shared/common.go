package shared

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	DEFAULT_ACCOUNT_HOLDER_FIRST_NAME = "PAY"
	DEFAULT_ACCOUNT_HOLDER_LAST_NAME  = "CHECKOUT"
)

func GeneratePayEmail() string {
	return fmt.Sprintf("checkout.%s@pay.com", uuid.NewString()[:4])
}

func GeneratePayPhoneNumber() string {
	return fmt.Sprintf("+23480%s", strconv.Itoa(int(time.Now().UnixMilli()))[5:])
}

func BindReplacer(query string, args ...interface{}) string {
	startAt := 1

	paramBuf := bytes.Buffer{}
	paramIndex := 0

	for {
		if paramIndex >= len(query) {
			break
		}

		query = query[paramIndex:]
		paramIndex = strings.IndexByte(query, '?')

		if paramIndex == -1 {
			paramBuf.WriteString(query)
			break
		}

		escapeIndex := strings.Index(query, `\?`)
		if escapeIndex != -1 && paramIndex > escapeIndex {
			paramBuf.WriteString(query[:escapeIndex] + "?")
			paramIndex++
			continue
		}

		paramBuf.WriteString(query[:paramIndex] + fmt.Sprintf("$%d", startAt))
		startAt++
		paramIndex++
	}

	return paramBuf.String()
}
