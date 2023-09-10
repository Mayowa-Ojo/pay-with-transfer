package logger

import (
	"context"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
)

const (
	LOG_FIELD_FUNCTION_NAME = "functionName"
	LOG_FIELD_ERROR         = "error"
	LOG_FIELD_TRACE_ID      = "traceId"
)

var std *log.Logger

func init() {
	std = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
	})
}

func With(keyvals ...interface{}) *log.Logger {
	return std.With(keyvals...)
}

func Info(msg interface{}, keyvals ...interface{}) {
	std.Info(msg, keyvals...)
}

func Error(msg interface{}, keyvals ...interface{}) {
	std.Error(msg, keyvals...)
}

func Warn(msg interface{}, keyvals ...interface{}) {
	std.Warn(msg, keyvals...)
}

func WithError(err error) *log.Logger {
	return std.With(LOG_FIELD_ERROR, err)
}

func WithTrace(ctx context.Context) *log.Logger {
	v := ctx.Value(LOG_FIELD_TRACE_ID)
	id, ok := v.(string)
	if ok {
		return std.With(LOG_FIELD_TRACE_ID, id)
	}
	return std.With(LOG_FIELD_TRACE_ID, uuid.NewString())
}
