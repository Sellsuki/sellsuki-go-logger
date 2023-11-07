package log

import (
	"encoding/json"
	"github.com/Sellsuki/sellsuki-go-logger/config"
	"github.com/Sellsuki/sellsuki-go-logger/level"
	"go.uber.org/zap"
)

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

type Event struct {
	Base
}

const (
	EventActionCreate EventAction = "create"
	EventActionUpdate EventAction = "update"
	EventActionDelete EventAction = "delete"

	EventResultSuccess    EventResult = "success"
	EventResultCompensate EventResult = "compensate"
)

func NewEvent(logger *zap.Logger, cfg config.Config, msg string, payload EventPayload) Log {
	l := New(logger, cfg, level.Info)
	l.SetMessage(msg)

	if payload.Data != nil {
		payload.DataJSON, _ = json.Marshal(payload.Data)
	}

	l.WithField("event", payload)

	return l
}
