package zap_logger

import (
	"github.com/Sellsuki/sellsuki-go-logger/v2/config"
	"github.com/Sellsuki/sellsuki-go-logger/v2/level"
	"github.com/Sellsuki/sellsuki-go-logger/v2/log"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Type string

const (
	TypeAudit        Type = "audit"
	TypeEvent        Type = "event"
	TypeApplication  Type = "application"
	TypeHandlerKafka Type = "handler.kafka"
	TypeHandlerHTTP  Type = "handler.http"
)

type Logger struct {
	logger log.ZapLogger
	config config.Config

	Type      Type
	Level     level.Level
	Alert     bool
	Message   string
	Data      map[string]any
	AppFields map[string]any
}

func (l Logger) Write() {
	if len(l.AppFields) > 0 {
		l.Data[l.config.AppName] = l.AppFields
	}

	f := []zap.Field{
		zap.String("app_name", l.config.AppName),
		zap.String("version", l.config.Version),
		zap.Int("alert", BoolToInt[l.Alert]),
		zap.String("log_type", string(l.Type)),
		zap.Any("data", l.Data),
	}

	l.logger.Log(level.ToZap(l.Level), l.Message, f...)
}

func (l Logger) SetMessage(msg string) log.Log {
	l.Message = msg
	return &l
}

func (l Logger) SetLevel(level level.Level) log.Log {
	l.Level = level
	return &l
}

func (l Logger) SetAlert(bool bool) log.Log {
	l.Alert = bool
	return &l
}

func (l Logger) WithAppData(key string, value any) log.Log {
	l.AppFields[key] = value

	return &l
}

func (l Logger) WithError(err error) log.Log {
	if err == nil {
		return &l
	}

	return l.WithField("error", err)
}

func (l Logger) WithTracing(sc trace.SpanContext) log.Log {
	return l.WithField("tracing", map[string]string{
		"trace_id": sc.TraceID().String(),
		"span_id":  sc.SpanID().String(),
	})
}

// WithField adds a single field to the log entry.
// for internal use only
func (l Logger) WithField(key string, value any) log.Log {
	l.Data[key] = value
	return &l
}

// WithFields adds multiple fields to the log entry.
// for internal use only
func (l Logger) WithFields(fields map[string]any) log.Log {
	for k, v := range fields {
		l.Data[k] = v
	}
	return &l
}

func (l Logger) WithStackTrace() log.Log {
	return l.WithField("stack_trace", CaptureStackTrace(2))
}

func New(logger *zap.Logger, cfg config.Config, l level.Level, t Type, msg string) *Logger {
	return &Logger{
		logger:    logger,
		config:    cfg,
		Data:      map[string]any{},
		AppFields: map[string]any{},
		Level:     l,
		Type:      t,
		Message:   msg,
	}
}
