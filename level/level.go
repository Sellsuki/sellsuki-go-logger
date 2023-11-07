package level

import "go.uber.org/zap/zapcore"

type Level int8

const (
	Debug Level = -1
	Info  Level = 0
	Warn  Level = 1
	Error Level = 2
	Panic Level = 4
	Fatal Level = 5
)

func ToZapLevel(level Level) zapcore.Level {
	switch level {
	case Debug:
		return zapcore.DebugLevel
	case Info:
		return zapcore.InfoLevel
	case Warn:
		return zapcore.WarnLevel
	case Error:
		return zapcore.ErrorLevel
	case Panic:
		return zapcore.PanicLevel
	case Fatal:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
