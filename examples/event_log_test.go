package examples

import (
	slog "github.com/Sellsuki/sellsuki-go-logger"
	"github.com/Sellsuki/sellsuki-go-logger/config"
	"github.com/Sellsuki/sellsuki-go-logger/log"
)

func Example_event_log() {
	slog.Init(config.Config{
		AppName:       "lord_of_the_rim",
		Version:       "the_rim_of_lovers",
		MaxBodySize:   1048576,
		HardCodedTime: "2023-11-09T14:48:14.803+0700",
	})

	slog.Event("Event message", log.EventPayload{
		Entity:      "rim",
		ReferenceID: "#1",
		Action:      log.EventActionCreate,
		Result:      log.EventResultSuccess,
		Data: map[string]interface{}{
			"rim_name":    "The One Ring",
			"power_level": "Ultimate",
			"creator":     "Sauron",
		},
	}).
		WithAppData("app_data", "app_data_value").
		// WithTracing(trace.SpanFromContext(ctx).SpanContext()). // add tracing from context (otel)
		Write()

	// Output:
	// {"level":"info","timestamp":"2023-11-09T14:48:14.803+0700","caller":"examples/event_log_test.go:30","message":"Event message","app_name":"lord_of_the_rim","version":"the_rim_of_lovers","alert":0,"log_type":"event","data":{"event":{"entity":"rim","reference_id":"#1","action":"create","result":"success","data":null},"lord_of_the_rim":{"app_data":"app_data_value"}}}
}
