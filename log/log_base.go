package log

import (
	"github.com/Sellsuki/sellsuki-go-logger/level"
	"go.uber.org/zap"
)

type Base struct {
	logger *zap.Logger

	Level     level.Level
	Alert     bool
	Message   string
	Fields    []zap.Field
	AppFields []zap.Field
}

func (l Base) Write() {
	l.Fields = append(l.Fields, zap.Int("alert", boolToInt[l.Alert]))

	l.logger.Log(level.LevelMapping(l.Level), l.Message, l.Fields...)
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
	l.AppFields = append(l.AppFields, zap.Any(key, value))
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

func New(logger *zap.Logger, l level.Level) Base {
	return Base{
		logger:    logger,
		Fields:    make([]zap.Field, 0, 1),
		AppFields: make([]zap.Field, 0),
		Level:     l,
	}
}
