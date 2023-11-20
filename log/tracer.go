package log

type SpanContext interface {
	TraceID() [16]byte
	SpanID() [8]byte
}
