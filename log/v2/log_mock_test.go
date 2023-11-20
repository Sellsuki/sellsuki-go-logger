package v2

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

// Define a TestSpanContext struct that implements the Tracer interface for testing.
type TestSpanContext struct {
	TraceIDVal [16]byte
	SpanIDVal  [8]byte
}

func (t TestSpanContext) TraceID() [16]byte {
	return t.TraceIDVal
}

func (t TestSpanContext) SpanID() [8]byte {
	return t.SpanIDVal
}

type MockLogger struct {
	logged bool
	level  zapcore.Level
	msg    string
	fields []zap.Field
}

func (m *MockLogger) Log(level zapcore.Level, msg string, fields ...zap.Field) {
	m.logged = true
	m.level = level
	m.msg = msg
	m.fields = fields
}

func FixedTimeEncoder(_ time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("fixed")
}
