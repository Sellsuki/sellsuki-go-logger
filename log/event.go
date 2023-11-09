package log

type EventAction string
type EventResult string

// EventPayload holds information about an event, including the entity being
// acted upon, its unique reference identifier, the action being performed,
// the result of the action, and the JSON string of event data.
type EventPayload struct {
	Entity      string      `json:"entity"`       // Entity represents the entity that is being acted upon, such as "{{your_bu}}.order", "system", "product", or "customer".
	ReferenceID string      `json:"reference_id"` // ReferenceID is the unique identifier of the entity, e.g., "ODR_1234567890".
	Action      EventAction `json:"action"`       // Action represents the action being performed, such as EventActionCreate, EventActionUpdate, or EventActionDelete.
	Result      EventResult `json:"result"`       // Result represents the result of the action, such as EventResultSuccess or EventResultFailure.
	Data        any         `json:"-"`            // Data is a JSON string containing event data.
	DataJSON    []byte      `json:"data"`         // DataJSON is a JSON string containing event data.
}

const (
	EventActionCreate EventAction = "create"
	EventActionUpdate EventAction = "update"
	EventActionDelete EventAction = "delete"

	EventResultSuccess    EventResult = "success"
	EventResultCompensate EventResult = "compensate"
)
