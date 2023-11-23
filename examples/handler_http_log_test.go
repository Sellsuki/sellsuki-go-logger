//go:build example
// +build example

package examples

import (
	"errors"
	slog "github.com/Sellsuki/sellsuki-go-logger/v2"
	"github.com/Sellsuki/sellsuki-go-logger/v2/config"
	"github.com/Sellsuki/sellsuki-go-logger/v2/log"
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
		Duration:  2,
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
	//{"level":"info","timestamp":"2023-11-09T14:48:14.803+0700","caller":"examples/handler_http_log_test.go:35","message":"HandlerHTTP request received","app_name":"sampleApp","version":"v1.0.0","alert":0,"log_type":"handler.http","data":{"http_request":{"method":"POST","handler":"GetResourceById","path":"/api/{{resource}}","remote_ip":"192.168.1.1","headers":{"Content-Type":"application/json"},"params":{"resource":"123"},"query":{"param1":"value1"},"body":"{\"key\": \"value\"}","request_id":"unique-request-id"}}}
	//{"level":"info","timestamp":"2023-11-09T14:48:14.803+0700","caller":"examples/handler_http_log_test.go:47","message":"HandlerHTTP request processed successfully","app_name":"sampleApp","version":"v1.0.0","alert":0,"log_type":"handler.http","data":{"http_response":{"status":200,"duration":2,"body":"{\"result\": \"success\"}","request_id":"unique-request-id","headers":null}}}
	//{"level":"info","timestamp":"2023-11-09T14:48:14.803+0700","caller":"examples/handler_http_log_test.go:53","message":"HandlerHTTP request processing failed","app_name":"sampleApp","version":"v1.0.0","alert":0,"log_type":"handler.http","data":{"error":"error message here","sampleApp":{"field2":"value2"}}}
}
