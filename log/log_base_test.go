package log

import (
	"fmt"
	"github.com/Sellsuki/sellsuki-go-logger/config"
	"github.com/Sellsuki/sellsuki-go-logger/level"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
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
		// Add more test cases with different request payloads if needed
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
			if got := l.WithHTTPReq(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithHTTPReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase_WithHTTPResp(t *testing.T) {
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
		want   Log
	}{
		// TODO: Add test cases.
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
			if got := l.WithHTTPResp(tt.args.resp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithHTTPResp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase_WithKafkaMessage(t *testing.T) {
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
		want   Log
	}{
		// TODO: Add test cases.
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
			if got := l.WithKafkaMessage(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithKafkaMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase_WithKafkaResult(t *testing.T) {
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
		want   Log
	}{
		// TODO: Add test cases.
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
			if got := l.WithKafkaResult(tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithKafkaResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase_WithStackTrace(t *testing.T) {
	type fields struct {
		logger    *zap.Logger
		config    config.Config
		Level     level.Level
		Alert     bool
		Message   string
		Fields    []zap.Field
		AppFields map[string]any
	}
	tests := []struct {
		name   string
		fields fields
		want   Log
	}{
		// TODO: Add test cases.
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
			if got := l.WithStackTrace(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithStackTrace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase_WithTracing(t *testing.T) {
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
		want   Log
	}{
		// TODO: Add test cases.
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
			if got := l.WithTracing(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithTracing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase_Write(t *testing.T) {
	type fields struct {
		logger    *zap.Logger
		config    config.Config
		Level     level.Level
		Alert     bool
		Message   string
		Fields    []zap.Field
		AppFields map[string]any
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
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
			l.Write()
		})
	}
}
