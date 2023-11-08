package log

import (
	"context"
	"github.com/Sellsuki/sellsuki-go-logger/config"
	"github.com/Sellsuki/sellsuki-go-logger/level"
	"reflect"
	"testing"
)

func TestBase_SetAlert(t *testing.T) {
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
		bool bool
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
			if got := l.SetAlert(tt.args.bool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAlert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase_SetLevel(t *testing.T) {
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
		level level.Level
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
			if got := l.SetLevel(tt.args.level); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase_SetMessage(t *testing.T) {
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
		msg string
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
			if got := l.SetMessage(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase_WithAppData(t *testing.T) {
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
		key   string
		value any
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
			if got := l.WithAppData(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithAppData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase_WithContextSession(t *testing.T) {
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
		ctx context.Context
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
			if got := l.WithContextSession(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithContextSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase_WithError(t *testing.T) {
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
		err error
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
			if got := l.WithError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase_WithField(t *testing.T) {
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
		key   string
		value any
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
			if got := l.WithField(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase_WithFields(t *testing.T) {
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
		fields map[string]any
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
			if got := l.WithFields(tt.args.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase_WithHTTPReq(t *testing.T) {
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
		req HTTPRequestPayload
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

func TestNew(t *testing.T) {
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
			if got := New(tt.args.logger, tt.args.cfg, tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
