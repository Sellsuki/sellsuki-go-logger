package slog

import (
	"fmt"
	"github.com/Sellsuki/sellsuki-go-logger/config"
	"github.com/Sellsuki/sellsuki-go-logger/level"
	"github.com/Sellsuki/sellsuki-go-logger/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
	"time"
)

type SukiLogger struct {
	config      config.Config
	zapInstance *zap.Logger
}

var sukiLoggerOnce sync.Once

var sukiLogger *SukiLogger

// Init initialize the logger
// Do not run this function in parallel
func Init(c ...config.Config) {
	sukiLoggerOnce.Do(func() {
		cfg := config.Config{
			LogLevel:    level.Info,
			AppName:     "unknown",
			Version:     "v0.0.0",
			MaxBodySize: 1048576,
		}
		if len(c) > 0 {
			cfg = c[0]
		}

		zCfg := zap.Config{
			Level:       zap.NewAtomicLevelAt(level.ToZap(cfg.LogLevel)),
			Development: false,
			Sampling: &zap.SamplingConfig{
				Initial:    100,
				Thereafter: 100,
			},
			Encoding: "json",
			EncoderConfig: zapcore.EncoderConfig{
				TimeKey:        "timestamp",
				LevelKey:       "level",
				NameKey:        "logger",
				CallerKey:      "caller",
				FunctionKey:    zapcore.OmitKey,
				MessageKey:     "message",
				StacktraceKey:  "stacktrace",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.SecondsDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			},
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stdout"},
		}

		if cfg.Readable {
			zCfg.Encoding = "console"
			zCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		}

		if cfg.HardCodedTime != "" {
			zCfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(cfg.HardCodedTime)
			}
		}

		logger, err := zCfg.Build(zap.AddCallerSkip(1))
		if err != nil {
			panic(fmt.Errorf("failed to init logger: %w", err))
		}

		defer logger.Sync()

		sukiLogger = &SukiLogger{zapInstance: logger, config: cfg}
	})
}

func Debug(msg string) log.Log {
	return log.NewDebug(sukiLogger.zapInstance, sukiLogger.config, msg)
}

func Info(msg string) log.Log {
	return log.NewInfo(sukiLogger.zapInstance, sukiLogger.config, msg)
}

func Warn(msg string) log.Log {
	return log.NewWarn(sukiLogger.zapInstance, sukiLogger.config, msg)
}

func Error(msg string) log.Log {
	return log.NewError(sukiLogger.zapInstance, sukiLogger.config, msg)
}

func Panic(msg string) log.Log {
	return log.NewPanic(sukiLogger.zapInstance, sukiLogger.config, msg)
}

func Fatal(msg string) log.Log {
	return log.NewFatal(sukiLogger.zapInstance, sukiLogger.config, msg)
}

func Event(msg string, payload log.EventPayload) log.Log {
	return log.NewEvent(sukiLogger.zapInstance, sukiLogger.config, msg, payload)
}

func Audit(msg string, payload log.AuditPayload) log.Log {
	return log.NewAudit(sukiLogger.zapInstance, sukiLogger.config, msg, payload)
}
