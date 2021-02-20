package leveledlogging

import (
	"context"
	"fmt"
	auloggingapi "github.com/StephanHCB/go-autumn-logging/api"
	"github.com/go-logr/logr"
	_ "github.com/go-logr/logr"
)

type RlogLeveledLoggingImplementation struct {
	Ctx context.Context
	Level int
	Err error
	Values map[string]string
}

func (lv *RlogLeveledLoggingImplementation) WithErr(err error) auloggingapi.LeveledLoggingImplementation {
	lv.Err = err
	return lv
}

func (lv *RlogLeveledLoggingImplementation) With(key string, value string) auloggingapi.LeveledLoggingImplementation {
	lv.Values[key] = value
	return lv
}

func (lv *RlogLeveledLoggingImplementation) Print(v ...interface{}) {
	rlogLogger := logr.FromContextOrDiscard(lv.Ctx).V(lv.Level)
	for k, v := range lv.Values {
		rlogLogger = rlogLogger.WithValues(k, v)
	}
	if lv.Err != nil {
		rlogLogger.Error(lv.Err, fmt.Sprint(v...))
	} else {
		rlogLogger.Info(fmt.Sprint(v...))
	}
}

func (lv *RlogLeveledLoggingImplementation) Printf(format string, v ...interface{}) {
	rlogLogger := logr.FromContextOrDiscard(lv.Ctx).V(lv.Level)
	for k, v := range lv.Values {
		rlogLogger = rlogLogger.WithValues(k, v)
	}
	if lv.Err != nil {
		rlogLogger.Error(lv.Err, fmt.Sprintf(format, v...))
	} else {
		rlogLogger.Info(fmt.Sprintf(format, v...))
	}
}
