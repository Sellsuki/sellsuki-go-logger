package examples

import (
	"errors"
	slog "github.com/Sellsuki/sellsuki-go-logger"
	"github.com/Sellsuki/sellsuki-go-logger/config"
	"github.com/Sellsuki/sellsuki-go-logger/log"
)

func Example_audit_log() {
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

}
