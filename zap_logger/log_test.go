package zap_logger

import (
	"bytes"
	"fmt"
	"github.com/Sellsuki/sellsuki-go-logger/v2/config"
	"github.com/Sellsuki/sellsuki-go-logger/v2/level"
	"github.com/Sellsuki/sellsuki-go-logger/v2/log"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"reflect"
	"testing"
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
		want   log.Log
	}{{
		name:   "Set alert to true",
		fields: fields{},
		args:   args{bool: true},
		want: &Logger{
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
			want: &Logger{
				config:  config.Config{},
				Level:   level.Level(0),
				Alert:   false,
				Message: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Logger{
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
		want   log.Log
	}{
		{
			name:   "Set level to Info",
			fields: fields{},
			args:   args{level: level.Info},
			want: &Logger{
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
			want: &Logger{
				config:  config.Config{},
				Level:   level.Error,
				Alert:   false,
				Message: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Logger{
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
		want   log.Log
	}{
		{
			name:   "Set an empty message",
			fields: fields{},
			args:   args{msg: ""},
			want: &Logger{
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
		//	want: &Logger{
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
			want: &Logger{
				config:  config.Config{},
				Level:   level.Level(0),
				Alert:   false,
				Message: "Hello, !@#$%^&*()_+{}:\"<>? World",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Logger{
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
		Fields    map[string]any
		AppFields map[string]any
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   log.Log
	}{
		{
			name: "Set app data with key and value",
			fields: fields{
				config: config.Config{
					AppName: "app_name",
				},
				AppFields: map[string]any{},
			},
			args: args{
				key:   "key",
				value: "value",
			},
			want: &Logger{
				config: config.Config{
					AppName: "app_name",
				},
				AppFields: map[string]any{
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
				AppFields: map[string]any{
					"existing_key": "existing_value",
				},
			},
			args: args{
				key:   "existing_key",
				value: "new_value",
			},
			want: &Logger{
				config: config.Config{
					AppName: "app_name",
				},
				AppFields: map[string]any{
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
				AppFields: map[string]any{},
			},
			args: args{
				key:   "numeric_key",
				value: 12345,
			},
			want: &Logger{
				config: config.Config{
					AppName: "app_name",
				},
				AppFields: map[string]any{
					"numeric_key": 12345,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Logger{
				config:    tt.fields.config,
				Level:     tt.fields.Level,
				Alert:     tt.fields.Alert,
				Message:   tt.fields.Message,
				Data:      tt.fields.Fields,
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
		Fields    map[string]any
		AppFields map[string]any
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Logger
	}{
		{
			name: "Add an error to fields",
			fields: fields{
				logger: logger,
				Fields: map[string]any{},
			},
			args: args{
				err: fmt.Errorf("Sample error message"),
			},
			want: &Logger{
				logger: logger,
				Data:   map[string]any{"error": "Sample error message"},
			},
		},
		{
			name: "Add a nil error to fields",
			fields: fields{
				logger: logger,
				Fields: map[string]any{},
			},
			args: args{
				err: nil,
			},
			want: &Logger{
				logger: logger,
				Data:   map[string]any{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Logger{
				logger:    tt.fields.logger,
				config:    tt.fields.config,
				Level:     tt.fields.Level,
				Alert:     tt.fields.Alert,
				Message:   tt.fields.Message,
				Data:      tt.fields.Fields,
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
		Fields map[string]any
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Logger
	}{
		{
			name: "Add a single string field",
			fields: fields{
				Fields: map[string]any{},
			},
			args: args{
				key:   "field_key",
				value: "field_value",
			},
			want: &Logger{
				Data: map[string]any{"field_key": "field_value"},
			},
		},
		{
			name: "Add a single integer field",
			fields: fields{
				Fields: map[string]any{}},
			args: args{
				key:   "key1",
				value: 123,
			},
			want: &Logger{
				Data: map[string]any{"key1": 123},
			},
		},
		{
			name: "Add a single float field",
			fields: fields{
				Fields: map[string]any{}},
			args: args{
				key:   "key1",
				value: 3.14,
			},
			want: &Logger{
				Data: map[string]any{"key1": 3.14},
			},
		},
		{
			name: "Add a single boolean field",
			fields: fields{
				Fields: map[string]any{}},
			args: args{
				key:   "key1",
				value: true,
			},
			want: &Logger{
				Data: map[string]any{"key1": true},
			},
		},
		{
			name: "Add a single nil field",
			fields: fields{
				Fields: map[string]any{}},
			args: args{
				key:   "key1",
				value: nil,
			},
			want: &Logger{
				Data: map[string]any{"key1": nil},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Logger{
				Data: tt.fields.Fields,
			}
			if got := l.WithField(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase_WithFields(t *testing.T) {
	type fields struct {
		Fields map[string]any
	}
	type args struct {
		fields map[string]any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Logger
	}{
		{
			name: "Add multiple fields",
			fields: fields{
				Fields: map[string]any{}},
			args: args{
				fields: map[string]any{
					"key1": "value1",
					"key2": 123,
				},
			},
			want: &Logger{
				Data: map[string]any{
					"key1": "value1",
					"key2": 123,
				},
			},
		},
		{
			name: "Add no fields",
			fields: fields{
				Fields: map[string]any{},
			},
			args: args{
				fields: map[string]any{},
			},
			want: &Logger{
				Data: map[string]any{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Logger{
				Data: tt.fields.Fields,
			}
			assert.Equal(t, tt.want, l.WithFields(tt.args.fields))
		})
	}
}

func TestBase_WithStackTrace(t *testing.T) {
	// Create a zap.Logger for testing purposes
	logger, _ := zap.NewDevelopment()

	// Create a sample config
	config := config.Config{
		// Initialize your configuration fields here.
	}

	// Initialize a Logger object with the required parameters
	base := Logger{
		logger:  logger,
		config:  config,
		Level:   level.Info, // Set the level as needed.
		Alert:   true,       // Set the Alert as needed.
		Message: "Test message",
		Data:    map[string]any{},
		AppFields: map[string]any{
			"key1": "value1",
			"key2": "value2",
		},
	}

	// Call the WithStackTrace method
	baseWithStackTrace := base.WithStackTrace()
	//expectedStackTrace := ``

	stack, ok := baseWithStackTrace.(*Logger).Data["stack_trace"]

	if !ok {
		t.Errorf("WithStackTrace() = %v", stack)
	}

	//assert.Equal(t, expectedStackTrace, stack)
}

func TestBase_WithTracing(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	type fields struct {
		logger    *zap.Logger
		config    config.Config
		Level     level.Level
		Alert     bool
		Message   string
		Fields    map[string]any
		AppFields map[string]any
	}

	type args struct {
		t trace.SpanContext
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Logger
	}{
		{
			name: "Add tracing information",
			fields: fields{
				logger: logger, // Assuming you have a logger instance available.
				Fields: map[string]any{},
			},
			args: args{
				t: trace.NewSpanContext(trace.SpanContextConfig{
					TraceID: [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
						11, 12, 13, 14, 15, 16},
					SpanID: [8]byte{1, 2, 3, 4, 5, 6, 7, 8},
				}),
			},
			want: &Logger{
				logger: logger,
				Data: map[string]any{
					"tracing": map[string]string{
						"trace_id": "0102030405060708090a0b0c0d0e0f10",
						"span_id":  "0102030405060708",
					},
				},
			},
		},
		// Add more test cases for different Tracer implementations
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Logger{
				logger:    tt.fields.logger,
				config:    tt.fields.config,
				Level:     tt.fields.Level,
				Alert:     tt.fields.Alert,
				Message:   tt.fields.Message,
				Data:      tt.fields.Fields,
				AppFields: tt.fields.AppFields,
			}

			assert.Equal(t, tt.want, l.WithTracing(tt.args.t))
		})
	}
}

func TestBase_Write(t *testing.T) {
	// Create a sample config
	c := config.Config{
		AppName: "app_name",
		Version: "1.0.0",
	}

	// Create a buffer to capture the log output
	var buf bytes.Buffer

	// Create a zapcore.Encoder and zapcore.WriteSyncer
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = FixedTimeEncoder // Disable timestamp
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	writeSyncer := zapcore.AddSync(&buf)

	// Create a zapcore.Core that writes to the buffer
	core := zapcore.NewCore(encoder, writeSyncer, zap.NewAtomicLevel())

	// Create a logger with the core
	logger := zap.New(core)

	// Initialize a Logger object with the required parameters
	base := Logger{
		logger:  logger,
		config:  c,
		Level:   level.Info, // Set the level as needed.
		Alert:   true,       // Set the Alert as needed.
		Message: "Test message",
		Data:    map[string]any{},
		AppFields: map[string]any{
			"key1": "value1",
			"key2": "value2",
		},
	}

	// Call the Write method
	base.Write()

	// Add assertions to check the expected behavior
	// For example, you can use the buffer to capture the output and then assert on it.
	// For simplicity, you can use buffer.String() to get the log output.

	// Capture the log output
	logOutput := buf.String()

	// Assert on the expected log message or other criteria based on the expected behavior
	expectedLog := "{\"level\":\"info\",\"ts\":\"fixed\",\"msg\":\"Test message\",\"app_name\":\"app_name\",\"version\":\"1.0.0\",\"alert\":1,\"log_type\":\"\",\"data\":{\"app_name\":{\"key1\":\"value1\",\"key2\":\"value2\"}}}\n"

	assert.Equal(t, expectedLog, logOutput)
}

func TestBase_New(t *testing.T) {
	// Create a zap.Logger for testing purposes
	logger, _ := zap.NewDevelopment()

	// Create a sample config
	c := config.Config{
		// Initialize your configuration fields here.
	}

	// Call the New function to create a Logger object
	lv := level.Info // Set the desired level
	base := New(logger, c, lv, TypeApplication, "abc")

	// Assert on the expected values in the created Logger object
	assert.Equal(t, logger, base.logger)
	assert.Equal(t, c, base.config)
	assert.Equal(t, lv, base.Level)
	assert.Equal(t, false, base.Alert)
	assert.Equal(t, "abc", base.Message)
}
