package log

import (
	"fmt"
	"github.com/Sellsuki/sellsuki-go-logger/config"
	"github.com/Sellsuki/sellsuki-go-logger/level"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"reflect"
	"testing"
	"time"
)

func TestBase_SetAlert(t *testing.T) {
	type fields struct {
		config  config.Config
		Level   level.Level
		Alert   bool
		Message string
	}
	type args struct {
		bool bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Log
	}{{
		name:   "Set alert to true",
		fields: fields{},
		args:   args{bool: true},
		want: Base{
			config:  config.Config{},
			Level:   level.Level(0),
			Alert:   true,
			Message: "",
		},
	},
		{
			name: "Set alert to false",
			fields: fields{
				Alert: true, // Initialize with Alert set to true
			},
			args: args{bool: false},
			want: Base{
				config:  config.Config{},
				Level:   level.Level(0),
				Alert:   false,
				Message: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Base{
				config:  tt.fields.config,
				Level:   tt.fields.Level,
				Alert:   tt.fields.Alert,
				Message: tt.fields.Message,
			}
			if got := l.SetAlert(tt.args.bool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAlert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase_SetLevel(t *testing.T) {
	type fields struct {
		config  config.Config
		Level   level.Level
		Alert   bool
		Message string
	}
	type args struct {
		level level.Level
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Log
	}{
		{
			name:   "Set level to Info",
			fields: fields{},
			args:   args{level: level.Info},
			want: Base{
				config:  config.Config{},
				Level:   level.Info,
				Alert:   false,
				Message: "",
			},
		},
		{
			name: "Set level to Error",
			fields: fields{
				Level: level.Info, // Initialize with Level set to Info
			},
			args: args{level: level.Error},
			want: Base{
				config:  config.Config{},
				Level:   level.Error,
				Alert:   false,
				Message: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Base{
				config:  tt.fields.config,
				Level:   tt.fields.Level,
				Alert:   tt.fields.Alert,
				Message: tt.fields.Message,
			}
			if got := l.SetLevel(tt.args.level); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase_SetMessage(t *testing.T) {
	type fields struct {
		config  config.Config
		Level   level.Level
		Alert   bool
		Message string
	}
	type args struct {
		msg string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Log
	}{
		{
			name:   "Set an empty message",
			fields: fields{},
			args:   args{msg: ""},
			want: Base{
				config:  config.Config{},
				Level:   level.Level(0),
				Alert:   false,
				Message: "",
			},
		},
		// Future feature ?
		//{
		//	name:   "Set a long message",
		//	fields: fields{},
		//	args:   args{msg: "This is a very long message that exceeds the character limit for messages and needs to be truncated."},
		//	want: Base{
		//		config: config.Config{
		//		},
		//		Level:   level.Level(0),
		//		Alert:   false,
		//		Message: "This is a very long message that exceeds the character limit for messages and needs to be trunca",
		//	},
		//},
		{
			name:   "Set a message with special characters",
			fields: fields{},
			args:   args{msg: "Hello, !@#$%^&*()_+{}:\"<>? World"},
			want: Base{
				config:  config.Config{},
				Level:   level.Level(0),
				Alert:   false,
				Message: "Hello, !@#$%^&*()_+{}:\"<>? World",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Base{
				config:  tt.fields.config,
				Level:   tt.fields.Level,
				Alert:   tt.fields.Alert,
				Message: tt.fields.Message,
			}
			if got := l.SetMessage(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase_WithAppData(t *testing.T) {
	type fields struct {
		config    config.Config
		Level     level.Level
		Alert     bool
		Message   string
		Fields    []zap.Field
		AppFields map[string]interface{}
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Log
	}{
		{
			name: "Set app data with key and value",
			fields: fields{
				config: config.Config{
					AppName: "app_name",
				},
				AppFields: map[string]interface{}{},
			},
			args: args{
				key:   "key",
				value: "value",
			},
			want: Base{
				config: config.Config{
					AppName: "app_name",
				},
				AppFields: map[string]interface{}{
					"key": "value",
				},
			},
		},
		{
			name: "Update existing app data",
			fields: fields{
				config: config.Config{
					AppName: "app_name",
				},
				AppFields: map[string]interface{}{
					"existing_key": "existing_value",
				},
			},
			args: args{
				key:   "existing_key",
				value: "new_value",
			},
			want: Base{
				config: config.Config{
					AppName: "app_name",
				},
				AppFields: map[string]interface{}{
					"existing_key": "new_value",
				},
			},
		},
		{
			name: "Set app data with a numeric value",
			fields: fields{
				config: config.Config{
					AppName: "app_name",
				},
				AppFields: map[string]interface{}{},
			},
			args: args{
				key:   "numeric_key",
				value: 12345,
			},
			want: Base{
				config: config.Config{
					AppName: "app_name",
				},
				AppFields: map[string]interface{}{
					"numeric_key": 12345,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Base{
				config:    tt.fields.config,
				Level:     tt.fields.Level,
				Alert:     tt.fields.Alert,
				Message:   tt.fields.Message,
				Fields:    tt.fields.Fields,
				AppFields: tt.fields.AppFields,
			}
			if got := l.WithAppData(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithAppData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase_WithError(t *testing.T) {
	logger, _ := zap.NewDevelopment()

	type fields struct {
		logger    *zap.Logger
		config    config.Config
		Level     level.Level
		Alert     bool
		Message   string
		Fields    []zap.Field
		AppFields map[string]interface{}
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Base
	}{
		{
			name: "Add an error to fields",
			fields: fields{
				logger: logger,
			},
			args: args{
				err: fmt.Errorf("Sample error message"),
			},
			want: Base{
				logger: logger,
				Fields: []zap.Field{zap.Error(fmt.Errorf("Sample error message"))},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Base{
				logger:    tt.fields.logger,
				config:    tt.fields.config,
				Level:     tt.fields.Level,
				Alert:     tt.fields.Alert,
				Message:   tt.fields.Message,
				Fields:    tt.fields.Fields,
				AppFields: tt.fields.AppFields,
			}
			if got := l.WithError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase_WithField(t *testing.T) {
	type fields struct {
		Fields []zap.Field
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Base
	}{
		{
			name:   "Add a single string field",
			fields: fields{},
			args: args{
				key:   "field_key",
				value: "field_value",
			},
			want: Base{
				Fields: []zap.Field{zap.String("field_key", "field_value")},
			},
		},
		{
			name:   "Add a single integer field",
			fields: fields{},
			args: args{
				key:   "key1",
				value: 123,
			},
			want: Base{
				Fields: []zap.Field{zap.Int("key1", 123)},
			},
		},
		{
			name:   "Add a single float field",
			fields: fields{},
			args: args{
				key:   "key1",
				value: 3.14,
			},
			want: Base{
				Fields: []zap.Field{zap.Float64("key1", 3.14)},
			},
		},
		{
			name:   "Add a single boolean field",
			fields: fields{},
			args: args{
				key:   "key1",
				value: true,
			},
			want: Base{
				Fields: []zap.Field{zap.Bool("key1", true)},
			},
		},
		{
			name:   "Add a single nil field",
			fields: fields{},
			args: args{
				key:   "key1",
				value: nil,
			},
			want: Base{
				Fields: []zap.Field{zap.Any("key1", nil)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Base{
				Fields: tt.fields.Fields,
			}
			if got := l.WithField(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase_WithFields(t *testing.T) {
	type fields struct {
		Fields []zap.Field
	}
	type args struct {
		fields map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Base
	}{
		{
			name:   "Add multiple fields",
			fields: fields{},
			args: args{
				fields: map[string]interface{}{
					"key1": "value1",
					"key2": 123,
				},
			},
			want: Base{
				Fields: []zap.Field{
					zap.String("key1", "value1"),
					zap.Int("key2", 123),
				},
			},
		},
		{
			name:   "Add no fields",
			fields: fields{},
			args: args{
				fields: map[string]interface{}{},
			},
			want: Base{
				Fields: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Base{
				Fields: tt.fields.Fields,
			}
			assert.Equal(t, tt.want, l.WithFields(tt.args.fields))
		})
	}
}

func TestBase_WithHTTPReq(t *testing.T) {
	logger, _ := zap.NewDevelopment()

	type fields struct {
		logger    *zap.Logger
		config    config.Config
		Level     level.Level
		Alert     bool
		Message   string
		Fields    []zap.Field
		AppFields map[string]interface{}
	}
	type args struct {
		req HTTPRequestPayload
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Base
	}{
		{
			name: "Add HTTP request payload with JSON body to fields",
			fields: fields{
				logger: logger,
			},
			args: args{
				req: HTTPRequestPayload{
					Method: "POST",
					Path:   "/api",
					Headers: map[string]string{
						"Content-Type": "application/json",
					},
					Body:      []byte(`{"key": "value", "nested": {"subkey": 123}}`),
					RequestID: "789012",
				},
			},
			want: Base{
				logger: logger,
				Fields: []zap.Field{
					zap.Any("http_request", HTTPRequestPayload{
						Method: "POST",
						Path:   "/api",
						Headers: map[string]string{
							"Content-Type": "application/json",
						},
						Body:      []byte(`{"key": "value", "nested": {"subkey": 123}}`),
						RequestID: "789012",
					}),
				},
			},
		},
		{
			name: "Trim body when it exceeds MaxBodySize",
			fields: fields{
				logger: logger,
				config: config.Config{
					MaxBodySize: 10,
				},
			},
			args: args{
				req: HTTPRequestPayload{
					Method: "POST",
					Path:   "/api",
					Headers: map[string]string{
						"Content-Type": "application/json",
					},
					// Body exceeds the configured MaxBodySize
					Body:      []byte(`{"large_body": "This is a large body that exceeds the maximum allowed size"}`),
					RequestID: "789012",
				},
			},
			want: Base{
				logger: logger,
				config: config.Config{
					MaxBodySize: 10,
				},
				Fields: []zap.Field{
					zap.Any("http_request", HTTPRequestPayload{
						Method: "POST",
						Path:   "/api",
						Headers: map[string]string{
							"Content-Type": "application/json",
						},
						// Body is trimmed to the maximum allowed size
						Body:      []byte(`{"large_bo`),
						RequestID: "789012",
					}),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Base{
				logger:    tt.fields.logger,
				config:    tt.fields.config,
				Level:     tt.fields.Level,
				Alert:     tt.fields.Alert,
				Message:   tt.fields.Message,
				Fields:    tt.fields.Fields,
				AppFields: tt.fields.AppFields,
			}

			assert.Equal(t, tt.want, l.WithHTTPReq(tt.args.req))
		})
	}
}

func TestBase_WithHTTPResp(t *testing.T) {
	logger, _ := zap.NewDevelopment()

	type fields struct {
		logger    *zap.Logger
		config    config.Config
		Level     level.Level
		Alert     bool
		Message   string
		Fields    []zap.Field
		AppFields map[string]any
	}
	type args struct {
		resp HTTPResponsePayload
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Base
	}{
		{
			name: "Add HTTP response payload with status, duration, and body",
			fields: fields{
				logger: logger,
			},
			args: args{
				resp: HTTPResponsePayload{
					Status:    200,
					Duration:  time.Millisecond * 100,
					Body:      []byte("Response body content"),
					RequestID: "123456",
				},
			},
			want: Base{
				logger: logger,
				Fields: []zap.Field{
					zap.Any("http_response", HTTPResponsePayload{
						Status:    200,
						Duration:  time.Millisecond * 100,
						Body:      []byte("Response body content"),
						RequestID: "123456",
					}),
				},
			},
		},
		{
			name: "Add HTTP response payload with an error message",
			fields: fields{
				logger: logger,
			},
			args: args{
				resp: HTTPResponsePayload{
					Status:    500,
					Duration:  time.Millisecond * 500,
					Error:     "Internal Server Error",
					RequestID: "789012",
				},
			},
			want: Base{
				logger: logger,
				Fields: []zap.Field{
					zap.Any("http_response", HTTPResponsePayload{
						Status:    500,
						Duration:  time.Millisecond * 500,
						Error:     "Internal Server Error",
						RequestID: "789012",
					}),
				},
			},
		},
		{
			name: "Trim body when it exceeds MaxBodySize",
			fields: fields{
				logger: logger,
				config: config.Config{
					MaxBodySize: 10,
				},
			},
			args: args{
				resp: HTTPResponsePayload{
					Status:   200,
					Duration: time.Millisecond * 100,
					// Body exceeds the configured MaxBodySize
					Body:      []byte("This is a large body that exceeds the maximum allowed size"),
					RequestID: "123456",
				},
			},
			want: Base{
				logger: logger,
				config: config.Config{
					MaxBodySize: 10,
				},
				Fields: []zap.Field{
					zap.Any("http_response", HTTPResponsePayload{
						Status:   200,
						Duration: time.Millisecond * 100,
						// Body is trimmed to the maximum allowed size
						Body:      []byte("This is a "),
						RequestID: "123456",
					}),
				},
			},
		},
		// Add more test cases for different response payloads
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Base{
				logger:    tt.fields.logger,
				config:    tt.fields.config,
				Level:     tt.fields.Level,
				Alert:     tt.fields.Alert,
				Message:   tt.fields.Message,
				Fields:    tt.fields.Fields,
				AppFields: tt.fields.AppFields,
			}

			assert.Equal(t, tt.want, l.WithHTTPResp(tt.args.resp))
		})
	}
}

func TestBase_WithKafkaMessage(t *testing.T) {
	logger, _ := zap.NewDevelopment()

	now := time.Now()
	type fields struct {
		logger    *zap.Logger
		config    config.Config
		Level     level.Level
		Alert     bool
		Message   string
		Fields    []zap.Field
		AppFields map[string]any
	}
	type args struct {
		msg KafkaMessagePayload
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Base
	}{
		{
			name: "Add Kafka message payload with details",
			fields: fields{
				logger: logger,
			},
			args: args{
				msg: KafkaMessagePayload{
					Topic:     "important-events",
					Partition: 1,
					Offset:    12345,
					Headers:   map[string]string{"key1": "value1", "key2": "value2"},
					Key:       "message-key",
					Payload:   []byte("This is a Kafka message payload."),
					Timestamp: now,
				},
			},
			want: Base{
				logger: logger,
				Fields: []zap.Field{
					zap.Any("kafka_message", KafkaMessagePayload{
						Topic:     "important-events",
						Partition: 1,
						Offset:    12345,
						Headers:   map[string]string{"key1": "value1", "key2": "value2"},
						Key:       "message-key",
						Payload:   []byte("This is a Kafka message payload."),
						Timestamp: now,
					}),
				},
			},
		},
		{
			name: "Add Kafka message payload with large payload",
			fields: fields{
				config: config.Config{
					MaxBodySize: 20,
				},
				logger: logger,
			},
			args: args{
				msg: KafkaMessagePayload{
					Topic:     "important-events",
					Partition: 1,
					Offset:    12345,
					Headers:   map[string]string{"key1": "value1", "key2": "value2"},
					Key:       "message-key",
					// Create a large payload exceeding the maximum allowed size
					Payload:   []byte("This is a very large Kafka message payload exceeding the maximum allowed size."),
					Timestamp: now,
				},
			},
			want: Base{
				config: config.Config{
					MaxBodySize: 20,
				},
				logger: logger,
				Fields: []zap.Field{
					zap.Any("kafka_message", KafkaMessagePayload{
						Topic:     "important-events",
						Partition: 1,
						Offset:    12345,
						Headers:   map[string]string{"key1": "value1", "key2": "value2"},
						Key:       "message-key",
						// Ensure that the payload is trimmed to the maximum allowed size
						Payload:   []byte("This is a very large"),
						Timestamp: now,
					}),
				},
			},
			// Add more test cases for different Kafka message payloads
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Base{
				logger:    tt.fields.logger,
				config:    tt.fields.config,
				Level:     tt.fields.Level,
				Alert:     tt.fields.Alert,
				Message:   tt.fields.Message,
				Fields:    tt.fields.Fields,
				AppFields: tt.fields.AppFields,
			}

			assert.Equal(t, tt.want, l.WithKafkaMessage(tt.args.msg))
		})
	}
}

func TestBase_WithKafkaResult(t *testing.T) {
	logger, _ := zap.NewDevelopment()

	duration := time.Millisecond * 500 // Sample duration
	errMessage := "Kafka error message"

	type fields struct {
		logger    *zap.Logger
		config    config.Config
		Level     level.Level
		Alert     bool
		Message   string
		Fields    []zap.Field
		AppFields map[string]any
	}

	type args struct {
		result KafkaResultPayload
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   Base
	}{
		{
			name: "Add Kafka result payload with success",
			fields: fields{
				logger: logger,
			},
			args: args{
				result: KafkaResultPayload{
					Duration: duration,
				},
			},
			want: Base{
				logger: logger,
				Fields: []zap.Field{
					zap.Any("kafka_result", KafkaResultPayload{
						Duration: duration,
					}),
				},
			},
		},
		{
			name: "Add Kafka result payload with error message",
			fields: fields{
				logger: logger,
			},
			args: args{
				result: KafkaResultPayload{
					Duration: duration,
					Error:    errMessage,
				},
			},
			want: Base{
				logger: logger,
				Fields: []zap.Field{
					zap.Any("kafka_result", KafkaResultPayload{
						Duration: duration,
						Error:    errMessage,
					}),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Base{
				logger:    tt.fields.logger,
				config:    tt.fields.config,
				Level:     tt.fields.Level,
				Alert:     tt.fields.Alert,
				Message:   tt.fields.Message,
				Fields:    tt.fields.Fields,
				AppFields: tt.fields.AppFields,
			}

			assert.Equal(t, tt.want, l.WithKafkaResult(tt.args.result))
		})
	}
}

func TestBase_WithStackTrace(t *testing.T) {
	// No idea how to test this
}

func TestBase_WithTracing(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	type fields struct {
		logger    *zap.Logger
		config    config.Config
		Level     level.Level
		Alert     bool
		Message   string
		Fields    []zap.Field
		AppFields map[string]any
	}

	type args struct {
		t Tracer
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   Base
	}{
		{
			name: "Add tracing information",
			fields: fields{
				logger: logger, // Assuming you have a logger instance available.
			},
			args: args{
				t: &TestTracer{
					TraceIDVal: "12345",
					SpanIDVal:  "67890",
				},
			},
			want: Base{
				logger: logger,
				Fields: []zap.Field{zap.Any("tracing", map[string]string{
					"trace_id": "12345",
					"span_id":  "67890",
				}),
				},
			},
		},
		// Add more test cases for different Tracer implementations
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Base{
				logger:    tt.fields.logger,
				config:    tt.fields.config,
				Level:     tt.fields.Level,
				Alert:     tt.fields.Alert,
				Message:   tt.fields.Message,
				Fields:    tt.fields.Fields,
				AppFields: tt.fields.AppFields,
			}

			assert.Equal(t, tt.want, l.WithTracing(tt.args.t))
		})
	}
}

func TestBase_Write(t *testing.T) {
	// Create an instance of your Base struct with a mock logger
	mockLogger := &MockLogger{}
	base := Base{
		logger:  mockLogger,
		config:  config.Config{}, // Replace with your config initialization
		Level:   level.Info,      // Set the desired log level
		Alert:   true,            // Set the alert flag to true or false as needed
		Message: "Test message",  // Set your desired log message
		// Set other fields and AppFields as needed
	}

	// Call the Write method
	base.Write()

	// Assert that the mock logger was called with the expected data
	if !mockLogger.logged {
		t.Errorf("Expected the logger to be called, but it wasn't.")
	}
	if mockLogger.level != zap.InfoLevel {
		t.Errorf("Expected log level Info, got %v", mockLogger.level)
	}
	if mockLogger.msg != "Test message" {
		t.Errorf("Expected message 'Test message', got %v", mockLogger.msg)
	}
	// Add more assertions for fields if needed
}

func TestBase_New(t *testing.T) {
	type args struct {
		logger *zap.Logger
		cfg    config.Config
		l      level.Level
	}
	tests := []struct {
		name string
		args args
		want Base
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, New(tt.args.logger, tt.args.cfg, tt.args.l), "New(%v, %v, %v)", tt.args.logger, tt.args.cfg, tt.args.l)
		})
	}
}
