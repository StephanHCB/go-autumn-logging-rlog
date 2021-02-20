package contextawarelogging

import (
	"context"
	"github.com/StephanHCB/go-autumn-logging-rlog/implementation/leveledlogging"
	"github.com/StephanHCB/go-autumn-logging/api"
)

type RlogContextAwareLoggingImplementation struct{
	Ctx context.Context
}

func (ca *RlogContextAwareLoggingImplementation) Trace() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.RlogLeveledLoggingImplementation{Ctx: ca.Ctx, Level: 5}
}

func (ca *RlogContextAwareLoggingImplementation) Debug() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.RlogLeveledLoggingImplementation{Ctx: ca.Ctx, Level: 4}
}

func (ca *RlogContextAwareLoggingImplementation) Info() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.RlogLeveledLoggingImplementation{Ctx: ca.Ctx, Level: 3}
}

func (ca *RlogContextAwareLoggingImplementation) Warn() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.RlogLeveledLoggingImplementation{Ctx: ca.Ctx, Level: 2}
}

func (ca *RlogContextAwareLoggingImplementation) Error() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.RlogLeveledLoggingImplementation{Ctx: ca.Ctx, Level: 1}
}

func (ca *RlogContextAwareLoggingImplementation) Fatal() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.RlogLeveledLoggingImplementation{Ctx: ca.Ctx, Level: 0}
}

func (ca *RlogContextAwareLoggingImplementation) Panic() auloggingapi.LeveledLoggingImplementation {
	return &leveledlogging.RlogLeveledLoggingImplementation{Ctx: ca.Ctx, Level: 0}
}
