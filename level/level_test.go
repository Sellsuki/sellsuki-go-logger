package level

import (
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestToZap(t *testing.T) {
	type args struct {
		level Level
	}
	tests := []struct {
		name string
		args args
		want zapcore.Level
	}{
		{
			name: "Debug",
			args: args{
				level: Debug,
			},
			want: zapcore.DebugLevel,
		},
		{
			name: "Info",
			args: args{
				level: Info,
			},
			want: zapcore.InfoLevel,
		},
		{
			name: "Warn",
			args: args{
				level: Warn,
			},
			want: zapcore.WarnLevel,
		},
		{
			name: "Error",
			args: args{
				level: Error,
			},
			want: zapcore.ErrorLevel,
		},
		{
			name: "Panic",
			args: args{
				level: Panic,
			},
			want: zapcore.PanicLevel,
		},
		{
			name: "Fatal",
			args: args{
				level: Fatal,
			},
			want: zapcore.FatalLevel,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToZap(tt.args.level); got != tt.want {
				t.Errorf("ToZap() = %v, want %v", got, tt.want)
			}
		})
	}
}
