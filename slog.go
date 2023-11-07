package slog

import (
	"fmt"
	"github.com/Sellsuki/sellsuki-go-logger/config"
	"github.com/Sellsuki/sellsuki-go-logger/level"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

var sukiLoggerOnce sync.Once

var sukiLogger *SukiLogger

type SukiLogger struct {
	config      config.Config
	zapInstance *zap.Logger
}

func Init(c ...config.Config) *SukiLogger {
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

		logger, err := zCfg.Build(zap.AddCallerSkip(1))
		if err != nil {
			panic(fmt.Errorf("failed to init logger: %w", err))
		}
		defer logger.Sync()

		sukiLogger = &SukiLogger{zapInstance: logger, config: cfg}
	})

	return sukiLogger
}
