# Sellsuki Logger
A Go wrapper for the sellsuki logging standard

## Installation

```bash
# Go library
go get github.com/Sellsuki/sellsuki-go-logger
```

Or install specific version

```bash
# Go client latest or explicit version
go get github.com/Sellsuki/sellsuki-go-logger/@v1.0.0
```

## Basic Usage

```go
import "github.com/Sellsuki/sellsuki-go-logger"

// Initialize Logger
config := slog.NewProductionConfig()
config.LogLevel = slog.LevelInfo
config.AppName = "sellsuki-logger"
config.Version = "1.0.0"
config.MaxBodySize = 1048576

slog.L().Configure(config)

// Simple Info Log
slog.L().Info(
    "Hello World",       // Log Message
    slog.Any("Yeet", 1), // Some object or variable to include in log
    slog.WithTracing("a", "b"), // Tracing information (Optional)
)
```

## Request Log

```go
import "github.com/Sellsuki/sellsuki-go-logger"

// Initialize Logger
config := slog.NewProductionConfig()
config.LogLevel = slog.LevelInfo
config.AppName = "sellsuki-logger"
config.Version = "1.0.0"
config.MaxBodySize = 1048576

slog.L().Configure(config)

// HTTP Request Log
slog.L().RequestHTTP(
    "such wow",
        slog.WithHTTPRequest(
        "POST",                     // Method
        "/dodge/wow",               // Path
        "127.0.0.1",                // Request IP Address 
        map[string]string{          // Request header
            "Content-Type": "application/json",
        },
        map[string]string{          // Params in URI
            "user_id": "777",
        },
        map[string]string{          // Query Params
            "keyword": "yikes",
        },
        "{\"such\": \"wow\"}",      // Raw Body string
    ),
    slog.WithHTTPResponse(
        200,                        // Status Code
        0.0167777,                  // Request Process Duration
        "{\"such\": \"wow\"}",      // Raw Response Body string
        slog.WithError(             // Error
			"item_not_found",       // Error name
            "/dodge/wow.go:35",     // Caller
			"some stack trace here" // Stacktrace
        ), 
    ),
    slog.WithTracing(
        "trace_id",                 // Tracing ID
        "span_id"                   // Span ID
    ),
)

// Kafka Request Log
slog.L().RequestKafka(
    "write something about kafka",
    slog.WithKafkaMessage(
        "topic.name.here",          // Kafka Topic Name
        0,                          // Partition
        500,                        // Offset
        map[string]string{          // Headers
            "header_key": "header_value",   
        },
        "kafka_key",                // Keys
        "kafka payload here",       // Message Payload
        time.Now(),                 // Timestamp
    ),
    slog.WithKafkaResult(
        0.016777,                   // Process Duration
        slog.WithError(
            "item_not_found",       // Error name
            "/dodge/wow.go:35",     // Caller
            "some stack trace here" // Stacktrace
        ),
    ),
    slog.WithTracing(
		"trace_id",                 // Tracing ID
		"span_id"                   // Span ID
    ),
)

```

## Event Log

```go
import "github.com/Sellsuki/sellsuki-go-logger"

// Initialize Logger
config := slog.NewProductionConfig()
config.LogLevel = slog.LevelInfo
config.AppName = "sellsuki-logger"
config.Version = "1.0.0"
config.MaxBodySize = 1048576

slog.L().Configure(config)

// Event Log
slog.L().Event(
    "event message",        // Log Message
    slog.WithEvent(         
        "order",            // Entity
        slog.ActionCreate,  // Event action (Create, Update, Delete)
        slog.ResultSuccess, // Event result (Success, Compensate)
        "",                 // Raw data
        "ref_id",           // Normalized reference id
    ),
    slog.WithTracing(
		"tracing_id", 
		"span_id"
    ),
)
```

## Application Log

```go
import "github.com/Sellsuki/sellsuki-go-logger"

// Initialize Logger
config := slog.NewProductionConfig()
config.LogLevel = slog.LevelInfo
config.AppName = "sellsuki-logger"
config.Version = "1.0.0"
config.MaxBodySize = 1048576

slog.L().Configure(config)


// Debug Log
slog.L().Debug(
    "Hello World",       // Log Message
    slog.Any("Yeet", 1), // Some object or variable to include in log
    slog.WithTracing("a", "b"), // Tracing information (Optional)
)

// Info Log
slog.L().Info(
    "Hello World",       // Log Message
    slog.Any("Yeet", 1), // Some object or variable to include in log
    slog.WithTracing("a", "b"), // Tracing information (Optional)
)

// Warning Log
slog.L().Warn(
    "Hello World",       // Log Message
    slog.Any("Yeet", 1), // Some object or variable to include in log
    slog.WithTracing("a", "b"), // Tracing information (Optional)
)

// Error Log
slog.L().Error(
    "Hello World",       // Log Message
    slog.Any("Yeet", 1), // Some object or variable to include in log
    slog.WithTracing("a", "b"), // Tracing information (Optional)
)

// Fatal Log, This log type exit the process after the log has written
slog.L().Fatal(
    "Hello World",       // Log Message
    slog.Any("Yeet", 1), // Some object or variable to include in log
    slog.WithTracing("a", "b"), // Tracing information (Optional)
)

// Fatal Log, This log type call panic after the log has written
slog.L().Panic(
    "Hello World",       // Log Message
    slog.Any("Yeet", 1), // Some object or variable to include in log
    slog.WithTracing("a", "b"), // Tracing information (Optional)
)
```
