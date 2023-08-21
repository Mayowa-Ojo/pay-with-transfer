package logger

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
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

func WithError(err error) {
	std.With(LOG_FIELD_ERROR, err)
}
