package log

import "time"

// KafkaMessagePayload represents the payload of a Kafka message.
type KafkaMessagePayload struct {
	Topic     string            `json:"topic"`     // Topic is the Kafka topic to which the message was sent.
	Partition int64             `json:"partition"` // Partition is the partition number where the message is located.
	Offset    int64             `json:"offset"`    // Offset is the offset within the partition for the message.
	Headers   map[string]string `json:"headers"`   // Headers contain key-value pairs of metadata associated with the message.
	Key       string            `json:"key"`       // Key is an optional key associated with the message.
	Payload   []byte            `json:"payload"`   // Payload is the raw message data.
	Timestamp time.Time         `json:"timestamp"` // Timestamp is the time when the message was produced.
}

type KafkaResultPayload struct {
	Duration  time.Duration `json:"duration"`            // Duration is the time it took to produce the message.
	Committed bool          `json:"committed,omitempty"` // Committed is true if the message was successfully commited.
}
