package slog

import (
	"fmt"
	"github.com/Sellsuki/sellsuki-go-logger/config"
	"github.com/Sellsuki/sellsuki-go-logger/level"
	"github.com/Sellsuki/sellsuki-go-logger/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

type SukiLogger struct {
	config      config.Config
	zapInstance *zap.Logger
}

var sukiLoggerOnce sync.Once

var sukiLogger *SukiLogger

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

		zCfg := zap.NewProductionConfig()
		zCfg.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		zCfg.EncoderConfig.MessageKey = "message"
		zCfg.EncoderConfig.TimeKey = "timestamp"
		zCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		zCfg.Level = zap.NewAtomicLevelAt(level.ToZap(cfg.LogLevel))

		if cfg.Readable {
			zCfg.Encoding = "console"
			zCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
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
	Init()

	return log.NewDebug(sukiLogger.zapInstance, sukiLogger.config, msg)
}

func Info(msg string) log.Log {
	Init()

	return log.NewInfo(sukiLogger.zapInstance, sukiLogger.config, msg)
}

func Warn(msg string) log.Log {
	Init()

	return log.NewWarn(sukiLogger.zapInstance, sukiLogger.config, msg)
}

func Error(msg string) log.Log {
	Init()

	return log.NewError(sukiLogger.zapInstance, sukiLogger.config, msg)
}

func Panic(msg string) log.Log {
	Init()

	return log.NewPanic(sukiLogger.zapInstance, sukiLogger.config, msg)
}

func Fatal(msg string) log.Log {
	Init()

	return log.NewFatal(sukiLogger.zapInstance, sukiLogger.config, msg)
}

func Event(msg string, payload log.EventPayload) log.Log {
	Init()

	return log.NewEvent(sukiLogger.zapInstance, sukiLogger.config, msg, payload)
}

func Audit(msg string, payload log.AuditPayload) log.Log {
	Init()

	return log.NewAudit(sukiLogger.zapInstance, sukiLogger.config, msg, payload)
}
