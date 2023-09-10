package shared

import (
	"bytes"
	"fmt"
	"strings"
)

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
