package log

import (
	"github.com/Sellsuki/sellsuki-go-logger/config"
	"github.com/Sellsuki/sellsuki-go-logger/level"
	"go.uber.org/zap"
	"testing"
)

// BenchmarkWrite measures the performance of the Write method.
func BenchmarkWrite(b *testing.B) {

	// Log to null
	logger := zap.NewNop()

	// Create a sample config
	c := config.Config{
		MaxBodySize: 1048576,
		AppName:     "unknown",
		Version:     "v0.0.0",
	}

	// Initialize a Base object with the required parameters
	base := New(logger, c, level.Info, TypeApplication).
		SetMessage("Benchmark message").
		withField("key", "value")

	// Reset the timer and start benchmarking
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Call the Write method and capture the log output
		base.Write()
	}
}

func BenchmarkAllMethods(b *testing.B) {
	// Create a zap.Logger for testing purposes
	logger := zap.NewNop()

	// Create a sample config
	c := config.Config{
		MaxBodySize: 1048576,
		AppName:     "unknown",
		Version:     "v0.0.0",
	}

	// Reset the timer and start benchmarking
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Chain multiple methods on the Base object
		base := New(logger, c, level.Info, TypeApplication).
			SetMessage("Benchmark message").
			withField("key", "value").
			WithError(nil).
			WithStackTrace().
			WithAppData("app_key", "app_value").
			WithHTTPReq(HTTPRequestPayload{}).
			WithHTTPResp(HTTPResponsePayload{}).
			WithKafkaMessage(KafkaMessagePayload{}).
			WithKafkaResult(KafkaResultPayload{})

		// Call the Write method
		base.Write()
	}
}

func BenchmarkAllMethodsExceptStackTrace(b *testing.B) {
	// Create a zap.Logger for testing purposes
	logger := zap.NewNop()

	// Create a sample config
	c := config.Config{
		MaxBodySize: 1048576,
		AppName:     "unknown",
		Version:     "v0.0.0",
	}

	// Reset the timer and start benchmarking
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Chain multiple methods on the Base object
		base := New(logger, c, level.Info, TypeApplication).
			SetMessage("Benchmark message").
			withField("key", "value").
			WithError(nil).
			WithAppData("app_key", "app_value").
			WithHTTPReq(HTTPRequestPayload{}).
			WithHTTPResp(HTTPResponsePayload{}).
			WithKafkaMessage(KafkaMessagePayload{}).
			WithKafkaResult(KafkaResultPayload{})

		// Call the Write method
		base.Write()
	}
}
