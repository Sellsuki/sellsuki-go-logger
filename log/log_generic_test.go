package log

import (
	"github.com/Sellsuki/sellsuki-go-logger/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestNewDebug(t *testing.T) {
	type args struct {
		logger *zap.Logger
		cfg    config.Config
		msg    string
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
			assert.Equalf(t, tt.want, NewDebug(tt.args.logger, tt.args.cfg, tt.args.msg), "NewDebug(%v, %v, %v)", tt.args.logger, tt.args.cfg, tt.args.msg)
		})
	}
}

func TestNewError(t *testing.T) {
	type args struct {
		logger *zap.Logger
		cfg    config.Config
		msg    string
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
			assert.Equalf(t, tt.want, NewError(tt.args.logger, tt.args.cfg, tt.args.msg), "NewError(%v, %v, %v)", tt.args.logger, tt.args.cfg, tt.args.msg)
		})
	}
}

func TestNewFatal(t *testing.T) {
	type args struct {
		logger *zap.Logger
		cfg    config.Config
		msg    string
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
			assert.Equalf(t, tt.want, NewFatal(tt.args.logger, tt.args.cfg, tt.args.msg), "NewFatal(%v, %v, %v)", tt.args.logger, tt.args.cfg, tt.args.msg)
		})
	}
}

func TestNewInfo(t *testing.T) {
	type args struct {
		logger *zap.Logger
		cfg    config.Config
		msg    string
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
			assert.Equalf(t, tt.want, NewInfo(tt.args.logger, tt.args.cfg, tt.args.msg), "NewInfo(%v, %v, %v)", tt.args.logger, tt.args.cfg, tt.args.msg)
		})
	}
}

func TestNewPanic(t *testing.T) {
	type args struct {
		logger *zap.Logger
		cfg    config.Config
		msg    string
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
			assert.Equalf(t, tt.want, NewPanic(tt.args.logger, tt.args.cfg, tt.args.msg), "NewPanic(%v, %v, %v)", tt.args.logger, tt.args.cfg, tt.args.msg)
		})
	}
}

func TestNewWarn(t *testing.T) {
	type args struct {
		logger *zap.Logger
		cfg    config.Config
		msg    string
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
			assert.Equalf(t, tt.want, NewWarn(tt.args.logger, tt.args.cfg, tt.args.msg), "NewWarn(%v, %v, %v)", tt.args.logger, tt.args.cfg, tt.args.msg)
		})
	}
}
