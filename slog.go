package slog

import (
	"fmt"
	"github.com/Sellsuki/sellsuki-go-logger/v2/config"
	"github.com/Sellsuki/sellsuki-go-logger/v2/level"
	"github.com/Sellsuki/sellsuki-go-logger/v2/log"
	"github.com/Sellsuki/sellsuki-go-logger/v2/zap_logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
	"time"
)

// SukiLogger is a basically a Singleton wrapper for slog/log/zap_logger

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
	return zap_logger.New(sukiLogger.zapInstance, sukiLogger.config, level.Debug, zap_logger.TypeApplication, msg)
}

func Info(msg string) log.Log {
	return zap_logger.New(sukiLogger.zapInstance, sukiLogger.config, level.Info, zap_logger.TypeApplication, msg)
}

func Warn(msg string) log.Log {
	return zap_logger.New(sukiLogger.zapInstance, sukiLogger.config, level.Warn, zap_logger.TypeApplication, msg)
}

func Error(msg string) log.Log {
	return zap_logger.New(sukiLogger.zapInstance, sukiLogger.config, level.Error, zap_logger.TypeApplication, msg)
}

func Panic(msg string) log.Log {
	return zap_logger.New(sukiLogger.zapInstance, sukiLogger.config, level.Panic, zap_logger.TypeApplication, msg)
}

func Fatal(msg string) log.Log {
	return zap_logger.New(sukiLogger.zapInstance, sukiLogger.config, level.Fatal, zap_logger.TypeApplication, msg)
}

func Event(msg string, payload log.EventPayload) log.Log {
	return zap_logger.New(sukiLogger.zapInstance, sukiLogger.config, level.Info, zap_logger.TypeEvent, msg).
		WithField("event", payload)
}

func Audit(msg string, payload log.AuditPayload) log.Log {
	return zap_logger.New(sukiLogger.zapInstance, sukiLogger.config, level.Info, zap_logger.TypeAudit, msg).
		WithField("audit", payload)
}

func Kafka(msg string, kMsg *log.KafkaMessagePayload, kRes *log.KafkaResultPayload) log.Log {
	payload := map[string]any{}

	if kMsg != nil {
		payload["kafka_message"] = kMsg
	}

	if kRes != nil {
		payload["kafka_result"] = kRes
	}

	return zap_logger.New(sukiLogger.zapInstance, sukiLogger.config, level.Info, zap_logger.TypeHandlerKafka, msg).
		WithFields(payload)
}

func HTTP(msg string, req *log.HTTPRequestPayload, res *log.HTTPResponsePayload) log.Log {
	payload := map[string]any{}

	if req != nil {
		payload["http_request"] = req
	}

	if res != nil {
		payload["http_response"] = res
	}

	return zap_logger.New(sukiLogger.zapInstance, sukiLogger.config, level.Info, zap_logger.TypeHandlerHTTP, msg).
		WithFields(payload)

}
