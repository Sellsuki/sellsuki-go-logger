package v2

import (
	"github.com/Sellsuki/sellsuki-go-logger/config"
	"github.com/Sellsuki/sellsuki-go-logger/level"
	"github.com/Sellsuki/sellsuki-go-logger/log"
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
	return l.WithField("error", err)
}

func (l Logger) WithTracing(t log.Tracer) log.Log {
	return l.WithField("tracing", map[string]string{
		"trace_id": t.TraceID().String(),
		"span_id":  t.SpanID().String(),
	})
}

func (l Logger) WithField(key string, value any) log.Log {
	l.Data[key] = value
	return &l
}

func (l Logger) WithFields(fields map[string]any) log.Log {
	for k, v := range fields {
		l.Data[k] = v
	}
	return &l
}

func (l Logger) WithStackTrace() log.Log {
	return l.WithField("stack_trace", CaptureStackTrace(2))
}

func (l Logger) WithHTTPReq(req log.HTTPRequestPayload) log.Log {
	if l.config.MaxBodySize > 0 && len(req.Body) > l.config.MaxBodySize {
		req.Body = req.Body[:l.config.MaxBodySize]
	}

	l.Data["http_request"] = req

	return &l
}

func (l Logger) WithHTTPResp(resp log.HTTPResponsePayload) log.Log {
	if l.config.MaxBodySize > 0 && len(resp.Body) > l.config.MaxBodySize {
		resp.Body = resp.Body[:l.config.MaxBodySize]
	}

	l.Data["http_response"] = resp

	return &l
}

func (l Logger) WithKafkaMessage(msg log.KafkaMessagePayload) log.Log {
	if l.config.MaxBodySize > 0 && len(msg.Payload) > l.config.MaxBodySize {
		msg.Payload = msg.Payload[:l.config.MaxBodySize]
	}

	l.Data["kafka_message"] = msg

	return &l
}

func (l Logger) WithKafkaResult(result log.KafkaResultPayload) log.Log {
	l.Data["kafka_result"] = result

	return &l
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