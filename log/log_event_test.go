package log

import (
	"github.com/Sellsuki/sellsuki-go-logger/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestNewEvent(t *testing.T) {
	type args struct {
		logger  *zap.Logger
		cfg     config.Config
		msg     string
		payload EventPayload
	}
	tests := []struct {
		name string
		args args
		want Log
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewEvent(tt.args.logger, tt.args.cfg, tt.args.msg, tt.args.payload), "NewEvent(%v, %v, %v, %v)", tt.args.logger, tt.args.cfg, tt.args.msg, tt.args.payload)
		})
	}
}
