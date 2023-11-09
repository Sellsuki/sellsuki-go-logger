package log

import (
	"fmt"
	"runtime"
)

func captureStackTrace(skip int) string {
	pc := make([]uintptr, 10) // Adjust the size as needed
	n := runtime.Callers(skip, pc)
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
