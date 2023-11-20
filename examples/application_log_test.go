//go:build example
// +build example

package examples

import (
	"errors"
	slog "github.com/Sellsuki/sellsuki-go-logger/v2"
	"github.com/Sellsuki/sellsuki-go-logger/v2/config"
)

func Example_application_log() {
	// Do this once in bootstrap file AKA main.go
	slog.Init(config.Config{
		AppName:       "sampleApp",
		Version:       "v1.0.0",
		MaxBodySize:   1048576,
		HardCodedTime: "2023-11-09T14:48:14.803+0700",
	})

	// Call the Info function
	slog.Info("Info message").Write()

	// Output:
	// {"level":"info","timestamp":"2023-11-09T14:48:14.803+0700","caller":"examples/application_log_test.go:22","message":"Info message","app_name":"sampleApp","version":"v1.0.0","alert":0,"log_type":"application","data":{}}

}

func Example_application_with_data() {
	// Do this once in bootstrap file AKA main.go
	slog.Init(config.Config{
		AppName:       "sampleApp",
		Version:       "v1.0.0",
		MaxBodySize:   1048576,
		HardCodedTime: "2023-11-09T14:48:14.803+0700",
	})

	// Call the Info function
	slog.Info("Info message").
		WithError(errors.New("error message here")).
		// WithStackTrace(). // uncomment this line to enable stack trace, because it's different on each machine
		WithAppData("field2", "value2").
		// WithTracing(trace.SpanFromContext(ctx).SpanContext()). // add tracing from context (otel)
		Write()

	// Output:
	// {"level":"info","timestamp":"2023-11-09T14:48:14.803+0700","caller":"examples/application_log_test.go:44","message":"Info message","app_name":"sampleApp","version":"v1.0.0","alert":0,"log_type":"application","data":{"error":{},"sampleApp":{"field2":"value2"}}}

}
