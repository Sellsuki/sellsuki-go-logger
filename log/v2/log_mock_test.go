package v2

import (
	"github.com/Sellsuki/sellsuki-go-logger/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

// Define a TestTracer struct that implements the Tracer interface for testing.
type TestTracer struct {
	TraceIDVal string
	SpanIDVal  string
}

func (t *TestTracer) TraceID() log.Stringer {
	return TestStringer(t.TraceIDVal)
}

func (t *TestTracer) SpanID() log.Stringer {
	return TestStringer(t.SpanIDVal)
}

// Define a TestStringer struct that implements the Stringer interface for testing.
type TestStringer string

func (s TestStringer) String() string {
	return string(s)
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
