package examples

import (
	slog "github.com/Sellsuki/sellsuki-go-logger"
	"github.com/Sellsuki/sellsuki-go-logger/config"
)

func Example_info() {
	slog.Init(config.Config{
		AppName:       "sampleApp",
		Version:       "v1.0.0",
		MaxBodySize:   1048576,
		HardCodedTime: "2023-11-09T14:48:14.803+0700",
	})

	// Call the Info function
	slog.Info("Info message").Write()

	// Output:
	//	{"level":"info","timestamp":"2023-11-09T14:48:14.803+0700","caller":"examples/application_log_test.go:17","message":"Info message","app_name":"sampleApp","version":"v1.0.0","alert":0,"log_type":"application","data":{}}
}
