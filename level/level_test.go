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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToZap(tt.args.level); got != tt.want {
				t.Errorf("ToZap() = %v, want %v", got, tt.want)
			}
		})
	}
}
