package slog

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"testing"
)

func TestError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want LogField
	}{
		{
			name: "Contain error case",
			args: args{
				err: errors.New("hello world"),
			},
			want: LogField{
				Key:   "error",
				Value: "hello world",
			},
		},
		{
			name: "Error was nil",
			args: args{
				err: nil,
			},
			want: LogField{
				Key:   "error",
				Value: "",
			},
		},
		{
			name: "Error is wrapped",
			args: args{
				err: fmt.Errorf("this is outer: %w", errors.New("this is inner")),
			},
			want: LogField{
				Key:   "error",
				Value: "this is outer: this is inner",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Error(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithEvent(t *testing.T) {
	type args struct {
		entity string
		action EventAction
		result EventResult
		data   interface{}
		refID  string
	}
	tests := []struct {
		name string
		args args
		want EventLog
	}{
		{
			name: "Event payload is string",
			args: args{
				entity: "order",
				action: ActionCreate,
				result: ResultSuccess,
				data:   "hello world",
				refID:  "1",
			},
			want: EventLog{
				Entity:      "order",
				Action:      "create",
				Result:      "success",
				ReferenceID: "1",
				Data:        "hello world",
			},
		},
		{
			name: "Event payload is nil",
			args: args{
				entity: "order",
				action: ActionCreate,
				result: ResultSuccess,
				data:   nil,
				refID:  "1",
			},
			want: EventLog{
				Entity:      "order",
				Action:      "create",
				Result:      "success",
				ReferenceID: "1",
				Data:        "",
			},
		},
		{
			name: "Event payload is object",
			args: args{
				entity: "order",
				action: ActionCreate,
				result: ResultSuccess,
				data: eventStruct{
					ID:   1,
					Name: "2",
				},
				refID: "1",
			},
			want: EventLog{
				Entity:      "order",
				Action:      "create",
				Result:      "success",
				ReferenceID: "1",
				Data:        "{\"ID\":1,\"Name\":\"2\"}",
			},
		},
		{
			name: "Event payload is int",
			args: args{
				entity: "order",
				action: ActionCreate,
				result: ResultSuccess,
				data:   1,
				refID:  "1",
			},
			want: EventLog{
				Entity:      "order",
				Action:      "create",
				Result:      "success",
				ReferenceID: "1",
				Data:        "1",
			},
		}, {
			name: "Event payload is boolean",
			args: args{
				entity: "order",
				action: ActionCreate,
				result: ResultSuccess,
				data:   true,
				refID:  "1",
			},
			want: EventLog{
				Entity:      "order",
				Action:      "create",
				Result:      "success",
				ReferenceID: "1",
				Data:        "true",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithEvent(tt.args.entity, tt.args.action, tt.args.result, tt.args.data, tt.args.refID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithTracing(t *testing.T) {
	type args struct {
		traceID   string
		spanID    string
		requestID []string
	}
	tests := []struct {
		name string
		args args
		want TraceInfo
	}{
		{
			name: "With request ID",
			args: args{
				traceID:   "trace_id",
				spanID:    "span_id",
				requestID: []string{"request_id"},
			},
			want: TraceInfo{
				TraceID:   "trace_id",
				SpanID:    "span_id",
				RequestID: "request_id",
			},
		},
		{
			name: "No request ID",
			args: args{
				traceID:   "trace_id",
				spanID:    "span_id",
				requestID: []string{},
			},
			want: TraceInfo{
				TraceID:   "trace_id",
				SpanID:    "span_id",
				RequestID: "",
			},
		},
		{
			name: "RequestID is nil",
			args: args{
				traceID:   "trace_id",
				spanID:    "span_id",
				requestID: nil,
			},
			want: TraceInfo{
				TraceID:   "trace_id",
				SpanID:    "span_id",
				RequestID: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithTracing(tt.args.traceID, tt.args.spanID, tt.args.requestID...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithTracing() = %v, want %v", got, tt.want)
			}
		})
	}
}
