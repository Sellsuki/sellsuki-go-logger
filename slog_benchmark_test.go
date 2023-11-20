package slog

import (
	"github.com/Sellsuki/sellsuki-go-logger/v2/config"
	"github.com/Sellsuki/sellsuki-go-logger/v2/level"
	"testing"
)

func BenchmarkDebug(b *testing.B) {
	// Initialize a zap Logger
	Init(config.Config{
		LogLevel:    level.Debug,
		AppName:     "sampleApp",
		Version:     "v1.0.0",
		MaxBodySize: 1048576,
	})

	// Reset the timer and start benchmarking
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Call the Debug function
		Debug("Benchmark debug message")
	}
}

// Example BenchmarkInfo function:
func BenchmarkInfo(b *testing.B) {
	// Create a sample config and initialize the zap Logger
	// ...

	// Reset the timer and start benchmarking
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Call the Info function
		Info("Benchmark info message")
	}
}
