package slog

import (
	"encoding/json"
	"fmt"
	"github.com/Sellsuki/sellsuki-go-logger/config"
	"github.com/Sellsuki/sellsuki-go-logger/level"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
	"time"
)

var sukiLoggerOnce sync.Once

var sukiLogger *SukiLogger

type SukiLogger struct {
	config      config.Config
	zapInstance *zap.Logger
}

type LogField struct {
	Key   string
	Value interface{}
}

type AlertLevel int

var (
	LevelNone  AlertLevel = 0
	LevelAlert AlertLevel = 1
)

type LogOption struct {
	Alert AlertLevel
}

type TraceInfo struct {
	TraceID string `json:"trace_id"`
	SpanID  string `json:"span_id"`
}

type HTTPRequestInfo struct {
	Method    string            `json:"method"`
	Path      string            `json:"path"`
	RemoteIP  string            `json:"remote_ip"`
	Headers   map[string]string `json:"headers"`
	Params    map[string]string `json:"params"`
	Query     map[string]string `json:"query"`
	Body      string            `json:"body"`
	RequestID string            `json:"request_id"`
}

type HTTPResponseInfo struct {
	Status    int64     `json:"status"`
	Duration  float64   `json:"duration"`
	Body      string    `json:"body"`
	Error     ErrorInfo `json:"error"`
	RequestID string    `json:"request_id"`
}

type ErrorInfo struct {
	Name       string `json:"name"`
	StackTrace string `json:"stack_trace"`
}

type KafkaMessage struct {
	Topic     string            `json:"topic"`
	Partition int64             `json:"partition"`
	Offset    int64             `json:"offset"`
	Headers   map[string]string `json:"headers"`
	Key       string            `json:"key"`
	Payload   string            `json:"payload"`
	Timestamp time.Time         `json:"timestamp"`
}

type KafkaResult struct {
	Duration float64   `json:"duration"`
	Error    ErrorInfo `json:"error"`
}

const (
	ActionCreate EventAction = "create"
	ActionUpdate EventAction = "update"
	ActionDelete EventAction = "delete"

	AuditActionCreate AuditAction = "create"
	AuditActionUpdate AuditAction = "update"
	AuditActionDelete AuditAction = "delete"
	AuditActionAccess AuditAction = "access"
)

const (
	ResultSuccess    EventResult = "success"
	ResultCompensate EventResult = "compensate"
)

type EventAction string
type EventResult string
type AuditAction string

type EventLog struct {
	Entity      string      `json:"entity"`
	Action      EventAction `json:"action"`
	Result      EventResult `json:"result"`
	ReferenceID string      `json:"reference_id"`
	Data        string      `json:"data"`
}

type AuditLog struct {
	ActorType   string      `json:"actor_type"`
	ActorID     string      `json:"actor_id"`
	Action      AuditAction `json:"action"`
	Entity      string      `json:"entity"`
	EntityRefs  []string    `json:"entity_refs"`
	EntityOwner string      `json:"entity_owner"`
}

func Any(key string, value interface{}) LogField {
	return LogField{
		Key:   key,
		Value: value,
	}
}

func Error(err error) LogField {
	if err != nil {
		return Any("error", err.Error())
	}
	return Any("error", "")
}

func WithOption(opts LogOption) LogOption {
	return opts
}

func WithEvent(entity string, action EventAction, result EventResult, data interface{}, refID string) EventLog {
	payload := ""

	if data == nil {
		payload = ""
	} else if str, ok := data.(string); ok {
		payload = str
	} else {
		r, err := json.Marshal(data)
		if err == nil {
			payload = string(r)
		}
	}

	return EventLog{
		Entity:      entity,
		Action:      action,
		Result:      result,
		ReferenceID: refID,
		Data:        payload,
	}
}

func WithAudit(actorType, actorId string, action AuditAction, entity, entityOwner string, entityRefs ...string) AuditLog {

	return AuditLog{
		ActorType:   actorType,
		ActorID:     actorId,
		Action:      action,
		Entity:      entity,
		EntityRefs:  entityRefs,
		EntityOwner: entityOwner,
	}
}

func WithHTTPRequest(
	method string,
	path string,
	remoteIP string,
	headers map[string]string,
	params map[string]string,
	query map[string]string,
	body string,
) HTTPRequestInfo {
	h := headers
	p := params
	q := query
	if h == nil {
		h = map[string]string{}
	}

	if p == nil {
		p = map[string]string{}
	}

	if q == nil {
		q = map[string]string{}
	}

	return HTTPRequestInfo{
		Method:   method,
		Path:     path,
		RemoteIP: remoteIP,
		Headers:  h,
		Params:   p,
		Query:    q,
		Body:     body,
	}
}

func WithError(name string, stacktrace ...string) ErrorInfo {

	var trace string

	if len(stacktrace) == 1 {
		trace = stacktrace[0]
	} else if len(stacktrace) > 1 {
		trace = stacktrace[1]
	}

	return ErrorInfo{
		Name:       name,
		StackTrace: trace,
	}
}

func WithHTTPResponse(
	status int64,
	duration float64,
	body string,
	error ...ErrorInfo,
) HTTPResponseInfo {
	var e ErrorInfo
	if len(error) > 0 {
		e = error[0]
	}
	return HTTPResponseInfo{
		Status:   status,
		Duration: duration,
		Body:     body,
		Error:    e,
	}
}

func WithKafkaMessage(
	topic string,
	partition int64,
	offset int64,
	headers map[string]string,
	key string,
	payload string,
	timestamp time.Time,
) KafkaMessage {

	h := headers
	if h == nil {
		h = map[string]string{}
	}

	return KafkaMessage{
		Topic:     topic,
		Partition: partition,
		Offset:    offset,
		Headers:   h,
		Key:       key,
		Payload:   payload,
		Timestamp: timestamp,
	}
}

func WithKafkaResult(
	duration float64,
	error ...ErrorInfo,
) KafkaResult {
	var e ErrorInfo
	if len(error) > 0 {
		e = error[0]
	}
	return KafkaResult{
		Duration: duration,
		Error:    e,
	}
}

func (s SukiLogger) RequestKafka(
	message string,
	kafkaMessage KafkaMessage,
	kafkaResult KafkaResult,
	args ...interface{},
) {
	appName := zap.String("app_name", s.config.AppName)
	version := zap.String("version", s.config.Version)
	logType := zap.String("log_type", "handler.kafka")
	data := make(map[string]interface{})
	alertLevel := LevelNone

	for i, _ := range args {
		if tracing, ok := args[i].(TraceInfo); ok {
			data["tracing"] = TraceInfo{
				TraceID: tracing.TraceID,
				SpanID:  tracing.SpanID,
			}
		} else if opts, ok := args[i].(LogOption); ok {
			alertLevel = opts.Alert
		}
	}

	data["kafka_message"] = kafkaMessage
	data["kafka_result"] = kafkaResult
	dataField := zap.Any("data", data)

	s.zapInstance.Info(
		message,
		appName,
		version,
		logType,
		zap.Int("alert", int(alertLevel)),
		dataField,
	)
}

func (s SukiLogger) RequestHTTP(
	message string,
	request HTTPRequestInfo,
	response HTTPResponseInfo,
	args ...interface{},
) {
	appName := zap.String("app_name", s.config.AppName)
	version := zap.String("version", s.config.Version)
	logType := zap.String("log_type", "handler.http")
	data := make(map[string]interface{})
	alertLevel := LevelNone

	for i, _ := range args {
		if tracing, ok := args[i].(TraceInfo); ok {
			data["tracing"] = TraceInfo{
				TraceID: tracing.TraceID,
				SpanID:  tracing.SpanID,
			}
		} else if opts, ok := args[i].(LogOption); ok {
			alertLevel = opts.Alert
		}
	}

	if s.config.MaxBodySize > 0 {
		if len(request.Body) > s.config.MaxBodySize {
			request.Body = "body is too large"
		}

		if len(response.Body) > s.config.MaxBodySize {
			response.Body = "body is too large"
		}
	}

	data["http_request"] = request
	data["http_response"] = response
	dataField := zap.Any("data", data)

	s.zapInstance.Info(
		message,
		appName,
		version,
		logType,
		zap.Int("alert", int(alertLevel)),
		dataField,
	)

}

func Init(c ...config.Config) *SukiLogger {
	sukiLoggerOnce.Do(func() {
		cfg := config.Config{
			LogLevel:    level.Info,
			AppName:     "unknown",
			Version:     "v0.0.0",
			MaxBodySize: 1048576,
		}
		if len(c) > 0 {
			cfg = c[0]
		}

		zCfg := zap.NewProductionConfig()
		zCfg.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		zCfg.EncoderConfig.MessageKey = "message"
		zCfg.EncoderConfig.TimeKey = "timestamp"
		zCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		zCfg.Level = zap.NewAtomicLevelAt(level.LevelMapping(cfg.LogLevel))

		logger, err := zCfg.Build(zap.AddCallerSkip(1))
		if err != nil {
			panic(fmt.Errorf("failed to init logger: %w", err))
		}
		defer logger.Sync()

		sukiLogger = &SukiLogger{zapInstance: logger, config: cfg}
	})

	return sukiLogger
}
