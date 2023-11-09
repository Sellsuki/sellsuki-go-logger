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
	return log.New(sukiLogger.zapInstance, sukiLogger.config, level.Debug, log.TypeApplication, msg)
}

func Info(msg string) log.Log {
	return log.New(sukiLogger.zapInstance, sukiLogger.config, level.Info, log.TypeApplication, msg)
}

func Warn(msg string) log.Log {
	return log.New(sukiLogger.zapInstance, sukiLogger.config, level.Warn, log.TypeApplication, msg)
}

func Error(msg string) log.Log {
	return log.New(sukiLogger.zapInstance, sukiLogger.config, level.Error, log.TypeApplication, msg)
}

func Panic(msg string) log.Log {
	return log.New(sukiLogger.zapInstance, sukiLogger.config, level.Panic, log.TypeApplication, msg)
}

func Fatal(msg string) log.Log {
	return log.New(sukiLogger.zapInstance, sukiLogger.config, level.Fatal, log.TypeApplication, msg)
}

func Event(msg string, payload log.EventPayload) log.Log {
	return log.New(sukiLogger.zapInstance, sukiLogger.config, level.Info, log.TypeEvent, msg).
		WithField("event", payload)
}

func Audit(msg string, payload log.AuditPayload) log.Log {
	return log.New(sukiLogger.zapInstance, sukiLogger.config, level.Info, log.TypeAudit, msg).
		WithField("audit", payload)
}

func Kafka(msg string, kMsg *log.KafkaMessagePayload, payload *log.KafkaResultPayload) log.Log {
	return log.New(sukiLogger.zapInstance, sukiLogger.config, level.Info, log.TypeHandlerKafka, msg).
		WithFields(map[string]interface{}{
			"kafka_message": kMsg,
			"kafka_result":  payload,
		})
}

func HTTP(msg string, req *log.HTTPRequestPayload, res *log.HTTPResponsePayload) log.Log {
	return log.New(sukiLogger.zapInstance, sukiLogger.config, level.Info, log.TypeHandlerHTTP, msg).
		WithField("http", map[string]interface{}{
			"http_request":  req,
			"http_response": res,
		})

}
