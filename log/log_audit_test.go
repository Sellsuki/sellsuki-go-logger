package log

import (
	"github.com/Sellsuki/sellsuki-go-logger/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestNewAudit(t *testing.T) {
	type args struct {
		logger  *zap.Logger
		cfg     config.Config
		msg     string
		payload AuditPayload
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
			assert.Equalf(t, tt.want, NewAudit(tt.args.logger, tt.args.cfg, tt.args.msg, tt.args.payload), "NewAudit(%v, %v, %v, %v)", tt.args.logger, tt.args.cfg, tt.args.msg, tt.args.payload)
		})
	}
}
