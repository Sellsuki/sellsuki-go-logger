package log

import (
	"context"
	"github.com/Sellsuki/sellsuki-go-logger/level"
)

// Log is a flexible logging interface with methods for customizing log entries.
type Log interface {
	Write()                                        // Logs the current entry to the output.
	SetMessage(msg string) Log                     // Sets or overrides the log message.
	SetLevel(level level.Level) Log                // Sets or overrides the log level (e.g., info, warning, error).
	SetAlert(bool bool) Log                        // Sets or overrides the alert flag.
	WithContextSession(ctx context.Context) Log    // Generates a session ID and associates it with the log entry.
	WithAppData(key string, value any) Log         // Adds application-specific data.
	WithError(err error) Log                       // Adds error information.
	WithTracing(t Tracer) Log                      // Adds tracing information.
	WithField(key string, value any) Log           // Adds a custom field.
	WithFields(fields map[string]any) Log          // Adds multiple custom fields.
	WithStackTrace() Log                           // Captures and adds a stack trace.
	WithHTTPReq(req HTTPRequestPayload) Log        // Adds an HTTP request payload.
	WithHTTPResp(resp HTTPResponsePayload) Log     // Adds an HTTP response payload.
	WithKafkaMessage(msg KafkaMessagePayload) Log  // Adds a Kafka message payload.
	WithKafkaResult(result KafkaResultPayload) Log // Adds a Kafka result payload.
}

type Payload struct {
	Level   level.Level `json:"level"`
	Alert   bool        `json:"alert"`
	AppName string      `json:"app_name"`
	Version string      `json:"version"`
	LogType string      `json:"log_type"`
}
