package log

import "time"

// HTTPRequestPayload represents the payload for an HTTP request.
type HTTPRequestPayload struct {
	Method    string            `json:"method"`     // The HTTP method of the request (e.g., "GET", "POST").
	Path      string            `json:"path"`       // The path of the request URL.
	RemoteIP  string            `json:"remote_ip"`  // The remote IP address of the client.
	Headers   map[string]string `json:"headers"`    // HTTP headers of the request.
	Params    map[string]string `json:"params"`     // URL path parameters.
	Query     map[string]string `json:"query"`      // URL query parameters.
	Body      []byte            `json:"body"`       // The raw HTTP request body.
	RequestID string            `json:"request_id"` // A unique identifier for the request.
}

// HTTPResponsePayload represents the payload for an HTTP response.
type HTTPResponsePayload struct {
	Status    int64         `json:"status"`     // The HTTP status code of the response.
	Duration  time.Duration `json:"duration"`   // The duration of the request processing.
	Body      []byte        `json:"body"`       // The body of the HTTP response.
	Error     string        `json:"error"`      // Any error message or information related to the response.
	RequestID string        `json:"request_id"` // The request identifier associated with this response.
}