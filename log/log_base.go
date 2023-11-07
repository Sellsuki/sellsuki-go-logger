package log

import (
	"context"
	"github.com/Sellsuki/sellsuki-go-logger/config"
	"github.com/Sellsuki/sellsuki-go-logger/level"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type Base struct {
	logger *zap.Logger
	config config.Config

	Level     level.Level
	Alert     bool
	Message   string
	Fields    []zap.Field
	AppFields map[string]any
}

func (l Base) Write() {
	l.Fields = append(l.Fields, zap.Int("alert", boolToInt[l.Alert]))

	if len(l.AppFields) > 0 {
		l.Fields = append(l.Fields, zap.Any(l.config.AppName, l.AppFields))
	}

	l.logger.Log(level.ToZap(l.Level), l.Message, l.Fields...)
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

func (l Base) WithContextSession(ctx context.Context) Log {
	// check if ctx have session_id or generate new one
	sessionID := ctx.Value("session_id")
	if sessionID == nil {
		sessionID = uuid.NewV4().String()
		ctx = context.WithValue(ctx, "session_id", sessionID)
	}

	return l.WithField("session_id", sessionID)

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
	l.Fields = append(l.Fields, zap.Any(key, value))

	return l
}

func (l Base) WithFields(fields map[string]any) Log {
	for k, v := range fields {
		l.Fields = append(l.Fields, zap.Any(k, v))
	}
	return l
}

func (l Base) WithStackTrace() Log {
	l.Fields = append(l.Fields, zap.Stack("stack_trace"))

	return l
}

func (l Base) WithHTTPReq(req HTTPRequestPayload) Log {
	if l.config.MaxBodySize > 0 && len(req.Body) > l.config.MaxBodySize {
		req.Body = req.Body[:l.config.MaxBodySize]
	}

	l.Fields = append(l.Fields, zap.Any("http_request", req))

	return l
}

func (l Base) WithHTTPResp(resp HTTPResponsePayload) Log {
	if l.config.MaxBodySize > 0 && len(resp.Body) > l.config.MaxBodySize {
		resp.Body = resp.Body[:l.config.MaxBodySize]
	}

	l.Fields = append(l.Fields, zap.Any("http_response", resp))

	return l
}

func (l Base) WithKafkaMessage(msg KafkaMessagePayload) Log {
	if l.config.MaxBodySize > 0 && len(msg.Payload) > l.config.MaxBodySize {
		msg.Payload = msg.Payload[:l.config.MaxBodySize]
	}

	l.Fields = append(l.Fields, zap.Any("kafka_message", msg))

	return l
}

func (l Base) WithKafkaResult(result KafkaResultPayload) Log {
	l.Fields = append(l.Fields, zap.Any("kafka_result", result))

	return l
}

func New(logger *zap.Logger, cfg config.Config, l level.Level) Base {
	return Base{
		logger:    logger,
		config:    cfg,
		Fields:    make([]zap.Field, 0, 1),
		AppFields: make(map[string]any),
		Level:     l,
	}
}
