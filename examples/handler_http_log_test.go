//go:build example
// +build example

package examples

import (
	"errors"
	slog "github.com/Sellsuki/sellsuki-go-logger"
	"github.com/Sellsuki/sellsuki-go-logger/config"
	"github.com/Sellsuki/sellsuki-go-logger/log"
	"time"
)

func Example_handler_http_log() {
	// Do this once in bootstrap file AKA main.go
	slog.Init(config.Config{
		AppName:       "sampleApp",
		Version:       "v1.0.0",
		MaxBodySize:   1048576,
		HardCodedTime: "2023-11-09T14:48:14.803+0700",
	})

	// Simulate an incoming HTTP request
	requestPayload := &log.HTTPRequestPayload{
		Method:    "POST",
		Handler:   "GetResourceById",   // You can put the function name here or the handler name, whatever you want
		Path:      "/api/{{resource}}", // If possible keep the path as a template for easier searching
		RemoteIP:  "192.168.1.1",
		Headers:   map[string]string{"Content-Type": "application/json"},
		Params:    map[string]string{"resource": "123"},
		Query:     map[string]string{"param1": "value1"},
		Body:      (`{"key": "value"}`),
		RequestID: "unique-request-id",
	}

	slog.HTTP("HandlerHTTP request received", requestPayload, nil).Write()

	// Processing the HTTP request

	// Simulate an HTTP response
	responsePayload := &log.HTTPResponsePayload{
		Status:    200,
		Duration:  2 * time.Second,
		Body:      (`{"result": "success"}`),
		RequestID: "unique-request-id",
	}

	slog.HTTP("HandlerHTTP request processed successfully", nil, responsePayload).Write()

	// Or you can use WithError and WithAppData
	slog.HTTP("HandlerHTTP request processing failed", nil, nil).
		WithError(errors.New("error message here")).
		WithAppData("field2", "value2").
		Write()

	// Output:
	// {"level":"info","timestamp":"2023-11-09T14:48:14.803+0700","caller":"examples/handler_http_log_test.go:36","message":"HandlerHTTP request received","app_name":"sampleApp","version":"v1.0.0","alert":0,"log_type":"handler.http","data":{"http_request":{"method":"POST","handler":"GetResourceById","path":"/api/{{resource}}","remote_ip":"192.168.1.1","headers":{"Content-Type":"application/json"},"params":{"resource":"123"},"query":{"param1":"value1"},"body":"eyJrZXkiOiAidmFsdWUifQ==","request_id":"unique-request-id"}}}
	// {"level":"info","timestamp":"2023-11-09T14:48:14.803+0700","caller":"examples/handler_http_log_test.go:48","message":"HandlerHTTP request processed successfully","app_name":"sampleApp","version":"v1.0.0","alert":0,"log_type":"handler.http","data":{"http_response":{"status":200,"duration":2000000000,"body":"eyJyZXN1bHQiOiAic3VjY2VzcyJ9","request_id":"unique-request-id"}}}
	// {"level":"info","timestamp":"2023-11-09T14:48:14.803+0700","caller":"examples/handler_http_log_test.go:54","message":"HandlerHTTP request processing failed","app_name":"sampleApp","version":"v1.0.0","alert":0,"log_type":"handler.http","data":{"error":{},"sampleApp":{"field2":"value2"}}}
}