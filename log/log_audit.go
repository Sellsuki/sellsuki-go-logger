package log

import (
	"github.com/Sellsuki/sellsuki-go-logger/config"
	"github.com/Sellsuki/sellsuki-go-logger/level"
	"go.uber.org/zap"
)

type AuditAction string

type AuditPayload struct {
	ActorType       string      `json:"actor_type"`        // represents the type of actor performing the action, such as "user", "{{your_bu}}.system", or "{{your_bu}}.store".
	ActorID         string      `json:"actor_id"`          // the unique identifier of the actor, e.g., "i_love_to_sell_1234", "USR_1234567890".
	Action          AuditAction `json:"action"`            // the action being performed on the entity, such as "create", "update", "delete", "access".
	Entity          string      `json:"entity"`            // the entity that is being acted upon, such as "order", "product", or "customer".
	EntityRefs      []string    `json:"entity_refs"`       // the unique identifier(s) of the entity, e.g., "ORD_1234567890", "some_id_12345".
	EntityOwnerType string      `json:"entity_owner_type"` // the owner of the entity, such as "{{your_bu}}.system", "{{your_bu}}.store".
	EntityOwnerID   string      `json:"entity_owner_id"`   // the unique identifier of the entity owner, e.g., "i_love_to_sell_1234", "USR_1234567890".
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

func NewAudit(logger *zap.Logger, cfg config.Config, msg string, payload AuditPayload) Log {
	l := New(logger, cfg, level.Info, TypeAudit).
		SetMessage(msg).
		WithField("data", payload).(Base)

	return Audit{
		Base: l,
	}
}
