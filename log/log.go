package log

import (
	"github.com/Sellsuki/sellsuki-go-logger/level"
)

type Log interface {
	Write()

	SetMessage(msg string) Log
	SetLevel(level level.Level) Log
	SetAlert(bool bool) Log

	WithAppData(key string, value any) Log
	WithError(err error) Log
	WithTracing(t Tracer) Log
	WithField(key string, value any) Log
	WithFields(fields map[string]any) Log
	WithStackTrace() Log
	WithHttpReq(req HTTPRequestPayload) Log
	WithHttpResp(resp HTTPResponsePayload) Log
}

type Payload struct {
	Level   level.Level `json:"level"`
	Alert   bool        `json:"alert"`
	AppName string      `json:"app_name"`
	Version string      `json:"version"`
	LogType string      `json:"log_type"`
}
