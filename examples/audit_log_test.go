//go:build example
// +build example

package examples

import (
	"errors"
	slog "github.com/Sellsuki/sellsuki-go-logger"
	"github.com/Sellsuki/sellsuki-go-logger/config"
	"github.com/Sellsuki/sellsuki-go-logger/log"
)

func Example_audit_log() {
	// Do this once in bootstrap file AKA main.go
	slog.Init(config.Config{
		AppName:       "harry_squatter",
		Version:       "the_boy_who_lifted",
		MaxBodySize:   1048576,
		HardCodedTime: "2023-11-09T14:48:14.803+0700",
	})

	slog.Audit("Audit message", log.AuditPayload{
		ActorType:       "hawkward.wizard",
		ActorID:         "magic_user_42",
		Action:          log.AuditActionCreate,
		Entity:          "hawkward.spell.banned",
		EntityRefs:      []string{"dead_rift", "bicep_curse"},
		EntityOwnerType: "fantasy_realm.system",
		EntityOwnerID:   "realm_keeper_5678",
	}).
		WithError(errors.New("you got mail")).
		//WithStackTrace().
		WithAppData("app_data", "app_data_value").
		// WithTracing(trace.SpanFromContext(ctx).SpanContext()). // add tracing from context (otel)
		Write()

	// Output:
	// {"level":"info","timestamp":"2023-11-09T14:48:14.803+0700","caller":"examples/audit_log_test.go:31","message":"Audit message","app_name":"harry_squatter","version":"the_boy_who_lifted","alert":0,"log_type":"audit","data":{"audit":{"actor_type":"hawkward.wizard","actor_id":"magic_user_42","action":"create","entity":"hawkward.spell.banned","entity_refs":["dead_rift","bicep_curse"],"entity_owner_type":"fantasy_realm.system","entity_owner_id":"realm_keeper_5678"},"error":{},"harry_squatter":{"app_data":"app_data_value"}}}
}
