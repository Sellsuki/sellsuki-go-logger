package log

type Stringer interface {
	String() string
}

type Tracer interface {
	TraceID() Stringer
	SpanID() Stringer
}
