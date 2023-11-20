package log

import (
	"github.com/Sellsuki/sellsuki-go-logger/level"
	"go.uber.org/zap/zapcore"
)

type Log interface {
	Write()                                // Logs the current entry to the output.
	SetMessage(msg string) Log             // Sets or overrides the log message.
	SetLevel(level level.Level) Log        // Sets or overrides the log level (e.g., info, warning, error).
	SetAlert(bool bool) Log                // Sets or overrides the alert flag.
	WithAppData(key string, value any) Log // Adds application-specific data.
	WithError(err error) Log               // Adds error information.
	WithTracing(t SpanContext) Log         // Adds tracing information.
	WithStackTrace() Log                   // Captures and adds a stack trace.
}

type ZapLogger interface {
	Log(level zapcore.Level, msg string, fields ...zapcore.Field)
}
