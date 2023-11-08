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

func NewDebug(logger *zap.Logger, cfg config.Config, msg string) Log {
	l := New(logger, cfg, level.Debug, TypeApplication).SetMessage(msg).(Base)

	return Debug{
		Base: l,
	}
}

func NewInfo(logger *zap.Logger, cfg config.Config, msg string) Log {
	l := New(logger, cfg, level.Info, TypeApplication).SetMessage(msg).(Base)

	return Info{
		Base: l,
	}
}

func NewWarn(logger *zap.Logger, cfg config.Config, msg string) Log {
	l := New(logger, cfg, level.Warn, TypeApplication).SetMessage(msg).(Base)

	return Warn{
		Base: l,
	}
}

func NewError(logger *zap.Logger, cfg config.Config, msg string) Log {
	l := New(logger, cfg, level.Error, TypeApplication).SetMessage(msg).(Base)

	return Error{
		Base: l,
	}
}

func NewPanic(logger *zap.Logger, cfg config.Config, msg string) Log {
	l := New(logger, cfg, level.Panic, TypeApplication).SetMessage(msg).(Base)

	return Panic{
		Base: l,
	}
}

func NewFatal(logger *zap.Logger, cfg config.Config, msg string) Log {
	l := New(logger, cfg, level.Fatal, TypeApplication).SetMessage(msg).(Base)

	return Fatal{
		Base: l,
	}
}
