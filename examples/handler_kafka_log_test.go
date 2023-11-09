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

func Example_handler_kafka_log() {
	// Do this once in bootstrap file AKA main.go
	slog.Init(config.Config{
		AppName:       "sampleApp",
		Version:       "v1.0.0",
		MaxBodySize:   1048576,
		HardCodedTime: "2023-11-09T14:48:14.803+0700",
	})

	slog.Kafka("HandlerKafka message received", &log.KafkaMessagePayload{
		Topic:     "topic",
		Partition: 0,
		Offset:    0,
		Payload:   []byte("payload"),
		Headers: map[string]string{
			"header1": "value1",
			"header2": "value2",
		},
		Key:       "key",
		Timestamp: time.Time{},
	}, nil).Write()

	// processing message

	slog.Kafka("HandlerKafka message processed successfully", nil, &log.KafkaResultPayload{
		Duration:  3 * time.Second,
		Committed: true,
	}).Write()

	// Or you can use WithError and WithAppData
	slog.Kafka("HandlerKafka message processed Failed", nil, &log.KafkaResultPayload{
		Duration:  3 * time.Second,
		Committed: false,
	}).WithError(errors.New("error message here")).
		WithAppData("field2", "value2").
		Write()

	// Output:
	// {"level":"info","timestamp":"2023-11-09T14:48:14.803+0700","caller":"examples/handler_kafka_log_test.go:34","message":"HandlerKafka message received","app_name":"sampleApp","version":"v1.0.0","alert":0,"log_type":"handler.kafka","data":{"kafka_message":{"topic":"topic","partition":0,"offset":0,"headers":{"header1":"value1","header2":"value2"},"key":"key","payload":"cGF5bG9hZA==","timestamp":"0001-01-01T00:00:00Z"}}}
	// {"level":"info","timestamp":"2023-11-09T14:48:14.803+0700","caller":"examples/handler_kafka_log_test.go:41","message":"HandlerKafka message processed successfully","app_name":"sampleApp","version":"v1.0.0","alert":0,"log_type":"handler.kafka","data":{"kafka_result":{"duration":3000000000,"committed":true}}}
	// {"level":"info","timestamp":"2023-11-09T14:48:14.803+0700","caller":"examples/handler_kafka_log_test.go:49","message":"HandlerKafka message processed Failed","app_name":"sampleApp","version":"v1.0.0","alert":0,"log_type":"handler.kafka","data":{"error":{},"kafka_result":{"duration":3000000000},"sampleApp":{"field2":"value2"}}}

}
