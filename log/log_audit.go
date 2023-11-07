package log

import (
	"github.com/Sellsuki/sellsuki-go-logger/level"
	"go.uber.org/zap"
)

type AuditAction string

type AuditPayload struct {
	// ActorType represents the type of actor performing the action, such as "user", "{{your_bu}}.system", or "{{your_bu}}.store".
	ActorType string `json:"actor_type"`
	// ActorID is the unique identifier of the actor, e.g., "i_love_to_sell_1234", "USR_1234567890".
	ActorID string `json:"actor_id"`
	// Action represents the action being performed
	Action AuditAction `json:"action"`
	// Entity represents the entity that is being acted upon, such as "order", "product", or "customer".
	Entity string `json:"entity"`
	// EntityRefs is a list of unique identifiers of the entity, e.g., ["ODR_1234567890", "ODR_1234567891"].
	EntityRefs []string `json:"entity_refs"`
	// EntityOwnerType represents the owner of the entity, such as "{{your_bu}}.system", "{{your_bu}}.store".
	EntityOwnerType string `json:"entity_owner_type"`
	// EntityOwnerID is the unique identifier of the entity owner, e.g., "i_love_to_sell_1234", "USR_1234567890".
	EntityOwnerID string `json:"entity_owner_id"`
}

type Audit struct {
	Base
}

const (
	AuditActionCreate AuditAction = "create"
	AuditActionUpdate AuditAction = "update"
	AuditActionDelete AuditAction = "delete"
	AuditActionAccess AuditAction = "access"
)

func NewAudit(logger *zap.Logger, msg string, payload AuditPayload) Log {
	l := New(logger, level.Info)
	l.SetMessage(msg)
	l.WithField("data", payload)

	return l
}
