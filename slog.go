package slog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type LogLevel int8

const (
	LevelDebug LogLevel = -1
	LevelInfo  LogLevel = 0
	LevelWarn  LogLevel = 1
	LevelError LogLevel = 2
	LevelPanic LogLevel = 4
	LevelFatal LogLevel = 5
)

var sukiLogger *SukiLogger

type Config struct {
	LogLevel LogLevel
	AppName  string
	Version  string
}

type SukiLogger struct {
	config      Config
	zapInstance *zap.Logger
}

type LogField struct {
	Key   string
	Value interface{}
}

type TraceInfo struct {
	TraceID string `json:"trace_id"`
	SpanID  string `json:"span_id"`
}

type HTTPRequestInfo struct {
	Method   string            `json:"method"`
	Path     string            `json:"path"`
	RemoteIP string            `json:"remote_ip"`
	Headers  map[string]string `json:"headers"`
	Params   map[string]string `json:"params"`
	Query    map[string]string `json:"query"`
	Body     string            `json:"body"`
}

type HTTPResponseInfo struct {
	Status   int64   `json:"status"`
	Duration float64 `json:"duration"`
	Body     string  `json:"body"`
	Error    Error   `json:"error"`
}

type Error struct {
	Name       string `json:"name"`
	Caller     string `json:"caller"`
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
	Duration float64 `json:"duration"`
	Error    Error   `json:"error"`
}

const (
	ActionCreate EventAction = "create"
	ActionUpdate EventAction = "update"
	ActionDelete EventAction = "delete"
)

const (
	ResultSuccess    EventResult = "success"
	ResultCompensate EventResult = "compensate"
)

type EventAction string
type EventResult string

type EventLog struct {
	Entity      string      `json:"entity"`
	Action      EventAction `json:"action"`
	Result      EventResult `json:"result"`
	ReferenceID string      `json:"reference_id"`
	Data        string      `json:"data"`
}

func Any(key string, value interface{}) LogField {
	return LogField{
		Key:   key,
		Value: value,
	}
}

func WithTracing(traceID string, spanID string) TraceInfo {
	return TraceInfo{
		TraceID: traceID,
		SpanID:  spanID,
	}
}

func WithEvent(entity string, action EventAction, result EventResult, data string, refID string) EventLog {
	return EventLog{
		Entity:      entity,
		Action:      action,
		Result:      result,
		ReferenceID: refID,
		Data:        data,
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

	return HTTPRequestInfo{
		Method:   method,
		Path:     path,
		RemoteIP: remoteIP,
		Headers:  headers,
		Params:   params,
		Query:    query,
		Body:     body,
	}
}

func WithError(name string, caller string, stacktrace string) Error {
	return Error{
		Name:       name,
		Caller:     caller,
		StackTrace: stacktrace,
	}
}

func WithHTTPResponse(
	status int64,
	duration float64,
	body string,
	err Error,
) HTTPResponseInfo {
	return HTTPResponseInfo{
		Status:   status,
		Duration: duration,
		Body:     body,
		Error:    err,
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
	return KafkaMessage{
		Topic:     topic,
		Partition: partition,
		Offset:    offset,
		Headers:   headers,
		Key:       key,
		Payload:   payload,
		Timestamp: timestamp,
	}
}

func WithKafkaResult(
	duration float64,
	error Error,
) KafkaResult {
	return KafkaResult{
		Duration: duration,
		Error:    error,
	}
}

func (s SukiLogger) RequestKafka(
	message string,
	kafkaMessage KafkaMessage,
	kafkaResult KafkaResult,
	tracing ...TraceInfo,
) {
	appName := zap.String("app_name", s.config.AppName)
	version := zap.String("version", s.config.Version)
	logType := zap.String("log_type", "handler.kafka")
	data := make(map[string]interface{})

	if len(tracing) > 0 {
		data["tracing"] = TraceInfo{
			TraceID: tracing[0].TraceID,
			SpanID:  tracing[0].SpanID,
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
		dataField,
	)
}

func (s SukiLogger) RequestHTTP(
	message string,
	request HTTPRequestInfo,
	response HTTPResponseInfo,
	tracing ...TraceInfo,
) {
	appName := zap.String("app_name", s.config.AppName)
	version := zap.String("version", s.config.Version)
	logType := zap.String("log_type", "handler.http")
	data := make(map[string]interface{})

	if len(tracing) > 0 {
		data["tracing"] = TraceInfo{
			TraceID: tracing[0].TraceID,
			SpanID:  tracing[0].SpanID,
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
		dataField,
	)

}

func (s SukiLogger) Event(message string, event EventLog, tracing ...TraceInfo) {
	appName := zap.String("app_name", s.config.AppName)
	version := zap.String("version", s.config.Version)
	logType := zap.String("log_type", "event")
	data := make(map[string]interface{})

	if len(tracing) > 0 {
		data["tracing"] = TraceInfo{
			TraceID: tracing[0].TraceID,
			SpanID:  tracing[0].SpanID,
		}
	}

	data["event"] = event
	dataField := zap.Any("data", data)

	s.zapInstance.Info(
		message,
		appName,
		version,
		logType,
		dataField,
	)

}

func (s SukiLogger) appLogBuilder(args ...interface{}) []zap.Field {
	appName := zap.String("app_name", s.config.AppName)
	version := zap.String("version", s.config.Version)
	logType := zap.String("log_type", "application")
	data := make(map[string]interface{})

	for i, _ := range args {
		if val, ok := args[i].(TraceInfo); ok {
			data["tracing"] = val
		} else if val, ok := args[i].(LogField); ok {
			data[val.Key] = val.Value
		}
	}

	dataField := zap.Any("data", data)

	return []zap.Field{
		appName,
		version,
		logType,
		dataField,
	}
}

func (s SukiLogger) Info(message string, args ...interface{}) {
	result := s.appLogBuilder(args...)

	s.zapInstance.Info(
		message,
		result...,
	)
}

func (s SukiLogger) Debug(message string, args ...interface{}) {
	result := s.appLogBuilder(args...)
	s.zapInstance.Debug(
		message,
		result...,
	)
}

func (s SukiLogger) Error(message string, args ...interface{}) {
	result := s.appLogBuilder(args...)
	s.zapInstance.Error(
		message,
		result...,
	)
}

func (s SukiLogger) Warn(message string, args ...interface{}) {
	result := s.appLogBuilder(args...)
	s.zapInstance.Warn(
		message,
		result...,
	)
}

func (s SukiLogger) Panic(message string, args ...interface{}) {
	result := s.appLogBuilder(args...)
	s.zapInstance.Panic(
		message,
		result...,
	)
}

func (s SukiLogger) Fatal(message string, args ...interface{}) {
	result := s.appLogBuilder(args...)
	s.zapInstance.Fatal(
		message,
		result...,
	)
}

func (s *SukiLogger) Configure(c Config) error {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	config.EncoderConfig.MessageKey = "message"
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.Level = zap.NewAtomicLevelAt(zapcore.Level(c.LogLevel))

	logger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		return err
	}
	defer logger.Sync()

	s.zapInstance = logger
	s.config = c
	return nil
}

func L() *SukiLogger {
	if sukiLogger == nil {
		logger, _ := zap.NewProduction()
		sukiLogger = &SukiLogger{zapInstance: logger}
	}
	return sukiLogger
}
