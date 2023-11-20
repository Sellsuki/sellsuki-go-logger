package zap_logger

import (
	"fmt"
	"runtime"
)

func CaptureStackTrace(skip int) string {
	pc := make([]uintptr, 10) // Adjust the size as needed
	n := runtime.Callers(skip+1, pc)
	frames := runtime.CallersFrames(pc[:n])
	stackTrace := ""
	for {
		frame, more := frames.Next()
		stackTrace += fmt.Sprintf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}
	return stackTrace
}
