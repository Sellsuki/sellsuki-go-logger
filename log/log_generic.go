package log

import (
	"github.com/Sellsuki/sellsuki-go-logger/config"
	"github.com/Sellsuki/sellsuki-go-logger/level"
	"go.uber.org/zap"
)

type Debug struct {
	Base
}

type Info struct {
	Base
}

type Warn struct {
	Base
}

type Error struct {
	Base
}

type Panic struct {
	Base
}

type Fatal struct {
	Base
}

func NewInfo(logger *zap.Logger, cfg config.Config, msg string) Log {
	l := New(logger, cfg, level.Info)
	l.SetMessage(msg)

	return l
}

func NewDebug(logger *zap.Logger, cfg config.Config, msg string) Log {
	l := New(logger, cfg, level.Debug)
	l.SetMessage(msg)

	return l
}

func NewWarn(logger *zap.Logger, cfg config.Config, msg string) Log {
	l := New(logger, cfg, level.Warn)
	l.SetMessage(msg)

	return l
}

func NewError(logger *zap.Logger, cfg config.Config, msg string) Log {
	l := New(logger, cfg, level.Error)
	l.SetMessage(msg)

	return l
}

func NewPanic(logger *zap.Logger, cfg config.Config, msg string) Log {
	l := New(logger, cfg, level.Panic)
	l.SetMessage(msg)

	return l
}

func NewFatal(logger *zap.Logger, cfg config.Config, msg string) Log {
	l := New(logger, cfg, level.Fatal)
	l.SetMessage(msg)

	return l
}
