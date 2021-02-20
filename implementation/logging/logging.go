package logging

import (
	"context"
	"github.com/StephanHCB/go-autumn-logging-rlog/implementation/contextawarelogging"
	auloggingapi "github.com/StephanHCB/go-autumn-logging/api"
)

type RlogLoggingImplementation struct{}

func (l *RlogLoggingImplementation) Ctx(ctx context.Context) auloggingapi.ContextAwareLoggingImplementation {
	return &contextawarelogging.RlogContextAwareLoggingImplementation{Ctx: ctx}
}

func (l *RlogLoggingImplementation) NoCtx() auloggingapi.ContextAwareLoggingImplementation {
	return &contextawarelogging.RlogContextAwareLoggingImplementation{}
}
