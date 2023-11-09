package log

import (
	"github.com/Sellsuki/sellsuki-go-logger/config"
	"github.com/Sellsuki/sellsuki-go-logger/level"
	"go.uber.org/zap"
)

type Type string

const (
	TypeAudit       Type = "audit"
	TypeEvent       Type = "event"
	TypeApplication Type = "application"
)

type Base struct {
	logger zapLogger
	config config.Config

	Type      Type
	Level     level.Level
	Alert     bool
	Message   string
	Data      map[string]any
	AppFields map[string]any
}

func (l Base) Write() {
	if len(l.AppFields) > 0 {
		l.Data[l.config.AppName] = l.AppFields
	}

	f := []zap.Field{
		zap.String("app_name", l.config.AppName),
		zap.String("version", l.config.Version),
		zap.Int("alert", boolToInt[l.Alert]),
		zap.String("log_type", string(l.Type)),
		zap.Any("data", l.Data),
	}

	l.logger.Log(level.ToZap(l.Level), l.Message, f...)
}

func (l Base) SetMessage(msg string) Log {
	l.Message = msg
	return l
}

func (l Base) SetLevel(level level.Level) Log {
	l.Level = level
	return l
}

func (l Base) SetAlert(bool bool) Log {
	l.Alert = bool
	return l
}

func (l Base) WithAppData(key string, value any) Log {
	l.AppFields[key] = value

	return l
}

func (l Base) WithError(err error) Log {
	return l.WithField("error", err)
}

func (l Base) WithTracing(t Tracer) Log {
	return l.WithField("tracing", map[string]string{
		"trace_id": t.TraceID().String(),
		"span_id":  t.SpanID().String(),
	})
}

func (l Base) WithField(key string, value any) Log {
	l.Data[key] = value
	return l
}

func (l Base) WithFields(fields map[string]any) Log {
	for k, v := range fields {
		l.Data[k] = v
	}
	return l
}

func (l Base) WithStackTrace() Log {
	return l.WithField("stack_trace", captureStackTrace(2))
}

func (l Base) WithHTTPReq(req HTTPRequestPayload) Log {
	if l.config.MaxBodySize > 0 && len(req.Body) > l.config.MaxBodySize {
		req.Body = req.Body[:l.config.MaxBodySize]
	}

	l.Data["http_request"] = req

	return l
}

func (l Base) WithHTTPResp(resp HTTPResponsePayload) Log {
	if l.config.MaxBodySize > 0 && len(resp.Body) > l.config.MaxBodySize {
		resp.Body = resp.Body[:l.config.MaxBodySize]
	}

	l.Data["http_response"] = resp

	return l
}

func (l Base) WithKafkaMessage(msg KafkaMessagePayload) Log {
	if l.config.MaxBodySize > 0 && len(msg.Payload) > l.config.MaxBodySize {
		msg.Payload = msg.Payload[:l.config.MaxBodySize]
	}

	l.Data["kafka_message"] = msg

	return l
}

func (l Base) WithKafkaResult(result KafkaResultPayload) Log {
	l.Data["kafka_result"] = result

	return l
}

func New(logger *zap.Logger, cfg config.Config, l level.Level, t Type) Base {
	return Base{
		logger:    logger,
		config:    cfg,
		Data:      map[string]any{},
		AppFields: map[string]any{},
		Level:     l,
		Type:      t,
	}
}
